package main

import (
	"testing"
)

func BenchmarkGetWhitelistType(b *testing.B) {
	var ret string
	for i := 0; i < b.N; i++ {
		ret = getWhiteListType("ILOVE WHKEK DSKFSKFDKFSDFSSDFKIE")
	}

	if ret == WhiteListTypeIP {
		b.Logf("ret:%s", ret)
	}
}
