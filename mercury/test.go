package main

import (
	"fmt"

	"github.com/sony/sonyflake"
)

func main() {
	//生产环境一定要设置machineID，使用zk或者etcd
	st := sonyflake.Settings{}
	sk := sonyflake.NewSonyflake(st)
	fmt.Println(sk.NextID())
}
