package Application

import (
	"adventofcode2023/Domain"
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day7 struct {
}

func (d *Day7) Part1(input *bufio.Scanner) error {
	hands := []Domain.D7Hand{}
	for input.Scan() {
		handSplit := strings.Split(input.Text(), " ")
		bid, _ := strconv.Atoi(handSplit[1])
		hand, e := Domain.NewD7Hand(handSplit[0], bid, Domain.Jay)
		if e != nil {
			continue
		}
		hands = append(hands, hand)
	}

	sort.Sort(Domain.SortHandsByValue(hands))

	sumBids := 0
	for rank, hand := range hands {
		sumBids += (rank + 1) * hand.Bid
	}

	fmt.Printf("HandBids is %d\n", sumBids)
	return nil
}

func (d *Day7) Part2(input *bufio.Scanner) error {
	hands := []Domain.D7Hand{}
	for input.Scan() {
		handSplit := strings.Split(input.Text(), " ")
		bid, _ := strconv.Atoi(handSplit[1])
		hand, e := Domain.NewD7Hand(handSplit[0], bid, Domain.Joker)
		if e != nil {
			continue
		}
		hands = append(hands, hand)
	}

	sort.Sort(Domain.SortHandsByValue(hands))

	sumBids := 0
	for rank, hand := range hands {
		sumBids += (rank + 1) * hand.Bid
	}

	fmt.Printf("HandBids is %d\n", sumBids)
	return nil
}
