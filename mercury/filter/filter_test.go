package filter

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {

	//加载敏感词库
	err := Init("../data/filter.dat.txt")
	if err != nil {
		t.Errorf("load filter data failed, err:%v", err)
		return
	}

	data := `乱伦`
	result, hit := Replace(data, "***")
	fmt.Printf("hit:%#v, str:%v\n", hit, result)

}
