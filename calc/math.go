package calc

import (
	"errors"
)

func Add(numbers ...int) (error, int) {
	sum := 0
	if len(numbers) < 2 {
		return errors.New("Слишком мало чисел"), sum
	}

	for _, i := range numbers {
		sum += i
	}

	return nil, sum
}
