package priorityqueue

func findByIndex(elements []*Node, target *Node) int {
	if len(elements) == 0 {
		return -1
	}

	for idx, element := range elements {
		if element.Value == target.Value {
			return idx
		}
	}

	return -1
}
