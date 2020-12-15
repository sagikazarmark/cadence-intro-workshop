package example03

import (
	"errors"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

type Input struct {
	Numbers []int
}

type Output struct {
	Count     int
	CountOdd  int
	CountEven int
	Sum       int
}

func Workflow(ctx workflow.Context, input Input) (Output, error) {
	workflow.GetLogger(ctx).Info("starting example 03")

	output := Output{
		Count: len(input.Numbers),
	}

	if output.Count == 0 {
		return output, errors.New("no numbers")
	}

	for _, number := range input.Numbers {
		output.Sum += number

		if number%2 == 0 {
			output.CountEven++
		} else {
			output.CountOdd++
		}

		if number%3 == 0 {
			workflow.GetLogger(ctx).Info("number divisible by 3", zap.Int("number", number))
		}
	}

	return output, nil
}
