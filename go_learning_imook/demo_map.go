package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) // empty map

	var m3 map[string]int // nil

	fmt.Println(m, m2, m3)

	fmt.Println("---Traversing map----") // 遍历(key无序)
	for k, v := range m {
		fmt.Println(k, v)
	}
	//for k := range m {  // 遍历key
	//	fmt.Println(k)
	//}
	//for _, v := range m {  // 遍历value
	//	fmt.Println(v)
	//}

	fmt.Println("----Getting values----")
	courseName := m["course"]
	fmt.Println(courseName)  // 取一个存在的key, return value

	coursName := m["cours"]
	fmt.Println(coursName)  // 取一个不存在的key, return empty

	courseName1, ok:= m["course"]
	fmt.Println(courseName1, ok)  // ok为true

	coursName1, ok := m["cours"]
	fmt.Println(coursName1, ok)  // ok为false

	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName)
	}else{
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values-----")  // map delete
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name,ok = m["name"]
	fmt.Println(name, ok)

}



