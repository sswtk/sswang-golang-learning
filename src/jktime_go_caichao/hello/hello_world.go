package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}
	// fmt.Println("hello World")
	// os.Exit(1) 正常退出
	// os.Exit(-1) 异常退出
}
