package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	var res = 0
	for i := 0; i < 1000; i++ {
		cmd := exec.Command("python", "-c ")
		start := time.Now().UnixNano()
		_ = cmd.Run()
		end := time.Now().UnixNano()
		res += int(end - start)
	}
	fmt.Println(res / 1000)
}
