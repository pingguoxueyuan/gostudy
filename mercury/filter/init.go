package filter

import (
	"bufio"
	"io"
	"os"

	"github.com/pingguoxueyuan/gostudy/mercury/util"
)

var (
	trie *util.Trie
)

func Init(filename string) (err error) {

	trie = util.NewTrie()
	file, err := os.Open(filename)
	if err != nil {
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		word, errRet := reader.ReadString('\n')
		if errRet == io.EOF {
			return
		}
		if errRet != nil {
			err = errRet
			return
		}

		err = trie.Add(word, nil)
		if err != nil {
			return
		}
	}
	return
}
