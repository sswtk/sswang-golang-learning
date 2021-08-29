package main

import "sswang-golang-learning/go_learning_imook/tree"

func main(){
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	//使用工厂函数创建
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()
	// 省略上述Node的写法
	//nodes := []Node{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	//root.print()
	//fmt.Println()

	//root.right.left.setValue(4)
	//root.right.left.print()
	//fmt.Println()

	//root.print()
	//root.setValue(100)

	//var pRoot *Node
	//pRoot.setValue(200)
	//pRoot = &root
	//pRoot.setValue(300)
	//pRoot.print()
}
