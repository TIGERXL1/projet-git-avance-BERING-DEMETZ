package main

import "fmt"

func test1() {
	fmt.Println("This is a first test")
}

func test2() {
	fmt.Println("This is a second test")
}

func test3() {
	fmt.Println("This is a third test")
}

func main() {
	fmt.Println("Hello world from a!")
	test1()
	test2()
	test3()
}
