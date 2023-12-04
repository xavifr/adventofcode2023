package Application

import (
	"adventofcode2023/Domain"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type d2Set struct {
	Red, Green, Blue int
}

type d2Game struct {
	ID   int
	Sets []d2Set
}

type Day2 struct {
	part1Set d2Set
}

func (d *Day2) Part1(input *bufio.Scanner) error {
	possibleGamesAdditioner := 0
	for input.Scan() {
		game := d.getDataFromGame(input.Text())

		possible := true
		for _, set := range game.Sets {
			if set.Red > d.part1Set.Red || set.Blue > d.part1Set.Blue || set.Green > d.part1Set.Green {
				possible = false
			}
		}

		if possible {
			possibleGamesAdditioner += game.ID
		}
	}

	fmt.Printf("Possible games additioner is %d\n", possibleGamesAdditioner)
	return nil
}

func (d *Day2) Part2(input *bufio.Scanner) error {
	powerGamesAdditioner := 0

	for input.Scan() {
		game := d.getDataFromGame(input.Text())

		minSet := d2Set{}
		for _, set := range game.Sets {
			minSet.Red = max(minSet.Red, set.Red)
			minSet.Green = max(minSet.Green, set.Green)
			minSet.Blue = max(minSet.Blue, set.Blue)
		}

		powerGamesAdditioner += minSet.Red * minSet.Green * minSet.Blue

	}

	fmt.Printf("Power games additioner is %d\n", powerGamesAdditioner)
	return nil
}

func (d *Day2) getDataFromGame(line string) d2Game {
	game := d2Game{}
	gameIdRegex, _ := regexp.Compile("^Game (\\d+): (.*)")
	setsRegex, _ := regexp.Compile("^([^;]+)(?:; (.*))?")

	gameIdString := gameIdRegex.FindStringSubmatch(line)
	game.ID, _ = strconv.Atoi(gameIdString[1])
	game.Sets = []d2Set{}

	setsString := gameIdString[2]
	for {
		setString := setsRegex.FindStringSubmatch(setsString)

		game.Sets = append(game.Sets, d.getSetFromRoll(setString[1]))

		if len(setString[2]) == 0 {
			break
		}
		setsString = setString[2]

	}

	return game
}

func (d *Day2) getSetFromRoll(roll string) (set d2Set) {
	set = d2Set{}
	setRegex, _ := regexp.Compile("^(\\d+) ([^,]+)(?:, (.*))?$")

	for {
		dices := setRegex.FindStringSubmatch(roll)
		switch dices[2] {
		case "red":
			set.Red, _ = strconv.Atoi(dices[1])
		case "blue":
			set.Blue, _ = strconv.Atoi(dices[1])
		case "green":
			set.Green, _ = strconv.Atoi(dices[1])
		}

		if len(dices[3]) == 0 {
			break
		}

		roll = dices[3]

	}
	return
}

func NewDay2(Red, Green, Blue int) Domain.Day {
	return &Day2{part1Set: d2Set{
		Red:   Red,
		Green: Green,
		Blue:  Blue,
	}}
}
