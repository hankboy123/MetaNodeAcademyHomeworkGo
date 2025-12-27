package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个函数类型
type Operation func()

func Add10(pointer *int) {
	*pointer = *pointer + 10
}

func MutiplyBy2(pointer *[]int) {
	for i := range *pointer {
		(*pointer)[i] = (*pointer)[i] * 2
	}
}

func PrintOddNum(limit int) {

	fmt.Println("奇数:")
	for i := 0; i <= limit; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func PrintEvenNum(limit int) {

	fmt.Println("偶数")
	for i := 0; i <= limit; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func PrintEvenAndOddNum() {
	// 使用 WaitGroup 等待两个协程完成
	var wg sync.WaitGroup
	wg.Add(2)

	// 协程1：打印奇数
	go func() {
		defer wg.Done()
		PrintOddNum(10)
	}()

	// 协程2：打印偶数
	go func() {
		defer wg.Done()
		PrintEvenNum(10)
	}()

	// 等待两个协程完成
	wg.Wait()
	fmt.Println("所有数字打印完成")
}

func PrintEvenAndOddNumV2() {
	// 创建两个通道用于协程间通信
	oddChan := make(chan int)
	evenChan := make(chan int)

	// 协程1：生成奇数
	go func() {
		for i := 1; i <= 10; i += 2 {
			oddChan <- i
		}
		close(oddChan)
	}()

	// 协程2：生成偶数
	go func() {
		for i := 2; i <= 10; i += 2 {
			evenChan <- i
		}
		close(evenChan)
	}()

	// 从通道接收并打印（先打印奇数，再打印偶数）
	for num := range oddChan {
		fmt.Printf("奇数: %d\n", num)
	}

	for num := range evenChan {
		fmt.Printf("偶数: %d\n", num)
	}

	fmt.Println("所有数字打印完成")
}

func Run(operations []Operation) {
	if len(operations) == 0 {
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(operations))

	for _, oper := range operations {
		go func() {
			startTime := time.Now()
			oper()
			duration := time.Since(startTime)
			defer wg.Done()
			fmt.Println("耗时", duration)
		}()

	}
	// 等待两个协程完成
	wg.Wait()
	fmt.Println("所有协程打印完成")
}

func Send(ch chan int, quit chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- 1
}

func Receive(ch chan int, quit chan int) {
	for true {
		select {
		case <-quit:
			fmt.Println("接收完成")
			return
		case num := <-ch:
			fmt.Println("receive:", num)
		}
	}

}
