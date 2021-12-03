package priorityqueue

type Node struct {
	Value int
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}
