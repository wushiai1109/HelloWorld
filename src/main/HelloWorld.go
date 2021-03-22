package main

import (
	"fmt"
	"os"
)

//func main() {
//	fmt.Println("Hello World!")
//}

//func main() {
//	fmt.Println("Hello World!")
//	os.Exit(0)
//}

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello World!", os.Args[1])
	}
	fmt.Println(Fib(9))
	change(8, 9)
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func change(m, n int) {
	var a = 1
	var b = 2
	a, b = b, a
	fmt.Println(a, b)
	m, n = n, m
	fmt.Println(m, n)
}
