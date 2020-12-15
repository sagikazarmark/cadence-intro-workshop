package example11

import (
	"errors"
	"time"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

type WorkflowInput struct {
	Numbers []int
}

type WorkflowOutput struct {
	Count     int
	CountOdd  int
	CountEven int
	Sum       int
}

func Workflow(ctx workflow.Context, input WorkflowInput) (WorkflowOutput, error) {
	workflow.GetLogger(ctx).Info("starting example 08")

	output := WorkflowOutput{
		Count: len(input.Numbers),
	}

	if output.Count == 0 {
		return output, errors.New("no numbers")
	}

	var currentNumber int

	err := workflow.SetQueryHandler(ctx, "current_number", func() (int, error) {
		return currentNumber, nil
	})
	if err != nil {
		return output, err
	}

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		TaskList:               "workshop",
		ScheduleToCloseTimeout: 3*time.Second + 3*time.Second,
		ScheduleToStartTimeout: 3 * time.Second,
		StartToCloseTimeout:    3 * time.Second,
		HeartbeatTimeout:       0 * time.Second,
		WaitForCancellation:    false,
		ActivityID:             "",
		RetryPolicy: &workflow.RetryPolicy{
			InitialInterval:          time.Second,
			BackoffCoefficient:       1.0,
			MaximumInterval:          10 * time.Second,
			ExpirationInterval:       60 * time.Second,
			MaximumAttempts:          5,
			NonRetriableErrorReasons: []string{"REASONS"},
		},
	})

	for _, number := range input.Numbers {
		currentNumber = number

		workflow.Sleep(ctx, 10*time.Second)

		output.Sum += number

		var activityOutput ActivityOutput

		err := workflow.ExecuteActivity(ctx, Activity, ActivityInput{number}).Get(ctx, &activityOutput)
		if err != nil {
			return output, err
		}

		if activityOutput.Even {
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
