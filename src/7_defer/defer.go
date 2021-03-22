package main

import "fmt"

func main() {
	////写入defer关键字
	//defer fmt.Println("main end1")
	//defer fmt.Println("main end2")
	//
	//fmt.Println("main::hello go 1")
	//fmt.Println("main::hello go 2")
	//
	//defer func1()
	//defer func2()
	//defer func3()

	returnAndDefer()

}

//知识点一：defer的执行顺序
func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

//知识点二：return之后的语句先执行，defer后的语句后执行
func deferFunc() int {
	fmt.Println("defer func~~~")
	return 0
}

func returnFunc() int {
	fmt.Println("return func~~~")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}
