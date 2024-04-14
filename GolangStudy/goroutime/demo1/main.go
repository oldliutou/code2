package main

import (
	"fmt"
	"time"
)

func main() {
	//新建goroutine
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	// 死循环
	// 如果不加此循环，main goroutine结束，则A,B的defer语句不会执行
	i := 0
	for {
		i++
		fmt.Println("main i= ", i)
		time.Sleep(1 * time.Second)
	}
}
