package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	var res = 0
	for i := 0; i < 1000; i++ {
		start := time.Now().UnixNano()
		cmd := exec.Command("python", "./test.py")
		end := time.Now().UnixNano()
		_ = cmd.Run()
		res += int(end - start)
	}
	fmt.Println(res / 1000)
}
