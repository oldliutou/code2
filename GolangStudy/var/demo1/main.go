package main

import "fmt"

type Book struct {
	Title string
}

func main() {
	fmt.Println("hello world")
	fmt.Println("---------------------------------------")
	var i1 int
	fmt.Printf("i=%d\n", i1)
	var i2 int = 10
	fmt.Printf("i=%d\n", i2)
	var i3 = 100
	fmt.Printf("i=%d\n", i3)
	i4 := 1000
	fmt.Printf("i=%d\n", i4)
	fmt.Println("---------------------------------------")
	b := &Book{Title: "Golang"}
	fmt.Printf("b type is %T", b)
}
