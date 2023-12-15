package Application

import (
	"adventofcode2023/Domain"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Day6 struct {
}

func (d *Day6) Part1(input *bufio.Scanner) error {
	races := d.parseInputP1(input)

	mulWinningTimes := 1
	for _, race := range races {
		winningTimes := race.GetWinningPrepareTimesQuadraticFunction()

		mulWinningTimes *= winningTimes
	}

	fmt.Printf("Mul is %d\n", mulWinningTimes)
	return nil
}

func (d *Day6) Part2(input *bufio.Scanner) error {
	race := d.parseInputP2(input)

	winningTimes := race.GetWinningPrepareTimesQuadraticFunction()

	fmt.Printf("Times is %d\n", winningTimes)
	return nil
}

func (d *Day6) parseInputP1(input *bufio.Scanner) []Domain.D6Race {
	races := []Domain.D6Race{}

	input.Scan()
	in := strings.TrimSpace(strings.Split(input.Text(), ":")[1])
	for strings.Count(in, "  ") > 0 {
		in = strings.ReplaceAll(in, "  ", " ")
	}
	times := strings.Split(in, " ")

	input.Scan()
	in = strings.TrimSpace(strings.Split(input.Text(), ":")[1])
	for strings.Count(in, "  ") > 0 {
		in = strings.ReplaceAll(in, "  ", " ")
	}
	distances := strings.Split(in, " ")

	for i, _ := range times {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		races = append(races, Domain.D6Race{MaxTime: time, MaxDistance: distance})
	}

	return races
}

func (d *Day6) parseInputP2(input *bufio.Scanner) Domain.D6Race {
	input.Scan()
	in := strings.TrimSpace(strings.Split(input.Text(), ":")[1])
	for strings.Count(in, " ") > 0 {
		in = strings.ReplaceAll(in, " ", "")
	}
	time, _ := strconv.Atoi(in)

	input.Scan()
	in = strings.TrimSpace(strings.Split(input.Text(), ":")[1])
	for strings.Count(in, " ") > 0 {
		in = strings.ReplaceAll(in, " ", "")
	}
	distance, _ := strconv.Atoi(in)

	return Domain.D6Race{MaxTime: time, MaxDistance: distance}
}
