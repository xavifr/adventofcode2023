package Application

import (
	"bufio"
	"fmt"
)

type Num struct {
	Value int
	Start Pos
	End   Pos
}

type Pos struct {
	X int
	Y int
}

func (n *Num) GetPositionsAround() []Pos {
	positions := make([]Pos, 0)

	for x := n.Start.X - 1; x <= n.End.X+1; x++ {
		for y := n.Start.Y - 1; y <= n.End.Y+1; y++ {
			positions = append(positions, Pos{X: x, Y: y})
		}
	}
	return positions
}

type Map struct {
	symMap [][]rune
	nums   []Num
}

type Day3 struct {
}

func (d *Day3) parseMap(input *bufio.Scanner) Map {
	m := Map{}
	m.symMap = make([][]rune, 0)

	for input.Scan() {
		actualNumber := Num{}
		line := make([]rune, len(input.Text()))
		for i, t := range input.Text() {
			if t >= 48 && t <= 58 {
				if actualNumber.Value == 0 {
					actualNumber.Start = Pos{X: i, Y: len(m.symMap)}
				}
				actualNumber.Value = actualNumber.Value*10 + (int(t) - 48)
				continue
			}

			if actualNumber.Value > 0 {
				actualNumber.End = Pos{X: i - 1, Y: len(m.symMap)}
				m.nums = append(m.nums, actualNumber)
				actualNumber = Num{}
			}

			if t != 46 {
				line[i] = t
			}
		}

		if actualNumber.Value > 0 {
			actualNumber.End = Pos{X: len(input.Text()), Y: len(m.symMap)}
			m.nums = append(m.nums, actualNumber)
			actualNumber = Num{}
		}

		m.symMap = append(m.symMap, line)
	}

	return m
}
func (d *Day3) Part1(input *bufio.Scanner) error {
	m := d.parseMap(input)

	sumNumbers := 0
	for _, num := range m.nums {
		positions := num.GetPositionsAround()
		for _, pos := range positions {
			if pos.Y < 0 || pos.X < 0 || pos.Y >= len(m.symMap) || pos.X >= len(m.symMap[0]) {
				continue
			}

			if m.symMap[pos.Y][pos.X] != 0 {
				//fmt.Printf("Symbol %c from %03d:%03d => %d\n", m.symMap[pos.Y][pos.X], num.Start.Y, num.Start.X, num.Value)
				sumNumbers = sumNumbers + num.Value
				break
			}
		}
	}

	fmt.Printf("Num is %d\n", sumNumbers)
	return nil
}

func (d *Day3) Part2(input *bufio.Scanner) error {
	m := d.parseMap(input)

	sumGears := 0
	for y := 0; y < len(m.symMap); y++ {
		for x := 0; x < len(m.symMap[y]); x++ {
			if m.symMap[y][x] != '*' {
				continue
			}

			fmt.Printf("Found possible gear at %d:%d\n", y, x)
			adjacentNumbers := []int{}
			for _, num := range m.nums {
				for _, pos := range num.GetPositionsAround() {
					if pos.X == x && pos.Y == y {
						fmt.Printf("  - found adjacent number %d\n", num.Value)
						adjacentNumbers = append(adjacentNumbers, num.Value)
						break
					}
				}
			}

			if len(adjacentNumbers) > 1 {
				gearValue := 1
				for _, v := range adjacentNumbers {
					gearValue *= v
				}

				sumGears += gearValue
			}
		}
	}

	fmt.Printf("Gears sum is %d\n", sumGears)
	return nil
}
