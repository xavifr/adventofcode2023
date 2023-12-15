package Application

import (
	"adventofcode2023/Domain"
	"bufio"
	"fmt"
)

type Day3 struct {
}

func (d *Day3) parseMap(input *bufio.Scanner) Domain.D3Map {
	m := Domain.D3Map{}
	m.SymMap = make([][]rune, 0)

	for input.Scan() {
		actualNumber := Domain.D3Num{}
		line := make([]rune, len(input.Text()))
		for i, t := range input.Text() {
			if t >= 48 && t <= 58 {
				if actualNumber.Value == 0 {
					actualNumber.Start = Domain.D3Pos{X: i, Y: len(m.SymMap)}
				}
				actualNumber.Value = actualNumber.Value*10 + (int(t) - 48)
				continue
			}

			if actualNumber.Value > 0 {
				actualNumber.End = Domain.D3Pos{X: i - 1, Y: len(m.SymMap)}
				m.Nums = append(m.Nums, actualNumber)
				actualNumber = Domain.D3Num{}
			}

			if t != 46 {
				line[i] = t
			}
		}

		if actualNumber.Value > 0 {
			actualNumber.End = Domain.D3Pos{X: len(input.Text()), Y: len(m.SymMap)}
			m.Nums = append(m.Nums, actualNumber)
			actualNumber = Domain.D3Num{}
		}

		m.SymMap = append(m.SymMap, line)
	}

	return m
}
func (d *Day3) Part1(input *bufio.Scanner) error {
	m := d.parseMap(input)

	sumNumbers := 0
	for _, num := range m.Nums {
		positions := num.GetPositionsAround()
		for _, pos := range positions {
			if pos.Y < 0 || pos.X < 0 || pos.Y >= len(m.SymMap) || pos.X >= len(m.SymMap[0]) {
				continue
			}

			if m.SymMap[pos.Y][pos.X] != 0 {
				//fmt.Printf("Symbol %c from %03d:%03d => %d\n", m.symMap[pos.Y][pos.X], num.Start.Y, num.Start.X, num.Value)
				sumNumbers = sumNumbers + num.Value
				break
			}
		}
	}

	fmt.Printf("D3Num is %d\n", sumNumbers)
	return nil
}

func (d *Day3) Part2(input *bufio.Scanner) error {
	m := d.parseMap(input)

	sumGears := 0
	for y := 0; y < len(m.SymMap); y++ {
		for x := 0; x < len(m.SymMap[y]); x++ {
			if m.SymMap[y][x] != '*' {
				continue
			}

			//fmt.Printf("Found possible gear at %d:%d\n", y, x)
			adjacentNumbers := []int{}
			for _, num := range m.Nums {
				for _, pos := range num.GetPositionsAround() {
					if pos.X == x && pos.Y == y {
						//fmt.Printf("  - found adjacent number %d\n", num.Value)
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
