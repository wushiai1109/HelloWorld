package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

//func main() {
//	var res = 0
//	for i := 0; i < 1000; i++ {
//		cmd := exec.Command("./b.out", "")
//		start := time.Now().UnixNano()
//		_ = cmd.Run()
//		//if err := cmd.Run(); err != nil {
//		//	fmt.Printf("failed to run command: %s\n", err)
//		//}
//		end := time.Now().UnixNano()
//		res += int(end - start)
//	}
//	fmt.Println(res / 1000)
//}

func main() {

	var res = 0
	for i := 0; i < 100; i++ {
		cmd := exec.Command("python", "./test.py")
		//cmd.Stdin = os.Stdin
		//cmd.Stdout = os.Stdout
		//cmd.Stderr = os.Stderr
		start := time.Now().UnixNano()
		//_ = cmd.Start()
		//_ = cmd.Wait()
		//_ = cmd.Run()
		if err := cmd.Run(); err != nil {
			fmt.Printf("failed to run command: %s\n", err)
		}
		end := time.Now().UnixNano()
		res += int(end - start)
	}
	fmt.Println(res / 100)

	//if err := cmd.Start(); err != nil {
	//	fmt.Printf("failed to run command: %s\n", err)
	//}
	//if err := cmd.Wait(); err != nil {
	//	fmt.Printf("failed to wait command: %s\n", err)
	//}
}

//func main() {
//cmd := exec.Command("tr", "a-z", "A-Z")
//cmd.Stdin = strings.NewReader("some input")
//var out bytes.Buffer
//cmd.Stdout = &out
//start := time.Now().UnixNano()
//err := cmd.Run()
//end := time.Now().UnixNano()
//fmt.Println(end - start)
//if err != nil {
//	log.Fatal(err)
//}
//fmt.Printf("in all caps: %q\n", out.String())
//}

type MyReader struct {
	file *os.File
}

func (r *MyReader) Read(b []byte) (n int, err error) {
	reader := bufio.NewReader(r.file)
	n, err = reader.Read(b)
	return n, err
}

//func main7() {
//	var res = 0
//	for i := 0; i < 100; i++ {
//		file, _ := os.Open("./test1.txt")
//		r := MyReader{file: file}
//		var b []byte
//		_, _ = r.Read(b)
//		cmd := exec.Command("./a.out", "")
//		cmd.Stdin = strings.NewReader(string(b))
//		//var out bytes.Buffer
//		//cmd.Stdout = &out
//		start := time.Now().UnixNano()
//		_ = cmd.Run()
//		//if err := cmd.Run(); err != nil {
//		//	fmt.Printf("failed to run command: %s\n", err)
//		//}
//		end := time.Now().UnixNano()
//		fmt.Println(end - start)
//		//fmt.Printf("file length: %d\n", len(out.String()))
//		res += int(end - start)
//	}
//	fmt.Println(res / 100)
//}

func main4() {
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
}

func main1() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}

func main7() {
	//var res = 0
	//for i := 0; i < 10; i++ {
	//cmd := exec.Command("tee", "out.txt")
	//file, _ := os.Open("./test.txt")
	//buf, _ := ioutil.ReadAll(file)
	//stdin := strings.NewReader(string(buf))
	cmd := exec.Command("./a.out", "")
	//cmd.Stdin = stdin
	var out bytes.Buffer
	cmd.Stdout = &out
	start := time.Now().UnixNano()
	if err := cmd.Run(); err != nil {
		fmt.Printf("failed to run command: %s\n", err)
	}
	end := time.Now().UnixNano()
	fmt.Println(end - start)
	fmt.Printf("%s", out.String())
	//res += int(end - start)
	//}
	//fmt.Println(res / 10)
}

//c文件中生成一个文件，在内容中输出
func main6() {
	//var res = 0
	//for i := 0; i < 10; i++ {
	cmd := exec.Command("./b.out", "")
	var out bytes.Buffer
	cmd.Stdout = io.Writer(&out)
	start := time.Now().UnixNano()
	if err := cmd.Run(); err != nil {
		fmt.Printf("failed to run command: %s\n", err)
	}
	end := time.Now().UnixNano()
	fmt.Println(end - start)
	//fmt.Printf("%s", out.String())
	//res += int(end - start)
	//}
	//fmt.Println(res / 10)
}

func main9() {
	file, _ := os.Open("./test10.txt")
	buf, _ := ioutil.ReadAll(file)
	cmd := exec.Command("echo", "-n", string(buf))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now().UnixNano()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	out := make([]byte, 0)
	if _, err = stdout.Read(out); err != nil {
		fmt.Printf("failed %s\n", err)
	}
	//var person struct {
	//	Name string
	//	Age  int
	//}
	//if err := json.NewDecoder(stdout).Decode(&person); err != nil {
	//	log.Fatal(err)
	//}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	end := time.Now().UnixNano()
	fmt.Println(end - start)
	//fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}

//func main() {
//	cmd := exec.Command("tr", "a-z", "A-Z")
//	file, _ := os.Open("./test1.txt")
//	buf, _ := ioutil.ReadAll(file)
//	cmd.Stdin = strings.NewReader(string(buf))
//	var out bytes.Buffer
//	cmd.Stdout = &out
//	err := cmd.Run()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("in all caps: %q\n", out.String())
//}

//func main() {
//	cmd := exec.Command("tr", "a-z", "A-Z")
//	file, _ := os.Open("./test10.txt")
//	buf, _ := ioutil.ReadAll(file)
//	cmd.Stdin = strings.NewReader(string(buf))
//	var out bytes.Buffer
//	cmd.Stdout = io.Writer(&out)
//	start := time.Now().UnixNano()
//	err := cmd.Run()
//	if err != nil {
//		log.Fatal(err)
//	}
//	end := time.Now().UnixNano()
//	fmt.Println(end - start)
//	//fmt.Printf("in all caps: %q\n", out.String())
//}
