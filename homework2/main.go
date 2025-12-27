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

	c := Circle{Radius: 5}
	var s Shape = c // 接口变量可以存储实现了接口的类型
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())

	p := Employee{Person: Person{Age: 30, Name: "Alice"}, EmployeeID: "E123"}
	p.PrintInfo()

	ch := make(chan int)
	quit := make(chan int)

	go Send(ch, quit)
	Receive(ch, quit)

	ch1 := make(chan int, 100)
	quit1 := make(chan int)

	go Send(ch1, quit1)
	Receive(ch1, quit1)
}
