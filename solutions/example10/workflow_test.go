package example10

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/cadence/testsuite"
)

func TestWorkflowTestSuite(t *testing.T) {
	suite.Run(t, new(WorkflowTestSuite))
}

type WorkflowTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite

	env *testsuite.TestWorkflowEnvironment
}

func (s *WorkflowTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *WorkflowTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func (s *WorkflowTestSuite) Test_Success() {
	s.env.RegisterWorkflow(Workflow)
	s.env.RegisterActivity(Activity)

	s.env.OnActivity(Activity, mock.Anything, ActivityInput{3, 4}).Return(ActivityOutput{7}, nil)

	s.env.ExecuteWorkflow(Workflow, WorkflowInput{3, 4})

	s.Require().True(s.env.IsWorkflowCompleted())
	s.Require().NoError(s.env.GetWorkflowError())

	var output WorkflowOutput
	s.Require().NoError(s.env.GetWorkflowResult(&output))

	s.Equal(7, output.Result)
}
