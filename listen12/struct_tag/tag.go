package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string  `json:"xxxx"`
	Sex      string  `json:"sex"`
	Score    float32 `json:"score"`
	Age      int32   `json:"age"`
}

func main() {
	user := &User{
		UserName: "user01",
		Sex:      "男",
		Score:    99.2,
		Age:      18,
	}

	data, _ := json.Marshal(user)
	fmt.Printf("json str:%s\n", string(data))
}
