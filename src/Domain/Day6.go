package Domain

import "math"

type D6Race struct {
	MaxTime     int
	MaxDistance int
}

func (r *D6Race) GetWinningPrepareTimesQuadraticFunction() int {
	insideSqrt := float64(r.MaxTime*r.MaxTime - 4*r.MaxDistance)
	if insideSqrt < 0 {
		return 0
	}

	firstRoot := math.Ceil(float64(r.MaxTime)-(math.Sqrt(insideSqrt))) / 2
	secondRoot := math.Floor(float64(r.MaxTime)+(math.Sqrt(insideSqrt))) / 2

	return int(math.Abs(secondRoot - firstRoot))
}

func (r *D6Race) GetWinningPrepareTimesBruteForce() []int {
	var winningPrepareTimes []int

	for i := 0; i <= r.MaxTime; i++ {
		if (r.MaxTime-i)*(i) > r.MaxDistance {
			winningPrepareTimes = append(winningPrepareTimes, i)
		}
	}

	return winningPrepareTimes
}
