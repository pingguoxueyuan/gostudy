// Go from multi-language-benchmark/src/havlak/go_pro

// Copyright 2011 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Test Program for the Havlak loop finder.
//
// This program constructs a fairly large control flow
// graph and performs loop recognition. This is the Go
// version.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

type BasicBlock struct {
	Name     int
	InEdges  []*BasicBlock
	OutEdges []*BasicBlock
}

func NewBasicBlock(name int) *BasicBlock {
	return &BasicBlock{Name: name}
}

func (bb *BasicBlock) Dump() {
	fmt.Printf("BB#%06d:", bb.Name)
	if len(bb.InEdges) > 0 {
		fmt.Printf(" in :")
		for _, iter := range bb.InEdges {
			fmt.Printf(" BB#%06d", iter.Name)
		}
	}
	if len(bb.OutEdges) > 0 {
		fmt.Print(" out:")
		for _, iter := range bb.OutEdges {
			fmt.Printf(" BB#%06d", iter.Name)
		}
	}
	fmt.Printf("\n")
}

func (bb *BasicBlock) NumPred() int {
	return len(bb.InEdges)
}

func (bb *BasicBlock) NumSucc() int {
	return len(bb.OutEdges)
}

func (bb *BasicBlock) AddInEdge(from *BasicBlock) {
	bb.InEdges = append(bb.InEdges, from)
}

func (bb *BasicBlock) AddOutEdge(to *BasicBlock) {
	bb.OutEdges = append(bb.OutEdges, to)
}

//-----------------------------------------------------------

type CFG struct {
	Blocks []*BasicBlock
	Start  *BasicBlock
}

func NewCFG() *CFG {
	return &CFG{}
}

func (cfg *CFG) NumNodes() int {
	return len(cfg.Blocks)
}

func (cfg *CFG) CreateNode(node int) *BasicBlock {
	if node < len(cfg.Blocks) {
		return cfg.Blocks[node]
	}
	if node != len(cfg.Blocks) {
		println("oops", node, len(cfg.Blocks))
		panic("wtf")
	}
	bblock := NewBasicBlock(node)
	cfg.Blocks = append(cfg.Blocks, bblock)

	if len(cfg.Blocks) == 1 {
		cfg.Start = bblock
	}

	return bblock
}

func (cfg *CFG) Dump() {
	for _, n := range cfg.Blocks {
		n.Dump()
	}
}

//-----------------------------------------------------------

type BasicBlockEdge struct {
	Dst *BasicBlock
	Src *BasicBlock
}

func NewBasicBlockEdge(cfg *CFG, from int, to int) *BasicBlockEdge {
	self := new(BasicBlockEdge)
	self.Src = cfg.CreateNode(from)
	self.Dst = cfg.CreateNode(to)

	self.Src.AddOutEdge(self.Dst)
	self.Dst.AddInEdge(self.Src)

	return self
}

//-----------------------------------------------------------
// Basic Blocks and Loops are being classified as regular, irreducible,
// and so on. This enum contains a symbolic name for all these classifications
//
const (
	_             = iota // Go has an interesting iota concept
	bbTop                // uninitialized
	bbNonHeader          // a regular BB
	bbReducible          // reducible loop
	bbSelf               // single BB loop
	bbIrreducible        // irreducible loop
	bbDead               // a dead BB
	bbLast               // sentinel
)

// UnionFindNode is used in the Union/Find algorithm to collapse
// complete loops into a single node. These nodes and the
// corresponding functionality are implemented with this class
//
type UnionFindNode struct {
	parent    *UnionFindNode
	bb        *BasicBlock
	loop      *SimpleLoop
	dfsNumber int
}

// Init explicitly initializes UnionFind nodes.
//
func (u *UnionFindNode) Init(bb *BasicBlock, dfsNumber int) {
	u.parent = u
	u.bb = bb
	u.dfsNumber = dfsNumber
	u.loop = nil
}

// FindSet implements the Find part of the Union/Find Algorithm
//
// Implemented with Path Compression (inner loops are only
// visited and collapsed once, however, deep nests would still
// result in significant traversals).
//
func (u *UnionFindNode) FindSet() *UnionFindNode {
	var nodeList []*UnionFindNode
	node := u

	for ; node != node.parent; node = node.parent {
		if node.parent != node.parent.parent {
			nodeList = append(nodeList, node)
		}

	}

	// Path Compression, all nodes' parents point to the 1st level parent.
	for _, ll := range nodeList {
		ll.parent = node.parent
	}

	return node
}

// Union relies on path compression.
//
func (u *UnionFindNode) Union(B *UnionFindNode) {
	u.parent = B
}

// Constants
//
// Marker for uninitialized nodes.
const unvisited = -1

// Safeguard against pathological algorithm behavior.
const maxNonBackPreds = 32 * 1024

// IsAncestor
//
// As described in the paper, determine whether a node 'w' is a
// "true" ancestor for node 'v'.
//
// Dominance can be tested quickly using a pre-order trick
// for depth-first spanning trees. This is why DFS is the first
// thing we run below.
//
// Go comment: Parameters can be written as w,v int, inlike in C, where
//   each parameter needs its own type.
//
func isAncestor(w, v int, last []int) bool {
	return ((w <= v) && (v <= last[w]))
}

// listContainsNode
//
// Check whether a list contains a specific element.
//
func listContainsNode(l []*UnionFindNode, u *UnionFindNode) bool {
	for _, ll := range l {
		if ll == u {
			return true
		}
	}
	return false
}

// DFS - Depth-First-Search and node numbering.
//
func DFS(currentNode *BasicBlock, nodes []*UnionFindNode, number []int, last []int, current int) int {
	nodes[current].Init(currentNode, current)
	number[currentNode.Name] = current

	lastid := current
	for _, target := range currentNode.OutEdges {
		if number[target.Name] == unvisited {
			lastid = DFS(target, nodes, number, last, lastid+1)
		}
	}
	last[number[currentNode.Name]] = lastid
	return lastid
}

func appendUnique(x []int, v int) []int {
	for _, i := range x {
		if i == v {
			return x
		}
	}
	x = append(x, v)
	return x
}

// FindLoops
//
// Find loops and build loop forest using Havlak's algorithm, which
// is derived from Tarjan. Variable names and step numbering has
// been chosen to be identical to the nomenclature in Havlak's
// paper (which, in turn, is similar to the one used by Tarjan).
//
func FindLoops(cfgraph *CFG, lsgraph *LSG) {
	if cfgraph.Start == nil {
		return
	}

	size := cfgraph.NumNodes()

	//nonBackPreds := make([]map[int]bool, size)
	nonBackPreds := make([][]int, size)
	backPreds := make([][]int, size)

	number := make([]int, size)
	header := make([]int, size, size)
	types := make([]int, size, size)
	last := make([]int, size, size)
	nodes := make([]*UnionFindNode, size, size)

	for i := 0; i < size; i++ {
		nodes[i] = new(UnionFindNode)
	}

	// Step a:
	//   - initialize all nodes as unvisited.
	//   - depth-first traversal and numbering.
	//   - unreached BB's are marked as dead.
	//
	for _, bb := range cfgraph.Blocks {
		number[bb.Name] = unvisited
		//nonBackPreds[i] = make(map[int]bool)
	}

	DFS(cfgraph.Start, nodes, number, last, 0)

	// Step b:
	//   - iterate over all nodes.
	//
	//   A backedge comes from a descendant in the DFS tree, and non-backedges
	//   from non-descendants (following Tarjan).
	//
	//   - check incoming edges 'v' and add them to either
	//     - the list of backedges (backPreds) or
	//     - the list of non-backedges (nonBackPreds)
	//
	for w := 0; w < size; w++ {
		header[w] = 0
		types[w] = bbNonHeader

		nodeW := nodes[w].bb
		if nodeW == nil {
			types[w] = bbDead
			continue // dead BB
		}

		if nodeW.NumPred() > 0 {
			for _, nodeV := range nodeW.InEdges {
				v := number[nodeV.Name]
				if v == unvisited {
					continue // dead node
				}

				if isAncestor(w, v, last) {
					backPreds[w] = append(backPreds[w], v)
				} else {
					//nonBackPreds[w][v] = true
					nonBackPreds[w] = appendUnique(nonBackPreds[w], v)
				}
			}
		}
	}

	// Start node is root of all other loops.
	header[0] = 0

	// Step c:
	//
	// The outer loop, unchanged from Tarjan. It does nothing except
	// for those nodes which are the destinations of backedges.
	// For a header node w, we chase backward from the sources of the
	// backedges adding nodes to the set P, representing the body of
	// the loop headed by w.
	//
	// By running through the nodes in reverse of the DFST preorder,
	// we ensure that inner loop headers will be processed before the
	// headers for surrounding loops.
	//
	for w := size - 1; w >= 0; w-- {
		// this is 'P' in Havlak's paper
		var nodePool []*UnionFindNode

		nodeW := nodes[w].bb
		if nodeW == nil {
			continue // dead BB
		}

		// Step d:
		for _, v := range backPreds[w] {
			if v != w {
				nodePool = append(nodePool, nodes[v].FindSet())
			} else {
				types[w] = bbSelf
			}
		}

		// Copy nodePool to workList.
		//
		workList := append([]*UnionFindNode(nil), nodePool...)

		if len(nodePool) != 0 {
			types[w] = bbReducible
		}

		// work the list...
		//
		for len(workList) > 0 {
			x := workList[0]
			workList = workList[1:]

			// Step e:
			//
			// Step e represents the main difference from Tarjan's method.
			// Chasing upwards from the sources of a node w's backedges. If
			// there is a node y' that is not a descendant of w, w is marked
			// the header of an irreducible loop, there is another entry
			// into this loop that avoids w.
			//

			// The algorithm has degenerated. Break and
			// return in this case.
			//
			nonBackSize := len(nonBackPreds[x.dfsNumber])
			if nonBackSize > maxNonBackPreds {
				return
			}

			for iter := range nonBackPreds[x.dfsNumber] {
				y := nodes[iter]
				ydash := y.FindSet()

				if !isAncestor(w, ydash.dfsNumber, last) {
					types[w] = bbIrreducible
					///nonBackPreds[w][ydash.dfsNumber] = true
					nonBackPreds[w] = appendUnique(nonBackPreds[w], ydash.dfsNumber)
				} else {
					if ydash.dfsNumber != w {
						if !listContainsNode(nodePool, ydash) {
							workList = append(workList, ydash)
							nodePool = append(nodePool, ydash)
						}
					}
				}
			}
		}

		// Collapse/Unionize nodes in a SCC to a single node
		// For every SCC found, create a loop descriptor and link it in.
		//
		if (len(nodePool) > 0) || (types[w] == bbSelf) {
			loop := lsgraph.NewLoop()

			loop.SetHeader(nodeW)
			if types[w] != bbIrreducible {
				loop.IsReducible = true
			}

			// At this point, one can set attributes to the loop, such as:
			//
			// the bottom node:
			//    iter  = backPreds[w].begin();
			//    loop bottom is: nodes[iter].node);
			//
			// the number of backedges:
			//    backPreds[w].size()
			//
			// whether this loop is reducible:
			//    type[w] != BasicBlockClass.bbIrreducible
			//
			nodes[w].loop = loop

			for _, node := range nodePool {
				// Add nodes to loop descriptor.
				header[node.dfsNumber] = w
				node.Union(nodes[w])

				// Nested loops are not added, but linked together.
				if node.loop != nil {
					node.loop.Parent = loop
				} else {
					loop.AddNode(node.bb)
				}
			}

			lsgraph.AddLoop(loop)
		} // nodePool.size
	} // Step c

}

// External entry point.
func FindHavlakLoops(cfgraph *CFG, lsgraph *LSG) int {
	FindLoops(cfgraph, lsgraph)
	return lsgraph.NumLoops()
}

//======================================================
// Scaffold Code
//======================================================

// Basic representation of loops, a loop has an entry point,
// one or more exit edges, a set of basic blocks, and potentially
// an outer loop - a "parent" loop.
//
// Furthermore, it can have any set of properties, e.g.,
// it can be an irreducible loop, have control flow, be
// a candidate for transformations, and what not.
//
type SimpleLoop struct {
	// No set, use map to bool
	basicBlocks map[*BasicBlock]bool
	Children    map[*SimpleLoop]bool
	Parent      *SimpleLoop
	header      *BasicBlock

	IsRoot       bool
	IsReducible  bool
	Counter      int
	NestingLevel int
	DepthLevel   int
}

func (loop *SimpleLoop) AddNode(bb *BasicBlock) {
	loop.basicBlocks[bb] = true
}

func (loop *SimpleLoop) AddChildLoop(child *SimpleLoop) {
	loop.Children[child] = true
}

func (loop *SimpleLoop) Dump(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Printf("  ")
	}

	// No ? operator ?
	fmt.Printf("loop-%d nest: %d depth %d ",
		loop.Counter, loop.NestingLevel, loop.DepthLevel)
	if !loop.IsReducible {
		fmt.Printf("(Irreducible) ")
	}

	// must have > 0
	if len(loop.Children) > 0 {
		fmt.Printf("Children: ")
		for ll := range loop.Children {
			fmt.Printf("loop-%d", ll.Counter)
		}
	}
	if len(loop.basicBlocks) > 0 {
		fmt.Printf("(")
		for bb := range loop.basicBlocks {
			fmt.Printf("BB#%06d ", bb.Name)
			if loop.header == bb {
				fmt.Printf("*")
			}
		}
		fmt.Printf("\b)")
	}
	fmt.Printf("\n")
}

func (loop *SimpleLoop) SetParent(parent *SimpleLoop) {
	loop.Parent = parent
	loop.Parent.AddChildLoop(loop)
}

func (loop *SimpleLoop) SetHeader(bb *BasicBlock) {
	loop.AddNode(bb)
	loop.header = bb
}

//------------------------------------
// Helper (No templates or such)
//
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// LoopStructureGraph
//
// Maintain loop structure for a given CFG.
//
// Two values are maintained for this loop graph, depth, and nesting level.
// For example:
//
// loop        nesting level    depth
//----------------------------------------
// loop-0      2                0
//   loop-1    1                1
//   loop-3    1                1
//     loop-2  0                2
//
var loopCounter = 0

type LSG struct {
	root  *SimpleLoop
	loops []*SimpleLoop
}

func NewLSG() *LSG {
	lsg := new(LSG)
	lsg.root = lsg.NewLoop()
	lsg.root.NestingLevel = 0

	return lsg
}

func (lsg *LSG) NewLoop() *SimpleLoop {
	loop := new(SimpleLoop)
	loop.basicBlocks = make(map[*BasicBlock]bool)
	loop.Children = make(map[*SimpleLoop]bool)
	loop.Parent = nil
	loop.header = nil

	loop.Counter = loopCounter
	loopCounter++
	return loop
}

func (lsg *LSG) AddLoop(loop *SimpleLoop) {
	lsg.loops = append(lsg.loops, loop)
}

func (lsg *LSG) Dump() {
	lsg.dump(lsg.root, 0)
}

func (lsg *LSG) dump(loop *SimpleLoop, indent int) {
	loop.Dump(indent)

	for ll := range loop.Children {
		lsg.dump(ll, indent+1)
	}
}

func (lsg *LSG) CalculateNestingLevel() {
	for _, sl := range lsg.loops {
		if sl.IsRoot {
			continue
		}
		if sl.Parent == nil {
			sl.SetParent(lsg.root)
		}
	}
	lsg.calculateNestingLevel(lsg.root, 0)
}

func (lsg *LSG) calculateNestingLevel(loop *SimpleLoop, depth int) {
	loop.DepthLevel = depth
	for ll := range loop.Children {
		lsg.calculateNestingLevel(ll, depth+1)

		ll.NestingLevel = max(loop.NestingLevel, ll.NestingLevel+1)
	}
}

func (lsg *LSG) NumLoops() int {
	return len(lsg.loops)
}

func (lsg *LSG) Root() *SimpleLoop {
	return lsg.root
}

//======================================================
// Testing Code
//======================================================

func buildDiamond(cfgraph *CFG, start int) int {
	bb0 := start
	NewBasicBlockEdge(cfgraph, bb0, bb0+1)
	NewBasicBlockEdge(cfgraph, bb0, bb0+2)
	NewBasicBlockEdge(cfgraph, bb0+1, bb0+3)
	NewBasicBlockEdge(cfgraph, bb0+2, bb0+3)

	return bb0 + 3
}

func buildConnect(cfgraph *CFG, start int, end int) {
	NewBasicBlockEdge(cfgraph, start, end)
}

func buildStraight(cfgraph *CFG, start int, n int) int {
	for i := 0; i < n; i++ {
		buildConnect(cfgraph, start+i, start+i+1)
	}
	return start + n
}

func buildBaseLoop(cfgraph *CFG, from int) int {
	header := buildStraight(cfgraph, from, 1)
	diamond1 := buildDiamond(cfgraph, header)
	d11 := buildStraight(cfgraph, diamond1, 1)
	diamond2 := buildDiamond(cfgraph, d11)
	footer := buildStraight(cfgraph, diamond2, 1)
	buildConnect(cfgraph, diamond2, d11)
	buildConnect(cfgraph, diamond1, header)

	buildConnect(cfgraph, footer, from)
	footer = buildStraight(cfgraph, footer, 1)
	return footer
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to this file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	lsgraph := NewLSG()
	cfgraph := NewCFG()

	cfgraph.CreateNode(0) // top
	cfgraph.CreateNode(1) // bottom
	NewBasicBlockEdge(cfgraph, 0, 2)

	for dummyloop := 0; dummyloop < 15000; dummyloop++ {
		FindHavlakLoops(cfgraph, NewLSG())
	}

	n := 2

	for parlooptrees := 0; parlooptrees < 10; parlooptrees++ {
		cfgraph.CreateNode(n + 1)
		buildConnect(cfgraph, 2, n+1)
		n = n + 1

		for i := 0; i < 100; i++ {
			top := n
			n = buildStraight(cfgraph, n, 1)
			for j := 0; j < 25; j++ {
				n = buildBaseLoop(cfgraph, n)
			}
			bottom := buildStraight(cfgraph, n, 1)
			buildConnect(cfgraph, n, top)
			n = bottom
		}
		buildConnect(cfgraph, n, 1)
	}

	FindHavlakLoops(cfgraph, lsgraph)
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}
