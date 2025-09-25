package main

import "fmt"

func Hello() {
	fmt.Println("Hello world !")
}

func Hello2() {
	fmt.Println("Hello world")
}

func Number() {
	fmt.Println("123")
}

func Number2() {
	fmt.Println("456")
}

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
	Hello()
	Hello2()
	Number()
	Number2()
	fmt.Println("Hello world from a!")
	test1()
	test2()
	test3()
}
