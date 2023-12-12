package Application

import (
	"adventofcode2023/Domain"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day5 struct {
	SeedsP1        []int64
	SeedsP2        []Domain.D5SeedRange
	Seed2Soil      Domain.D5Map
	Soil2Fert      Domain.D5Map
	Fert2Water     Domain.D5Map
	Water2Light    Domain.D5Map
	Light2Temp     Domain.D5Map
	Temp2Humid     Domain.D5Map
	Humid2Location Domain.D5Map
}

func (d *Day5) Part1(input *bufio.Scanner) error {
	d.parseInput(input)

	minLocation := int64(999999999999999999)

	for _, seed := range d.SeedsP1 {
		location := d.Humid2Location.Transform(d.Temp2Humid.Transform(d.Light2Temp.Transform(d.Water2Light.Transform(d.Fert2Water.Transform(d.Soil2Fert.Transform(d.Seed2Soil.Transform(seed)))))))

		if location < minLocation {
			fmt.Printf(" - Seed %d produces a new min! %d\n", seed, location)
			minLocation = location
		}
	}

	fmt.Println()
	fmt.Printf("Min location is %d\n", minLocation)

	return nil
}

func (d *Day5) Part2(input *bufio.Scanner) error {
	d.parseInput(input)

	minLocation := int64(999999999999999999)

	fmt.Printf("#### TODO: Add goroutines for each seed range ####")
	fmt.Printf("#### ATM it will take up to 4 minutes to run the task ####")

	for _, seedRange := range d.SeedsP2 {
		fmt.Printf(" - SeedRange %d + %d\n", seedRange.Start, seedRange.Size)

		for seed := seedRange.Start; seed < seedRange.Start+seedRange.Size; seed++ {
			location := d.Humid2Location.Transform(d.Temp2Humid.Transform(d.Light2Temp.Transform(d.Water2Light.Transform(d.Fert2Water.Transform(d.Soil2Fert.Transform(d.Seed2Soil.Transform(seed)))))))

			if location < minLocation {
				fmt.Printf(" * new min!! %d\n", minLocation)
				minLocation = location
			}
		}
	}

	fmt.Println()
	fmt.Printf("Min location is %d\n", minLocation)

	return nil
}

func (d *Day5) parseInput(input *bufio.Scanner) {
	d.SeedsP1 = []int64{}
	d.SeedsP2 = []Domain.D5SeedRange{}
	d.Seed2Soil = Domain.D5Map{}
	d.Soil2Fert = Domain.D5Map{}
	d.Fert2Water = Domain.D5Map{}
	d.Water2Light = Domain.D5Map{}
	d.Light2Temp = Domain.D5Map{}
	d.Temp2Humid = Domain.D5Map{}
	d.Humid2Location = Domain.D5Map{}

	seedsRegex, _ := regexp.Compile("^seeds: (.*)$")
	numsRegex, _ := regexp.Compile("^(\\d+) ?(.*)$")
	mapRegex, _ := regexp.Compile("^(.*) map:$")
	ruleRegex, _ := regexp.Compile("^(\\d+) (\\d+) (\\d+)$")

	var atMap *Domain.D5Map
	for input.Scan() {
		if input.Text() == "" {
			atMap = nil
			continue
		}

		if seedsRegex.MatchString(input.Text()) {
			atMap = nil
			seedsString := seedsRegex.FindStringSubmatch(input.Text())

			// loops seeds
			numsString := strings.TrimSpace(seedsString[1])
			base := int64(-1)

			for {
				numString := numsRegex.FindStringSubmatch(numsString)
				if len(numString) == 0 {
					break
				}

				num, _ := strconv.ParseInt(numString[1], 10, 64)
				d.SeedsP1 = append(d.SeedsP1, num)

				if base == -1 {
					base = num
				} else {
					d.SeedsP2 = append(d.SeedsP2, Domain.D5SeedRange{Start: base, Size: num})
					base = -1
				}

				if len(numString[2]) == 0 {
					break
				}

				numsString = strings.TrimSpace(numString[2])
			}

			continue
		}

		if mapRegex.MatchString(input.Text()) {
			mapString := mapRegex.FindStringSubmatch(input.Text())
			switch mapString[1] {
			case "seed-to-soil":
				atMap = &d.Seed2Soil
			case "soil-to-fertilizer":
				atMap = &d.Soil2Fert
			case "fertilizer-to-water":
				atMap = &d.Fert2Water
			case "water-to-light":
				atMap = &d.Water2Light
			case "light-to-temperature":
				atMap = &d.Light2Temp
			case "temperature-to-humidity":
				atMap = &d.Temp2Humid
			case "humidity-to-location":
				atMap = &d.Humid2Location
			}

			continue
		}

		if ruleRegex.MatchString(input.Text()) {
			ruleString := ruleRegex.FindStringSubmatch(input.Text())

			if atMap != nil {
				destination, _ := strconv.ParseInt(ruleString[1], 10, 64)
				source, _ := strconv.ParseInt(ruleString[2], 10, 64)
				size, _ := strconv.ParseInt(ruleString[3], 10, 64)
				atMap.AddRule(destination, source, size)
			}
			continue
		}
	}
}

func NewDay5() *Day5 {
	return &Day5{}
}
