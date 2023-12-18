package Domain

type D8Direction string

const (
	Left  D8Direction = "L"
	Right             = "R"
)

type D8Node struct {
	Name, Left, Right string
}

func (n *D8Node) Move(dir D8Direction) string {
	if dir == Left {
		return n.Left
	}

	return n.Right
}

type D8Map struct {
	Nodes  map[string]*D8Node
	RootP1 *D8Node
	RootP2 []*D8Node
}

func (m *D8Map) Navigate(startNode, finalNode, path string) int {
	steps := 0
	actualNode := m.Nodes[startNode]

	for true {
		if (len(finalNode) == 0 && actualNode.Name[len(actualNode.Name)-1:] == "Z") || (len(finalNode) != 0 && actualNode.Name == finalNode) {
			break
		}
		direction := D8Direction(path[steps%len(path)])
		actualNode = m.Nodes[actualNode.Move(direction)]
		steps++

	}

	return steps
}

func (m *D8Map) NavigateP2(path string) int {
	steps := 0
	actualNodes := make([]*D8Node, len(m.RootP2))
	copy(actualNodes, m.RootP2)

	for true {
		allAtEnd := true
		for _, node := range actualNodes {
			if node.Name[len(node.Name)-1:] != "Z" {
				allAtEnd = false
				break
			}
		}

		if allAtEnd {
			break
		}

		direction := D8Direction(path[steps%len(path)])
		for i, _ := range actualNodes {
			actualNodes[i] = m.Nodes[actualNodes[i].Move(direction)]
		}
		steps++
	}

	return steps
}
