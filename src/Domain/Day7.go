package Domain

import (
	"errors"
	"slices"
	"sort"
	"strings"
)

type D7Card rune

func (c D7Card) Value() int {
	if c >= '2' && c <= '9' {
		return int(c - '0')
	} else if c == 'T' {
		return 10
	} else if c == 'J' {
		return 11
	} else if c == 'Q' {
		return 12
	} else if c == 'K' {
		return 13
	} else if c == 'A' {
		return 14
	} else if c == '*' {
		return 1
	}

	return 0
}

type D7Hand struct {
	Cards    []D7Card
	HandType JType
	hand     string
	Bid      int
}

func (h *D7Hand) HandValue() int {
	cards := h.hand

	var t []int
	jokers := 0

	for len(cards) > 0 {
		cardType := cards[0]

		if cardType == '*' {
			jokers = strings.Count(cards, string(cardType))
		} else {
			t = append(t, strings.Count(cards, string(cardType)))
		}

		cards = strings.ReplaceAll(cards, string(cardType), "")
	}

	sort.Ints(t)
	slices.Reverse(t)

	if jokers > 0 {
		if len(t) == 0 {
			t = append(t, jokers)
		} else {
			t[0] += jokers
		}
	}

	if len(t) == 1 {
		return 7 // five of a kind
	} else if len(t) == 2 && (t[0] == 4) {
		return 6 // four of a kind
	} else if len(t) == 2 && (t[0] == 3) {
		return 5 // full house
	} else if len(t) == 3 && (t[0] == 3) {
		return 4 // three of a kind
	} else if len(t) == 3 && (t[0] == 2) {
		return 3 // two pairs
	} else if len(t) == 4 {
		return 2 // one pair
	} else if len(t) == 5 {
		return 1 // high card
	}

	return 0
}

type JType int

const (
	Joker JType = iota
	Jay
)

func NewD7Hand(hand string, bid int, handType JType) (D7Hand, error) {
	d7Hand := D7Hand{}
	if len(hand) != 5 {
		return d7Hand, errors.New("invalid hand")
	}

	if handType == Joker {
		hand = strings.ReplaceAll(hand, "J", "*")
	}

	for _, card := range hand {
		if D7Card(card).Value() == 0 {
			return d7Hand, errors.New("invalid card in hand")
		}

		d7Hand.Cards = append(d7Hand.Cards, D7Card(card))
	}

	d7Hand.hand = hand
	d7Hand.Bid = bid

	return d7Hand, nil
}

type SortHandsByValue []D7Hand

func (h SortHandsByValue) Len() int {
	return len(h)
}

func (h SortHandsByValue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h SortHandsByValue) Less(i, j int) bool {
	if h[i].HandValue() == h[j].HandValue() {
		for k, _ := range h[i].Cards {
			if h[i].Cards[k].Value() != h[j].Cards[k].Value() {
				return h[i].Cards[k].Value() < h[j].Cards[k].Value()
			}
		}
	}

	return h[i].HandValue() < h[j].HandValue()
}
