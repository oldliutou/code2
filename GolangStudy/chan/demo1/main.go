package main

import "fmt"

func main() {
	c := make(chan int, 3)
	fmt.Println("len(c)=", len(c), " cap(c)= ", cap(c))
	go func() {
		defer fmt.Println("goroutine exit")
		for i := 0; i < 14; i++ {
			c <- i
			fmt.Println("send:", i)
		}
		close(c)
	}()
	i := 0
	for {
		if data, ok := <-c; ok {
			fmt.Println("recv:", data)
		}
		i++
		fmt.Println("i=", i)

	}
	fmt.Println("Main Finshed!!!")

}
