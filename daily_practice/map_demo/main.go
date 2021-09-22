package main

import "fmt"

func main() {
	//常规创建
	m1 := map[string]string{
		"m1": "v1",
	}
	fmt.Printf("%v\n", m1) //map[m1:v1]

	//使用make函数创建，make可以创建slice/map
	m2 := make(map[string]string)
	m2["m2"] = "v2"
	fmt.Printf("%v\n", m2) //map[m2:v2]

	//定义一个空map
	m3 := map[string]string{}
	m4 := make(map[string]string)
	fmt.Printf("%v\n", m3) //map[]
	fmt.Printf("%v\n", m4) // map[]

	// map中key的类型
	// key为切片类型---不支持
	//a1 := []int{1, 2, 3}
	//b1 := []int{1, 2, 3}
	//if a1 == b1 {  //编译不通过
	//}

	// key 为数组类型（因为数组是值类型）--- 支持
	a1 := [3]int{1, 2, 3}
	b1 := [3]int{1, 2, 3}
	if a1 == b1 { //编译通过
		fmt.Println("true")
	}

	//interface{}类型可以作为key，但是需要加入的key的类型是可以比较的
	//var m5 map[interface{}]string
	//m6 = make(map[interface{}]string)
	//m5[[]byte("k2")] = "v2" // panic: runtime error: hash of unhashable type []uint8
	//m5[123] = "123"
	//m6[12.3] = "123"
	//fmt.Println(m5)
	//fmt.Println(m6)

	// book里面的元素都是支持== !=
	type book struct {
		name string
	}
	var m7 map[book]string
	fmt.Println(m7) //map[]

	// map的增删改
	//创建
	mm := map[string]string{
		"a": "va",
		"b": "vb",
	}
	fmt.Println(len(mm)) //2

	// 增加、修改
	// k不存在为增加，k存在为修改
	mm["c"] = ""
	mm["c"] = "11"                       // 重复增加(key相同)，使用新的值覆盖
	fmt.Printf("%#v %#v\n", mm, len(mm)) // map[string]string{"a":"va", "b":"vb", "c":"11"} 3

	// 删除
	// delete(m, k) 将k以及k对应的v从m中删掉；如果k不在m中，不执行任何操作
	delete(mm, "x")                      // 删除不存在的key,原m不影响
	delete(mm, "a")                      // 删除存在的key
	fmt.Printf("%#v %#v\n", mm, len(mm)) // map[string]string{"b":"vb", "c":"11"} 2
	delete(mm, "a")                      // 重复删除不报错,m无影响
	fmt.Printf("%#v %#v\n", mm, len(mm)) /// map[string]string{"b":"vb", "c":"11"} 2

	// 查找
	// v := m[k] // 从m中取键k对应的值给v，如果k在m中不存在，则将value类型的零值赋值给v
	// v, ok := m[k] // 从m中取键k对应的值给v，如果k存在，ok=true,如果k不存在，将value类型的零值赋值给v同时ok=false
	// 查1 - 元素不存在
	v1 := mm["x"] //
	v2, ok2 := mm["x"]
	fmt.Printf("%#v %#v %#v\n", v1, v2, ok2) // "" "" false

	// 查2 - 元素存在
	v3 := mm["a"]
	v4, ok4 := mm["a"]
	fmt.Printf("%#v %#v %#v\n", v3, v4, ok4) //"va" "va" true

	//遍历--遍历顺序随机，k,v使用的同一块内存
	ma := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range ma {
		fmt.Printf("k:[%v]. v:[%v]\n", k, v)
		//k:[a]. v:[1]
		//k:[b]. v:[2]
		//k:[c]. v:[3]
	}
}
