package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	child *bool
)

func startChild(file *os.File) {

	args := []string{"-child"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// put socket FD at the first entry
	cmd.ExtraFiles = []*os.File{file}
	err := cmd.Start()
	if err != nil {
		fmt.Printf("start child failed, err:%v\n", err)
		return
	}
}

func init() {
	child = flag.Bool("child", false, "继承于父进程(internal use only)")
}

func readFromParent() {
	f := os.NewFile(3, "")
	count := 0
	for {
		str := fmt.Sprintf("hello, i'child process, write:%d line\n", count)
		count++
		f.WriteString(str)
		time.Sleep(time.Second)
	}
}

func main() {
	if child != nil && *child == true {
		fmt.Printf("继承于父进程的文件句柄\n")
		readFromParent()
		return
	}
	//父进程的逻辑
	file, err := os.OpenFile("c:/tmp/test_inherit.log", os.O_APPEND|os.O_CREATE, 0766)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	startChild(file)
	fmt.Printf("parant exited")
}
