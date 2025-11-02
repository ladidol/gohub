package main

import (
	"fmt"
	"math/rand"

	. "github.com/bytedance/mockey"
)

func main() {
	Mock(rand.Int).Return(1).Build() // mock `rand.Int` 返回 1

	fmt.Printf("rand.Int() 总是返回: %v\n", rand.Int())
}
