package test

import "fmt"

/**
 * @Author: Bruce
 * @Date: 2021/1/16 9:16 下午
 * @Description:
 */
func main1() {
	//2_var a int = 1
	//2_var b int = 1
	//2_var (
	//	a int = 1
	//	b int = 1
	//)
	//2_var (
	//	a = 1
	//	b = 1
	//)
	a := 1
	b := 1
	//fmt.Println(a)
	for i := 0; i < 5; i++ {
		//fmt.Print(" ", b)
		fmt.Println(" ", a)
		tmp := a
		a = b
		b = tmp + a
	}
}
