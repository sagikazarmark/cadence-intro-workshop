package main

import (
	"github.com/sagikazarmark/cadence-intro-workshop/examples/example01"
	"github.com/sagikazarmark/cadence-intro-workshop/examples/example02"
	"github.com/sagikazarmark/cadence-intro-workshop/examples/example03"
	"github.com/sagikazarmark/cadence-intro-workshop/examples/example04"
	"github.com/sagikazarmark/cadence-intro-workshop/examples/example05"
	"github.com/sagikazarmark/cadence-intro-workshop/examples/example06"
	"github.com/sagikazarmark/cadence-intro-workshop/examples/example07"
	"github.com/sagikazarmark/cadence-intro-workshop/examples/example08"
	"github.com/sagikazarmark/cadence-intro-workshop/examples/example09"
	"github.com/sagikazarmark/cadence-intro-workshop/examples/example10"
	"github.com/sagikazarmark/cadence-intro-workshop/examples/example11"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
)

func register(worker worker.Worker) {
	worker.RegisterWorkflow(example01.Workflow)
	worker.RegisterWorkflowWithOptions(example01.Workflow, workflow.RegisterOptions{Name: "example01"})

	worker.RegisterWorkflowWithOptions(example02.Workflow, workflow.RegisterOptions{Name: "example02"})

	worker.RegisterWorkflowWithOptions(example03.Workflow, workflow.RegisterOptions{Name: "example03"})

	worker.RegisterWorkflowWithOptions(example04.Workflow, workflow.RegisterOptions{Name: "example04"})

	worker.RegisterWorkflowWithOptions(example05.Workflow, workflow.RegisterOptions{Name: "example05"})

	worker.RegisterWorkflowWithOptions(example06.Workflow, workflow.RegisterOptions{Name: "example06"})

	worker.RegisterWorkflowWithOptions(example07.Workflow, workflow.RegisterOptions{Name: "example07"})

	worker.RegisterWorkflowWithOptions(example08.Workflow, workflow.RegisterOptions{Name: "example08"})

	worker.RegisterWorkflowWithOptions(example09.Workflow, workflow.RegisterOptions{Name: "example09"})
	worker.RegisterActivity(example09.Activity)
	worker.RegisterActivityWithOptions(example09.Activity, activity.RegisterOptions{Name: "example09"})

	worker.RegisterWorkflowWithOptions(example10.Workflow, workflow.RegisterOptions{Name: "example10"})
	worker.RegisterActivity(example10.Activity)
	worker.RegisterActivityWithOptions(example10.Activity, activity.RegisterOptions{Name: "example10"})

	worker.RegisterWorkflowWithOptions(example11.Workflow, workflow.RegisterOptions{Name: "example11"})
	worker.RegisterActivity(example11.Activity)
	worker.RegisterActivityWithOptions(example11.Activity, activity.RegisterOptions{Name: "example11"})
}
