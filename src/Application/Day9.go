package Application

import (
	"adventofcode2023/Domain"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Day9 struct {
}

func (d *Day9) Part1(input *bufio.Scanner) error {
	predictors := d.parseInput(input)

	sumPredictions := 0
	for _, pred := range predictors {
		nextValue := pred.NextValue()
		sumPredictions += nextValue
	}

	fmt.Printf("Sum of next predictions is %d\n", sumPredictions)

	return nil
}

func (d *Day9) Part2(input *bufio.Scanner) error {
	predictors := d.parseInput(input)

	sumPredictions := 0
	for _, pred := range predictors {
		prevValue := pred.PrevValue()
		sumPredictions += prevValue
	}

	fmt.Printf("Sum of prev predictions is %d\n", sumPredictions)

	return nil
}

func (d *Day9) parseInput(input *bufio.Scanner) []Domain.D9Predictor {
	var predictors []Domain.D9Predictor

	for input.Scan() {
		intArr := strings.Split(input.Text(), " ")

		pred := Domain.D9Predictor{}
		for _, intStr := range intArr {
			intVal, _ := strconv.Atoi(intStr)
			pred.Histogram = append(pred.Histogram, intVal)
		}

		predictors = append(predictors, pred)
	}

	return predictors
}
