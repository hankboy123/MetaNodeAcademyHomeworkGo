package main

import (
	"fmt"
)

func main() {

	intPara := 20
	Add10(&intPara)
	fmt.Println(intPara)

	slice1 := []int{1, 2, 3, 4, 5}
	mutiplyBy2(&slice1)
	fmt.Println("方法1结果:", slice1) // 输出: [2 4 6 8 10]
}
