package calc

import (
	"errors"
)

// return sum of two intergers
func Add(numbers ...int) (error, int) {
	sum := 0

	if len(numbers) < 2 {
		return errors.New("Provide more than 2 numbers"), sum
	} else {
		for _, num := range numbers {
			sum += num
		}

		return sum
	}
}
