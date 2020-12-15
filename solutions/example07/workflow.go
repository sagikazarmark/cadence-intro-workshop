package example07

import (
	"time"

	"go.uber.org/cadence/workflow"
)

func Workflow(ctx workflow.Context) error {
	workflow.GetLogger(ctx).Info("starting example 07")

	var number int

	err := workflow.SetQueryHandler(ctx, "current_number", func() (int, error) {
		return number, nil
	})
	if err != nil {
		return err
	}

	for _, v := range []int{0, 1, 2, 3, 4, 5} {
		number = v

		workflow.Sleep(ctx, 10*time.Second)
	}

	return nil
}
