package calc

import (
	"errors"
	"math"
)

// return sum of two intergers
func Add(numbers ...int) (error, int) {
	sum := math.MaxInt32

	if len(numbers) < 2 {
		return errors.New("Provide more than 2 numbers"), sum
	} else {
		for _, num := range numbers {
			sum += num
		}

		return nil, sum
	}
}
