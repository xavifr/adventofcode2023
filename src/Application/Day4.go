package Application

import (
	"adventofcode2023/Domain"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day4 struct {
}

func (d *Day4) Part1(input *bufio.Scanner) error {
	score := 0
	for input.Scan() {
		card := d.parseCard(input.Text())

		//fmt.Printf("%+v\n", card)
		score += card.GetScore()
	}

	fmt.Printf("Final score is %d\n", score)
	return nil
}

func (d *Day4) Part2(input *bufio.Scanner) error {
	stack := Domain.D4CardStack{}
	for input.Scan() {
		card := d.parseCard(input.Text())

		stack.Cards = append(stack.Cards, Domain.D4CardWithCopies{D4Card: card})
	}

	stack.Play()
	fmt.Printf("Final number of cards is %d\n", stack.GetNumberOfCopies())

	return nil
}

func (d *Day4) parseCard(line string) Domain.D4Card {
	card := Domain.D4Card{}

	cardRegex, _ := regexp.Compile("^Card\\s+(\\d+): (.*) \\| (.*)")
	cardMatches := cardRegex.FindStringSubmatch(line)

	card.ID, _ = strconv.Atoi(cardMatches[1])

	numRegex, _ := regexp.Compile("^(\\d+)(.*)$")

	// winning
	numsString := strings.TrimSpace(cardMatches[2])
	for {
		numString := numRegex.FindStringSubmatch(numsString)
		if len(numString) == 0 {
			break
		}

		num, _ := strconv.Atoi(numString[1])
		card.WinningNumbers = append(card.WinningNumbers, num)

		if len(numString[2]) == 0 {
			break
		}

		numsString = strings.TrimSpace(numString[2])
	}

	// played
	numsString = strings.TrimSpace(cardMatches[3])
	for {
		numString := numRegex.FindStringSubmatch(numsString)
		if len(numString) == 0 {
			break
		}

		num, _ := strconv.Atoi(numString[1])
		card.PlayedNumbers = append(card.PlayedNumbers, num)

		if len(numString[2]) == 0 {
			break
		}

		numsString = strings.TrimSpace(numString[2])
	}

	return card
}
