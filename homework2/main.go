package main

import (
	"fmt"
	"time"
)

func main() {

	intPara := 20
	Add10(&intPara)
	fmt.Println(intPara)

	slice1 := []int{1, 2, 3, 4, 5}
	MutiplyBy2(&slice1)
	fmt.Println("方法1结果:", slice1) // 输出: [2 4 6 8 10]

	PrintEvenAndOddNum()

	operations := []Operation{}
	operations = append(operations, func() {
		fmt.Println("方法1")
		time.Sleep(1 * time.Second) // 重试延迟
	})
	operations = append(operations, func() {
		fmt.Println("方法2")
		time.Sleep(2 * time.Second) // 重试延迟
	})

	Run(operations)
}
