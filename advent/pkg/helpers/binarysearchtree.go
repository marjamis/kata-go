package helpers

type Node struct {
	Name   string
	Data   int
	Parent *Node
	Left   *Node
	Right  *Node
}

func CreateBinarySearchTree() *Node {
	return &Node{
		"Root Node",
		0,
		nil,
		nil,
		nil,
	}
}

func InsertNode(root *Node, data int) {

	// parent := findParentNode(root, data)
	//
	// node := &Node{
	// 	Name:   "temp",
	// 	Data:   data,
	// 	Parent: parent,
	// 	Left:   parent.Left,
	// 	Right:  parent.Right,
	// }

}

func findParentNode(node *Node, data int) *Node {
	// if node.Left.Data > data {
	//
	// }
	//
	// if node.Right.Data > data {
	//
	// }

	return nil
}
