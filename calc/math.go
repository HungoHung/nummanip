package calc

import (
	"errors"

	"github.com/fatih/color"
)

func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		errMessage := color.RedString("provide more than 2 numbers")
		return sum, errors.New(errMessage)
	} else {
		for _, num := range numbers {
			sum += num
		}
		return sum, nil
	}

}
