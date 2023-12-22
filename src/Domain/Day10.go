package Domain

type D10Tile struct {
	Tile    rune
	Visited bool
	Step    *D10Step
}

const (
	TL rune = 'F'
	TR      = '7'
	BR      = 'J'
	BL      = 'L'
	UD      = '|'
	LR      = '-'
	NO      = '.'
	ST      = 'S'
)

func (d D10Tile) GetNeighbors() []D10Offset {
	var offsets []D10Offset
	switch d.Tile {
	case TL:
		offsets = append(offsets, D10Offset{X: 1, Y: 0})
		offsets = append(offsets, D10Offset{X: 0, Y: 1})
	case TR:
		offsets = append(offsets, D10Offset{X: -1, Y: 0})
		offsets = append(offsets, D10Offset{X: 0, Y: 1})
	case BL:
		offsets = append(offsets, D10Offset{X: 1, Y: 0})
		offsets = append(offsets, D10Offset{X: 0, Y: -1})
	case BR:
		offsets = append(offsets, D10Offset{X: -1, Y: 0})
		offsets = append(offsets, D10Offset{X: 0, Y: -1})
	case UD:
		offsets = append(offsets, D10Offset{X: 0, Y: -1})
		offsets = append(offsets, D10Offset{X: 0, Y: 1})
	case LR:
		offsets = append(offsets, D10Offset{X: -1, Y: 0})
		offsets = append(offsets, D10Offset{X: 1, Y: 0})
	case ST:
		offsets = append(offsets, D10Offset{X: -1, Y: 0})
		offsets = append(offsets, D10Offset{X: 1, Y: 0})
		offsets = append(offsets, D10Offset{X: 0, Y: -1})
		offsets = append(offsets, D10Offset{X: 0, Y: 1})
	}

	return offsets
}

type D10Offset struct {
	X, Y int
}

type D10Step struct {
	X, Y  int
	From  *D10Step
	Steps int
}

func (s D10Step) Walk(tile D10Tile) []D10Step {
	var nextSteps []D10Step
	for _, neigh := range tile.GetNeighbors() {
		//	if s.From != nil {
		//fmt.Printf("Compare %d:%d <=> %d:%d\n", s.X+neigh.X, s.Y+neigh.Y, s.From.X, s.From.Y)
		//		}
		if s.From != nil && s.X+neigh.X == s.From.X && s.Y+neigh.Y == s.From.Y {
			continue
		}

		nextStep := D10Step{X: s.X + neigh.X, Y: s.Y + neigh.Y, Steps: s.Steps + 1, From: &s}
		nextSteps = append(nextSteps, nextStep)
	}

	return nextSteps
}

type D10Map struct {
	Map  [][]*D10Tile
	Next []D10Step
}

func (m *D10Map) FollowPath() *D10Step {
	var nextNext []D10Step
	for _, step := range m.Next {
		if step.X < 0 || step.Y < 0 || step.X >= len(m.Map[0]) || step.Y >= len(m.Map) {
			continue
		}

		tile := m.Map[step.Y][step.X]
		if tile.Visited == true {
			return &step
		}

		tile.Visited = true
		tile.Step = &step

		nextSteps := step.Walk(*tile)

		nextNext = append(nextNext, nextSteps...)
	}

	m.Next = nextNext

	return nil
}
