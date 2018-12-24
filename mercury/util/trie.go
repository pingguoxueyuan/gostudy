package util

type Node struct {
	//rune表示一个utf8字符
	char   rune
	Data   interface{}
	parent *Node
	Depth  int
	//childs 用来当前节点的所有孩子节点
	childs map[rune]*Node
	term   bool
}

type Trie struct {
	root *Node
	size int
}

func NewNode() *Node {
	return &Node{
		childs: make(map[rune]*Node, 32),
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

func (p *Trie) Add(key string, data interface{}) (err error) {

	node := p.root
	runes := []rune(key)
	for _, r := range runes {
		ret, ok := node.childs[r]
		if !ok {
			ret = NewNode()
			ret.Depth = node.Depth + 1
			ret.char = r
			node.childs[r] = ret
		}

		node = ret
	}

	node.term = true
	node.Data = data
	return
}

func (p *Trie) findNode(key string) (result *Node) {

	node := p.root
	chars := []rune(key)
	for _, v := range chars {
		ret, ok := node.childs[v]
		if !ok {
			return
		}

		node = ret
	}

	result = node
	return
}

func (p *Trie) collectNode(node *Node) (result []*Node) {

	if node == nil {
		return
	}

	if node.term {
		result = append(result, node)
		return
	}

	var queue []*Node
	queue = append(queue, node)

	for i := 0; i < len(queue); i++ {
		if queue[i].term {
			result = append(result, queue[i])
			continue
		}

		for _, v1 := range queue[i].childs {
			queue = append(queue, v1)
		}
	}

	return
}

func (p *Trie) PrefixSearch(key string) (result []*Node) {

	node := p.findNode(key)
	if node == nil {
		return
	}

	result = p.collectNode(node)
	return
}

func (p *Trie) Check(text, replace string) (result bool, str string) {

	chars := []rune(text)
	if p.root == nil {
		return
	}

	var left []rune
	node := p.root
	start := 0
	for index, v := range chars {
		ret, ok := node.childs[v]
		if !ok {
			left = append(left, chars[start:index+1]...)
			start = index + 1
			node = p.root
			continue
		}

		node = ret
		if ret.term {
			result = true
			node = p.root
			left = append(left, ([]rune(replace))...)
			start = index + 1
			continue
		}
	}

	str = string(left)
	return
}
