package tree


// 中树遍历（先左再右）
func (node *Node) Traverse() {
	if node == nil{
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
