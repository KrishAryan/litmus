// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	chaos_experiment "github.com/litmuschaos/litmus/chaoscenter/graphql/server/pkg/database/mongodb/chaos_experiment"

	data_store "github.com/litmuschaos/litmus/chaoscenter/graphql/server/pkg/data-store"

	mock "github.com/stretchr/testify/mock"

	model "github.com/litmuschaos/litmus/chaoscenter/graphql/server/graph/model"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
)

// ChaosExperimentService is an autogenerated mock type for the Service type
type ChaosExperimentService struct {
	mock.Mock
}

type ChaosExperimentService_Expecter struct {
	mock *mock.Mock
}

func (_m *ChaosExperimentService) EXPECT() *ChaosExperimentService_Expecter {
	return &ChaosExperimentService_Expecter{mock: &_m.Mock}
}

// ProcessExperiment provides a mock function with given fields: ctx, workflow, projectID, revID
func (_m *ChaosExperimentService) ProcessExperiment(ctx context.Context, workflow *model.ChaosExperimentRequest, projectID string, revID string) (*model.ChaosExperimentRequest, *chaos_experiment.ChaosExperimentType, error) {
	ret := _m.Called(ctx, workflow, projectID, revID)

	if len(ret) == 0 {
		panic("no return value specified for ProcessExperiment")
	}

	var r0 *model.ChaosExperimentRequest
	var r1 *chaos_experiment.ChaosExperimentType
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.ChaosExperimentRequest, string, string) (*model.ChaosExperimentRequest, *chaos_experiment.ChaosExperimentType, error)); ok {
		return rf(ctx, workflow, projectID, revID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.ChaosExperimentRequest, string, string) *model.ChaosExperimentRequest); ok {
		r0 = rf(ctx, workflow, projectID, revID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ChaosExperimentRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.ChaosExperimentRequest, string, string) *chaos_experiment.ChaosExperimentType); ok {
		r1 = rf(ctx, workflow, projectID, revID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*chaos_experiment.ChaosExperimentType)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *model.ChaosExperimentRequest, string, string) error); ok {
		r2 = rf(ctx, workflow, projectID, revID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ChaosExperimentService_ProcessExperiment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessExperiment'
type ChaosExperimentService_ProcessExperiment_Call struct {
	*mock.Call
}

// ProcessExperiment is a helper method to define mock.On call
//   - ctx context.Context
//   - workflow *model.ChaosExperimentRequest
//   - projectID string
//   - revID string
func (_e *ChaosExperimentService_Expecter) ProcessExperiment(ctx interface{}, workflow interface{}, projectID interface{}, revID interface{}) *ChaosExperimentService_ProcessExperiment_Call {
	return &ChaosExperimentService_ProcessExperiment_Call{Call: _e.mock.On("ProcessExperiment", ctx, workflow, projectID, revID)}
}

func (_c *ChaosExperimentService_ProcessExperiment_Call) Run(run func(ctx context.Context, workflow *model.ChaosExperimentRequest, projectID string, revID string)) *ChaosExperimentService_ProcessExperiment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.ChaosExperimentRequest), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *ChaosExperimentService_ProcessExperiment_Call) Return(_a0 *model.ChaosExperimentRequest, _a1 *chaos_experiment.ChaosExperimentType, _a2 error) *ChaosExperimentService_ProcessExperiment_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ChaosExperimentService_ProcessExperiment_Call) RunAndReturn(run func(context.Context, *model.ChaosExperimentRequest, string, string) (*model.ChaosExperimentRequest, *chaos_experiment.ChaosExperimentType, error)) *ChaosExperimentService_ProcessExperiment_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessExperimentCreation provides a mock function with given fields: ctx, input, username, projectID, wfType, revisionID, r
func (_m *ChaosExperimentService) ProcessExperimentCreation(ctx context.Context, input *model.ChaosExperimentRequest, username string, projectID string, wfType *chaos_experiment.ChaosExperimentType, revisionID string, r *data_store.StateData) error {
	ret := _m.Called(ctx, input, username, projectID, wfType, revisionID, r)

	if len(ret) == 0 {
		panic("no return value specified for ProcessExperimentCreation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.ChaosExperimentRequest, string, string, *chaos_experiment.ChaosExperimentType, string, *data_store.StateData) error); ok {
		r0 = rf(ctx, input, username, projectID, wfType, revisionID, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChaosExperimentService_ProcessExperimentCreation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessExperimentCreation'
type ChaosExperimentService_ProcessExperimentCreation_Call struct {
	*mock.Call
}

// ProcessExperimentCreation is a helper method to define mock.On call
//   - ctx context.Context
//   - input *model.ChaosExperimentRequest
//   - username string
//   - projectID string
//   - wfType *chaos_experiment.ChaosExperimentType
//   - revisionID string
//   - r *data_store.StateData
func (_e *ChaosExperimentService_Expecter) ProcessExperimentCreation(ctx interface{}, input interface{}, username interface{}, projectID interface{}, wfType interface{}, revisionID interface{}, r interface{}) *ChaosExperimentService_ProcessExperimentCreation_Call {
	return &ChaosExperimentService_ProcessExperimentCreation_Call{Call: _e.mock.On("ProcessExperimentCreation", ctx, input, username, projectID, wfType, revisionID, r)}
}

func (_c *ChaosExperimentService_ProcessExperimentCreation_Call) Run(run func(ctx context.Context, input *model.ChaosExperimentRequest, username string, projectID string, wfType *chaos_experiment.ChaosExperimentType, revisionID string, r *data_store.StateData)) *ChaosExperimentService_ProcessExperimentCreation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.ChaosExperimentRequest), args[2].(string), args[3].(string), args[4].(*chaos_experiment.ChaosExperimentType), args[5].(string), args[6].(*data_store.StateData))
	})
	return _c
}

func (_c *ChaosExperimentService_ProcessExperimentCreation_Call) Return(_a0 error) *ChaosExperimentService_ProcessExperimentCreation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChaosExperimentService_ProcessExperimentCreation_Call) RunAndReturn(run func(context.Context, *model.ChaosExperimentRequest, string, string, *chaos_experiment.ChaosExperimentType, string, *data_store.StateData) error) *ChaosExperimentService_ProcessExperimentCreation_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessExperimentDelete provides a mock function with given fields: query, workflow, username, r
func (_m *ChaosExperimentService) ProcessExperimentDelete(query primitive.D, workflow chaos_experiment.ChaosExperimentRequest, username string, r *data_store.StateData) error {
	ret := _m.Called(query, workflow, username, r)

	if len(ret) == 0 {
		panic("no return value specified for ProcessExperimentDelete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.D, chaos_experiment.ChaosExperimentRequest, string, *data_store.StateData) error); ok {
		r0 = rf(query, workflow, username, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChaosExperimentService_ProcessExperimentDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessExperimentDelete'
type ChaosExperimentService_ProcessExperimentDelete_Call struct {
	*mock.Call
}

// ProcessExperimentDelete is a helper method to define mock.On call
//   - query primitive.D
//   - workflow chaos_experiment.ChaosExperimentRequest
//   - username string
//   - r *data_store.StateData
func (_e *ChaosExperimentService_Expecter) ProcessExperimentDelete(query interface{}, workflow interface{}, username interface{}, r interface{}) *ChaosExperimentService_ProcessExperimentDelete_Call {
	return &ChaosExperimentService_ProcessExperimentDelete_Call{Call: _e.mock.On("ProcessExperimentDelete", query, workflow, username, r)}
}

func (_c *ChaosExperimentService_ProcessExperimentDelete_Call) Run(run func(query primitive.D, workflow chaos_experiment.ChaosExperimentRequest, username string, r *data_store.StateData)) *ChaosExperimentService_ProcessExperimentDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(primitive.D), args[1].(chaos_experiment.ChaosExperimentRequest), args[2].(string), args[3].(*data_store.StateData))
	})
	return _c
}

func (_c *ChaosExperimentService_ProcessExperimentDelete_Call) Return(_a0 error) *ChaosExperimentService_ProcessExperimentDelete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChaosExperimentService_ProcessExperimentDelete_Call) RunAndReturn(run func(primitive.D, chaos_experiment.ChaosExperimentRequest, string, *data_store.StateData) error) *ChaosExperimentService_ProcessExperimentDelete_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessExperimentUpdate provides a mock function with given fields: workflow, username, wfType, revisionID, updateRevision, projectID, r
func (_m *ChaosExperimentService) ProcessExperimentUpdate(workflow *model.ChaosExperimentRequest, username string, wfType *chaos_experiment.ChaosExperimentType, revisionID string, updateRevision bool, projectID string, r *data_store.StateData) error {
	ret := _m.Called(workflow, username, wfType, revisionID, updateRevision, projectID, r)

	if len(ret) == 0 {
		panic("no return value specified for ProcessExperimentUpdate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.ChaosExperimentRequest, string, *chaos_experiment.ChaosExperimentType, string, bool, string, *data_store.StateData) error); ok {
		r0 = rf(workflow, username, wfType, revisionID, updateRevision, projectID, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChaosExperimentService_ProcessExperimentUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessExperimentUpdate'
type ChaosExperimentService_ProcessExperimentUpdate_Call struct {
	*mock.Call
}

// ProcessExperimentUpdate is a helper method to define mock.On call
//   - workflow *model.ChaosExperimentRequest
//   - username string
//   - wfType *chaos_experiment.ChaosExperimentType
//   - revisionID string
//   - updateRevision bool
//   - projectID string
//   - r *data_store.StateData
func (_e *ChaosExperimentService_Expecter) ProcessExperimentUpdate(workflow interface{}, username interface{}, wfType interface{}, revisionID interface{}, updateRevision interface{}, projectID interface{}, r interface{}) *ChaosExperimentService_ProcessExperimentUpdate_Call {
	return &ChaosExperimentService_ProcessExperimentUpdate_Call{Call: _e.mock.On("ProcessExperimentUpdate", workflow, username, wfType, revisionID, updateRevision, projectID, r)}
}

func (_c *ChaosExperimentService_ProcessExperimentUpdate_Call) Run(run func(workflow *model.ChaosExperimentRequest, username string, wfType *chaos_experiment.ChaosExperimentType, revisionID string, updateRevision bool, projectID string, r *data_store.StateData)) *ChaosExperimentService_ProcessExperimentUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.ChaosExperimentRequest), args[1].(string), args[2].(*chaos_experiment.ChaosExperimentType), args[3].(string), args[4].(bool), args[5].(string), args[6].(*data_store.StateData))
	})
	return _c
}

func (_c *ChaosExperimentService_ProcessExperimentUpdate_Call) Return(_a0 error) *ChaosExperimentService_ProcessExperimentUpdate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ChaosExperimentService_ProcessExperimentUpdate_Call) RunAndReturn(run func(*model.ChaosExperimentRequest, string, *chaos_experiment.ChaosExperimentType, string, bool, string, *data_store.StateData) error) *ChaosExperimentService_ProcessExperimentUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRuntimeCronWorkflowConfiguration provides a mock function with given fields: cronWorkflowManifest, experiment
func (_m *ChaosExperimentService) UpdateRuntimeCronWorkflowConfiguration(cronWorkflowManifest v1alpha1.CronWorkflow, experiment chaos_experiment.ChaosExperimentRequest) (v1alpha1.CronWorkflow, []string, error) {
	ret := _m.Called(cronWorkflowManifest, experiment)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRuntimeCronWorkflowConfiguration")
	}

	var r0 v1alpha1.CronWorkflow
	var r1 []string
	var r2 error
	if rf, ok := ret.Get(0).(func(v1alpha1.CronWorkflow, chaos_experiment.ChaosExperimentRequest) (v1alpha1.CronWorkflow, []string, error)); ok {
		return rf(cronWorkflowManifest, experiment)
	}
	if rf, ok := ret.Get(0).(func(v1alpha1.CronWorkflow, chaos_experiment.ChaosExperimentRequest) v1alpha1.CronWorkflow); ok {
		r0 = rf(cronWorkflowManifest, experiment)
	} else {
		r0 = ret.Get(0).(v1alpha1.CronWorkflow)
	}

	if rf, ok := ret.Get(1).(func(v1alpha1.CronWorkflow, chaos_experiment.ChaosExperimentRequest) []string); ok {
		r1 = rf(cronWorkflowManifest, experiment)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	if rf, ok := ret.Get(2).(func(v1alpha1.CronWorkflow, chaos_experiment.ChaosExperimentRequest) error); ok {
		r2 = rf(cronWorkflowManifest, experiment)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ChaosExperimentService_UpdateRuntimeCronWorkflowConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRuntimeCronWorkflowConfiguration'
type ChaosExperimentService_UpdateRuntimeCronWorkflowConfiguration_Call struct {
	*mock.Call
}

// UpdateRuntimeCronWorkflowConfiguration is a helper method to define mock.On call
//   - cronWorkflowManifest v1alpha1.CronWorkflow
//   - experiment chaos_experiment.ChaosExperimentRequest
func (_e *ChaosExperimentService_Expecter) UpdateRuntimeCronWorkflowConfiguration(cronWorkflowManifest interface{}, experiment interface{}) *ChaosExperimentService_UpdateRuntimeCronWorkflowConfiguration_Call {
	return &ChaosExperimentService_UpdateRuntimeCronWorkflowConfiguration_Call{Call: _e.mock.On("UpdateRuntimeCronWorkflowConfiguration", cronWorkflowManifest, experiment)}
}

func (_c *ChaosExperimentService_UpdateRuntimeCronWorkflowConfiguration_Call) Run(run func(cronWorkflowManifest v1alpha1.CronWorkflow, experiment chaos_experiment.ChaosExperimentRequest)) *ChaosExperimentService_UpdateRuntimeCronWorkflowConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(v1alpha1.CronWorkflow), args[1].(chaos_experiment.ChaosExperimentRequest))
	})
	return _c
}

func (_c *ChaosExperimentService_UpdateRuntimeCronWorkflowConfiguration_Call) Return(_a0 v1alpha1.CronWorkflow, _a1 []string, _a2 error) *ChaosExperimentService_UpdateRuntimeCronWorkflowConfiguration_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ChaosExperimentService_UpdateRuntimeCronWorkflowConfiguration_Call) RunAndReturn(run func(v1alpha1.CronWorkflow, chaos_experiment.ChaosExperimentRequest) (v1alpha1.CronWorkflow, []string, error)) *ChaosExperimentService_UpdateRuntimeCronWorkflowConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// NewChaosExperimentService creates a new instance of ChaosExperimentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChaosExperimentService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChaosExperimentService {
	mock := &ChaosExperimentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}