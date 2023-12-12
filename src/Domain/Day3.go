package Domain

type D3Num struct {
	Value int
	Start D3Pos
	End   D3Pos
}

type D3Pos struct {
	X int
	Y int
}

func (n *D3Num) GetPositionsAround() []D3Pos {
	positions := make([]D3Pos, 0)

	for x := n.Start.X - 1; x <= n.End.X+1; x++ {
		for y := n.Start.Y - 1; y <= n.End.Y+1; y++ {
			positions = append(positions, D3Pos{X: x, Y: y})
		}
	}
	return positions
}

type D3Map struct {
	SymMap [][]rune
	Nums   []D3Num
}
