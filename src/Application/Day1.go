package Application

import (
	"bufio"
	"fmt"
	"strconv"
)

type Day1 struct {
	debug bool
}

func (d *Day1) SetDebug(debug bool) {
	d.debug = debug
}

func (d *Day1) Part1(input *bufio.Scanner) error {
	calibration := 0
	for input.Scan() {
		firstDigit := -1
		lastDigit := -1
		text := input.Text()
		for i := 0; i < len(text); i++ {
			char := text[i]
			num, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}

			if firstDigit == -1 {
				firstDigit = num
			}

			lastDigit = num
		}

		calibrationLine, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
		if err != nil {
			return err
		}

		calibration += calibrationLine
	}

	fmt.Printf("Calibration number is %d\n", calibration)
	return nil
}

func (d *Day1) Part2(input *bufio.Scanner) error {
	calibration := 0
	for input.Scan() {
		firstDigit := -1
		lastDigit := -1
		text := input.Text()
		for i := 0; i < len(text); i++ {
			char := text[i]
			num, err := strconv.Atoi(string(char))
			if err != nil {
				num = d.searchDigit(text[i:])
				if num == -1 {
					continue
				}
			}

			if firstDigit == -1 {
				firstDigit = num
			}

			lastDigit = num
		}

		calibrationLine, err := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
		if err != nil {
			return err
		}

		calibration += calibrationLine
	}

	fmt.Printf("Calibration number is %d\n", calibration)
	return nil
}

func (d *Day1) searchDigit(name string) int {
	if len(name) >= 3 && name[:3] == "one" {
		return 1
	} else if len(name) >= 3 && name[:3] == "two" {
		return 2
	} else if len(name) >= 5 && name[:5] == "three" {
		return 3
	} else if len(name) >= 4 && name[:4] == "four" {
		return 4
	} else if len(name) >= 4 && name[:4] == "five" {
		return 5
	} else if len(name) >= 3 && name[:3] == "six" {
		return 6
	} else if len(name) >= 5 && name[:5] == "seven" {
		return 7
	} else if len(name) >= 5 && name[:5] == "eight" {
		return 8
	} else if len(name) >= 4 && name[:4] == "nine" {
		return 9
	}

	return -1
}
