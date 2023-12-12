package Domain

import "math"

type D4Card struct {
	ID             int
	WinningNumbers []int
	PlayedNumbers  []int
}

func (c *D4Card) GetScore() int {
	guessedNumbers := c.GetGuessedNumbers()

	if guessedNumbers == 0 {
		return 0
	}

	return int(math.Pow(2, float64(guessedNumbers-1)))
}

func (c *D4Card) GetGuessedNumbers() int {
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

type D4CardWithCopies struct {
	D4Card
	Copies int
}

type D4CardStack struct {
	Cards []D4CardWithCopies
}

func (cs *D4CardStack) Reset() {
	for i, _ := range cs.Cards {
		cs.Cards[i].Copies = 1
	}
}

func (cs *D4CardStack) Play() int {
	cs.Reset()

	for index, card := range cs.Cards {
		if card.Copies == 0 {
			break
		}
		//fmt.Printf("At card %d\n", card.ID)

		cardsWon := card.GetGuessedNumbers()

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

func (cs *D4CardStack) GetNumberOfCopies() int {
	cards := 0
	for _, c := range cs.Cards {
		cards += c.Copies
	}

	return cards
}
