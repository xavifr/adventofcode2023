package Domain

import (
	"slices"
)

type D9Histogram []int

func (h D9Histogram) AllZeros() bool {
	for _, v := range h {
		if v != 0 {
			return false
		}
	}

	return true
}

func (h D9Histogram) GetDifferences() D9Histogram {
	var nextHistogram []int

	prevV := 0
	for i, v := range h {
		if i > 0 {
			nextHistogram = append(nextHistogram, v-prevV)
		}

		prevV = v
	}

	return nextHistogram
}

type D9Predictor struct {
	Histogram   D9Histogram
	Differences []D9Histogram
}

func (p *D9Predictor) fillDifferences() {
	p.Differences = nil
	difference := p.Histogram
	for {
		difference = difference.GetDifferences()
		if difference == nil {
			break
		}
		p.Differences = append(p.Differences, difference)
		if difference.AllZeros() {
			break
		}
	}
}

func (p *D9Predictor) NextValue() int {
	p.fillDifferences()

	slices.Reverse(p.Differences)
	nextSum := 0
	for i, diff := range p.Differences {
		nextSum += diff[len(diff)-1]
		p.Differences[i] = append(diff, nextSum)
	}
	slices.Reverse(p.Differences)

	nextSum += p.Histogram[len(p.Histogram)-1]
	p.Histogram = append(p.Histogram, nextSum)
	return nextSum
}

func (p *D9Predictor) PrevValue() int {
	p.fillDifferences()

	slices.Reverse(p.Differences)
	prevSum := 0
	for i, diff := range p.Differences {
		prevSum = diff[0] - prevSum
		p.Differences[i] = append(D9Histogram{prevSum}, diff...)
	}
	slices.Reverse(p.Differences)

	prevSum = p.Histogram[0] - prevSum
	p.Histogram = append(D9Histogram{prevSum}, p.Histogram...)
	return prevSum
}
