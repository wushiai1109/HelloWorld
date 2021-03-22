package main

import (
	//_ "./lib1"
	"./lib1"
	mylib2 "./lib2"
)

func main() {

	//import _ “fmt”
	//给fmt包起一个别名，匿名， ⽆法使用当前包的方法，但是会执⾏当前的包内部的init()方法

	//import aa “fmt”
	//给fmt包起一个别名，aa， aa.Println()来直接调用。

	//import . “fmt” (不建议)
	//将当前fmt包中的全部⽅方法，导⼊到当前本包的作用中，fmt包中 的全部的⽅方法可以直接使⽤API来调用，不需要fmt.API来调⽤

	lib1.Lib1Test()

	//lib2.Lib2Test()
	mylib2.Lib2Test()
}
