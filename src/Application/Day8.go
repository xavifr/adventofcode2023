package Application

import (
	"adventofcode2023/Domain"
	"bufio"
	"fmt"
	"regexp"
)

type Day8 struct {
}

func (d *Day8) Part1(input *bufio.Scanner) error {
	path, nodesMap, startNode := d.parseInputP1(input)

	moves := nodesMap.Navigate(startNode.Name, "ZZZ", path)

	fmt.Printf("Moves to reach: %d\n", moves)

	return nil
}

func (d *Day8) Part2(input *bufio.Scanner) error {
	path, nodesMap, startNodes := d.parseInputP2(input)

	moves := []int{}
	for _, startNode := range startNodes {
		moves = append(moves, nodesMap.Navigate(startNode.Name, "", path))
	}

	finalMoves := 1
	for _, move := range moves {
		finalMoves = d.lcd(finalMoves, move)
	}

	fmt.Printf("Moves to reach: %d\n", finalMoves)

	return nil
}

func (d *Day8) parseInputP1(input *bufio.Scanner) (string, Domain.D8Map, *Domain.D8Node) {
	nodesMap := Domain.D8Map{Nodes: map[string]*Domain.D8Node{}}
	input.Scan()

	path := input.Text()

	nodeRegexp, _ := regexp.Compile("^(\\w+) = \\((\\w+), (\\w+)\\)$")

	for input.Scan() {
		if len(input.Text()) == 0 {
			continue
		}

		if !nodeRegexp.MatchString(input.Text()) {
			continue
		}

		match := nodeRegexp.FindStringSubmatch(input.Text())

		node := &Domain.D8Node{Name: match[1], Left: match[2], Right: match[3]}
		nodesMap.Nodes[match[1]] = node
	}

	startNode := nodesMap.Nodes["AAA"]

	return path, nodesMap, startNode
}

func (d *Day8) gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func (d *Day8) lcd(a, b int, integers ...int) int {
	result := a * b / d.gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = d.lcd(result, integers[i])
	}

	return result
}

func (d *Day8) parseInputP2(input *bufio.Scanner) (string, Domain.D8Map, []*Domain.D8Node) {
	startNodes := []*Domain.D8Node{}
	nodesMap := Domain.D8Map{Nodes: map[string]*Domain.D8Node{}}
	input.Scan()

	path := input.Text()

	nodeRegexp, _ := regexp.Compile("^(\\w+) = \\((\\w+), (\\w+)\\)$")
	startNodeRegexp, _ := regexp.Compile("^(\\w+)A$")

	for input.Scan() {
		if len(input.Text()) == 0 {
			continue
		}

		if !nodeRegexp.MatchString(input.Text()) {
			continue
		}

		match := nodeRegexp.FindStringSubmatch(input.Text())

		node := &Domain.D8Node{Name: match[1], Left: match[2], Right: match[3]}
		nodesMap.Nodes[match[1]] = node

		if startNodeRegexp.MatchString(node.Name) {
			startNodes = append(startNodes, node)
		}
	}

	return path, nodesMap, startNodes
}
