package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//func bounded(v, int) int {
//	if v > 100 {
//		return 100
//	}else if v < 0 {
//		return 0
//	} else {
//		return v
//	}
//}

func grade(score int) string {
	g :=""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong scorce:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}


func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string ) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

//死循环
func forever(){
	for {
		fmt.Println("abab")
	}
}

func main() {
	//forever()
	//printFile("go_learning_imook/abc.txt")
	//fmt.Println(
	//	convertToBin(5), // 101
	//	convertToBin(13), // 1101
	//)
	//fmt.Println(grade(85))
	//fmt.Println(grade(73))
	//fmt.Println(grade(58))
	//fmt.Println(grade(101))

	//const filename = "abc.txt"
	// 方式一
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}
	// 方式二
	//if contents, err := ioutil.ReadFile(filename); err != nil {
	//	fmt.Println(err)
	//} else{
	//	fmt.Printf("%s\n", contents)
	//}
}


