package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

//给结构体定义方法； (node Node)是方法的接收者
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}


// 使用指针作为方法接收者(nil指针也可以）
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.Value = value
}


//创建结构体的方式：Factory function -> 工厂函数
func CreateNode(value int) *Node{
	return &Node{Value: value}
}
