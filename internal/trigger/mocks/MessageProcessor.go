// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/yfpanne/app-functions-sdk-go/v2/pkg/interfaces"
	mock "github.com/stretchr/testify/mock"

	types "github.com/edgexfoundry/go-mod-messaging/v2/pkg/types"
)

// MessageProcessor is an autogenerated mock type for the MessageProcessor type
type MessageProcessor struct {
	mock.Mock
}

// MessageReceived provides a mock function with given fields: ctx, envelope, outputHandler
func (_m *MessageProcessor) MessageReceived(ctx interfaces.AppFunctionContext, envelope types.MessageEnvelope, outputHandler interfaces.PipelineResponseHandler) error {
	ret := _m.Called(ctx, envelope, outputHandler)

	var r0 error
	if rf, ok := ret.Get(0).(func(interfaces.AppFunctionContext, types.MessageEnvelope, interfaces.PipelineResponseHandler) error); ok {
		r0 = rf(ctx, envelope, outputHandler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Process provides a mock function with given fields: ctx, envelope
func (_m *MessageProcessor) Process(ctx interfaces.AppFunctionContext, envelope types.MessageEnvelope) error {
	ret := _m.Called(ctx, envelope)

	var r0 error
	if rf, ok := ret.Get(0).(func(interfaces.AppFunctionContext, types.MessageEnvelope) error); ok {
		r0 = rf(ctx, envelope)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReceivedInvalidMessage provides a mock function with given fields:
func (_m *MessageProcessor) ReceivedInvalidMessage() {
	_m.Called()
}

type NewMessageProcessorT interface {
	mock.TestingT
	Cleanup(func())
}

// NewMessageProcessor creates a new instance of MessageProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMessageProcessor(t NewMessageProcessorT) *MessageProcessor {
	mock := &MessageProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
