package priorityqueue

type PriorityQueue interface {
	Insert(*Node) error
	Poll() *Node
	Remove(*Node) error
	IsEmpty() bool
	Print()
	Size() int
	GetElements() []*Node
	Peak() *Node
}
