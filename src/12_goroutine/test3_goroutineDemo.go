package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
 * @Author: Bruce
 * @Date: 2021/3/14 12:39 下午
 * @Description: 并不是顺序执行，而是 3，1，2 这样，比如 4 个协程，就会是 4，1，2，3   ???(未知)
 */
func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(3)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			fmt.Print("@")
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}
