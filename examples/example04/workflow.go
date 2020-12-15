package example04

import (
	"go.uber.org/cadence/workflow"
)

func Workflow(ctx workflow.Context) error {
	workflow.GetLogger(ctx).Info("starting example 04")

	return nil
}
