package main

import (
	"fmt"

	"github.com/sony/sonyflake"
)

func main() {
	settings := sonyflake.Settings{}
	sk := sonyflake.NewSonyflake(settings)

	fmt.Println(sk.NextID())
}
