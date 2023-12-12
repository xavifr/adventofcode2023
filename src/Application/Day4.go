package Application

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Day4 struct {
}

type Card struct {
	ID             int
	WinningNumbers []int
	PlayedNumbers  []int
}

func (c *Card) getScore() int {
	guessedNumbers := c.getGuessedNumbers()

	if guessedNumbers == 0 {
		return 0
	}

	return int(math.Pow(2, float64(guessedNumbers-1)))
}

func (c *Card) getGuessedNumbers() int {
	found := 0
	for _, win := range c.WinningNumbers {
		for _, played := range c.PlayedNumbers {
			if win == played {
				found++
				break
			}
		}
	}

	return found
}

type CardWithCopies struct {
	Card
	Copies int
}

type CardStack struct {
	Cards []CardWithCopies
}

func (cs *CardStack) Reset() {
	for i, _ := range cs.Cards {
		cs.Cards[i].Copies = 1
	}
}

func (cs *CardStack) Play() int {
	cs.Reset()

	for index, card := range cs.Cards {
		if card.Copies == 0 {
			break
		}
		//fmt.Printf("At card %d\n", card.ID)

		cardsWon := card.getGuessedNumbers()

		//fmt.Printf("  - won %d cards\n", cardsWon)

		for i := index + 1; i <= index+cardsWon; i++ {
			if i >= len(cs.Cards) {
				break
			}

			cs.Cards[i].Copies += card.Copies
			//fmt.Printf("  - now card %d has %d copies\n", cs.Cards[i].ID, cs.Cards[i].Copies)
		}
	}

	return 0
}

func (cs *CardStack) GetNumberOfCopies() int {
	cards := 0
	for _, c := range cs.Cards {
		cards += c.Copies
	}

	return cards
}

func (d *Day4) Part1(input *bufio.Scanner) error {
	score := 0
	for input.Scan() {
		card := d.parseCard(input.Text())

		//fmt.Printf("%+v\n", card)
		score += card.getScore()
	}

	fmt.Printf("Final score is %d\n", score)
	return nil
}

func (d *Day4) Part2(input *bufio.Scanner) error {
	stack := CardStack{}
	for input.Scan() {
		card := d.parseCard(input.Text())

		stack.Cards = append(stack.Cards, CardWithCopies{Card: card})
	}

	stack.Play()
	fmt.Printf("Final number of cards is %d\n", stack.GetNumberOfCopies())

	return nil
}

func (d *Day4) parseCard(line string) Card {
	card := Card{}

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
