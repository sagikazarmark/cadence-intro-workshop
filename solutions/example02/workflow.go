package example02

import (
	"go.uber.org/cadence/workflow"
)

type Input struct {
	A int
	B int
}

type Output struct {
	Result int
}

func Workflow(ctx workflow.Context, input Input) (Output, error) {
	workflow.GetLogger(ctx).Info("starting example 02")

	return Output{input.A + input.B}, nil
}
