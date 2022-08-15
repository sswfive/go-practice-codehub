package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 为结构体定义方法
func (node Node) Print() {
	fmt.Println(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil " + "node. Ignored")
		return
	}
	node.Value = value
}

// 工厂函数
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
