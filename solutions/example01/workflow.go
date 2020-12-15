package example01

import "go.uber.org/cadence/workflow"

func Workflow(ctx workflow.Context, a int, b int) (int, error) {
	return a + b, nil
}
