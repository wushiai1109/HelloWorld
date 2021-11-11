package main

import (
	"fmt"
	"strings"
)

//strings.Split()有个坑
//按照指定的分隔符拆分字符串
//strings.Split(str, "分隔符")
//比如我现在有个字符串：
//str := "Zk;5cp;"
//当用 ; 去拆分后，你会发现拆分后的数组len() 为3, 而不是2.
func main() {
	str := ";Zk;5cp;"
	fmt.Println(str[0:1])
	if str[0:1] == ";" {
		str = str[1:]
	}
	if str[len(str)-1:] == ";" {
		str = str[:len(str)-1]
	}
	data := strings.Split(str, ";")
	fmt.Println(data)
	fmt.Println(data[0])
	fmt.Println(data[1])
	fmt.Println(len(data))
}
