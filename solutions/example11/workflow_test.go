package example11

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

	s.env.OnActivity(Activity, mock.Anything, ActivityInput{0}).Return(ActivityOutput{true}, nil)
	s.env.OnActivity(Activity, mock.Anything, ActivityInput{1}).Return(ActivityOutput{false}, nil)
	s.env.OnActivity(Activity, mock.Anything, ActivityInput{2}).Return(ActivityOutput{true}, nil)
	s.env.OnActivity(Activity, mock.Anything, ActivityInput{3}).Return(ActivityOutput{false}, nil)
	s.env.OnActivity(Activity, mock.Anything, ActivityInput{4}).Return(ActivityOutput{true}, nil)
	s.env.OnActivity(Activity, mock.Anything, ActivityInput{5}).Return(ActivityOutput{false}, nil)

	s.env.ExecuteWorkflow(Workflow, WorkflowInput{[]int{0, 1, 2, 3, 4, 5}})

	s.Require().True(s.env.IsWorkflowCompleted())
	s.Require().NoError(s.env.GetWorkflowError())

	var output WorkflowOutput
	s.Require().NoError(s.env.GetWorkflowResult(&output))

	expectedOutput := WorkflowOutput{
		Count:     6,
		CountOdd:  3,
		CountEven: 3,
		Sum:       15,
	}

	s.Equal(expectedOutput, output)
}

func (s *WorkflowTestSuite) Test_NoNumbers() {
	s.env.RegisterWorkflow(Workflow)
	s.env.ExecuteWorkflow(Workflow, WorkflowInput{})

	s.Require().True(s.env.IsWorkflowCompleted())

	s.EqualError(s.env.GetWorkflowError(), "no numbers")
}
