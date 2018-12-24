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

	data := `
	这种婚外情心理很常见，比例开始逐渐增多。一般出轨的男人大多都是这种想法，当然，也有少部分女人也是这么想的。他们左手家庭右手情人，哪个都不舍得扔！一边嘴里夸着爱人，一边心里想着情人。一边炫耀夫妻恩爱，一边又觉得自己有魅力。家里和和美美，家外又有情人对自己服服帖帖。
	
	这种婚外情一般发生在熟人之间，以索取为目的。一般发生的主体是“一次恋爱就结婚”的夫妻，出轨对象大多是同事或者上司。在没彻底做出抉择前，多数人选择过一天是一天，一石二鸟一箭双雕！至于将来和情人能发展到什么程度，他们一般不会去深刻考虑。大不了到时候让情人出局就是了，趁着现在还能享受，哪还管那么多。
	裸体，喜欢，小黄片，乱伦
	
	`
	result, isReplace := Replace(data, "***")
	fmt.Printf("isReplace:%#v, str:%v\n", isReplace, result)

}
