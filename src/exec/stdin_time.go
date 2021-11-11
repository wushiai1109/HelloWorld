package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	var res = 0
	for i := 0; i < 100; i++ {
		file, _ := os.Open("./test10.txt")
		buf, _ := ioutil.ReadAll(file)
		cmd := exec.Command("./a.out", "")
		cmd.Stdin = strings.NewReader(string(buf))
		start := time.Now().UnixNano()
		if err := cmd.Run(); err != nil {
			fmt.Printf("failed to run command: %s\n", err)
		}
		end := time.Now().UnixNano()
		fmt.Println(end - start)
		res += int(end - start)
	}
	fmt.Println(res / 100)

	fmt.Println(1024*1024)
}
