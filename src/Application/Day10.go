package Application

import (
	"adventofcode2023/Domain"
	"bufio"
	"fmt"
	"math"
)

type Day10 struct {
}

func (d *Day10) Part1(input *bufio.Scanner) error {
	mmap := d.parseInput(input)

	var lastStep *Domain.D10Step
	for true {
		lastStep = mmap.FollowPath()
		if lastStep != nil {
			break
		}
	}

	fmt.Printf("Last step was at %d:%d with %d steps\n", lastStep.X, lastStep.Y, int(math.Ceil(float64(lastStep.Steps)/2)))
	return nil
}

func (d *Day10) Part2(input *bufio.Scanner) error {
	mmap := d.parseInput(input)

	for mmap.FollowPath() == nil {
		// no-do
	}

	inputTiles := 0
	for y := 0; y < len(mmap.Map); y++ {
		isInput := false
		wall := Domain.NO
		for x := 0; x < len(mmap.Map[y]); x++ {
			tile := mmap.Map[y][x]
			if tile.Visited {
				switch tile.Tile {
				case Domain.LR:
				case Domain.UD:
					isInput = !isInput
					wall = Domain.NO

				case Domain.BL:
					wall = tile.Tile
				case Domain.TL:
					wall = tile.Tile
				case Domain.TR:
					if wall == Domain.BL {
						isInput = !isInput
						wall = Domain.NO
					} else {
						wall = tile.Tile
					}
				case Domain.BR:
					if wall == Domain.TL {
						isInput = !isInput
						wall = Domain.NO
					} else {
						wall = tile.Tile
					}
				}

				continue
			}

			wall = Domain.NO
			if isInput {
				inputTiles++
			}
		}
	}

	fmt.Printf("Input tiles number: %d\n", inputTiles)
	return nil
}

func (d *Day10) parseInput(input *bufio.Scanner) Domain.D10Map {
	mmap := Domain.D10Map{}
	start := Domain.D10Step{}
	y := 0
	for input.Scan() {
		line := input.Text()
		var yLine []*Domain.D10Tile
		for x := 0; x < len(line); x++ {
			yLine = append(yLine, &Domain.D10Tile{Tile: rune(line[x])})

			if line[x] == 'S' {
				start.X = x
				start.Y = y
				start.From = &start
			}
		}

		mmap.Map = append(mmap.Map, yLine)
		y++
	}

	//mmap.Next = append(mmap.Next, start)
	//mmap.Next = start.Walk(mmap.Map[start.Y][start.X])
	stTile := mmap.Map[start.Y][start.X]
	for _, neigh := range stTile.GetNeighbors() {
		if start.Y+neigh.Y < 0 || start.X+neigh.X < 0 || start.Y+neigh.Y > len(mmap.Map) || start.X+neigh.X > len(mmap.Map[0]) {
			continue
		}
		neighTile := mmap.Map[start.Y+neigh.Y][start.X+neigh.X]
		for _, nn := range neighTile.GetNeighbors() {
			if start.X+neigh.X+nn.X == start.X && start.Y+neigh.Y+nn.Y == start.Y {
				mmap.Next = append(mmap.Next, Domain.D10Step{X: start.X + neigh.X, Y: start.Y + neigh.Y, From: &start, Steps: 1})
				break
			}
		}
	}

	stTile.Visited = true
	stTile.Step = &start

	// not happy about this code...
	if mmap.Next[0].X == mmap.Next[1].X {
		stTile.Tile = Domain.UD
	} else if mmap.Next[0].Y == mmap.Next[1].Y {
		stTile.Tile = Domain.LR
	} else if mmap.Next[0].X == start.X {
		if mmap.Next[0].Y < start.Y {
			if mmap.Next[1].X > start.X {
				stTile.Tile = Domain.BL
			} else {
				stTile.Tile = Domain.BR
			}
		} else {
			if mmap.Next[1].X > start.X {
				stTile.Tile = Domain.TL
			} else {
				stTile.Tile = Domain.TR
			}
		}
	} else if mmap.Next[0].Y == start.Y {
		if mmap.Next[0].X < start.X {
			if mmap.Next[1].Y > start.Y {
				stTile.Tile = Domain.TR
			} else {
				stTile.Tile = Domain.BR
			}
		} else {
			if mmap.Next[1].Y > start.Y {
				stTile.Tile = Domain.TL
			} else {
				stTile.Tile = Domain.BL
			}
		}
	}

	mmap.Next = mmap.Next[:1]

	return mmap
}
