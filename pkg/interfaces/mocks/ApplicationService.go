// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	bootstrapinterfaces "github.com/edgexfoundry/go-mod-bootstrap/v2/bootstrap/interfaces"
	clientsinterfaces "github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"

	config "github.com/edgexfoundry/go-mod-bootstrap/v2/config"

	http "net/http"

	interfaces "github.com/yfpanne/app-functions-sdk-go/v2/pkg/interfaces"

	logger "github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"

	mock "github.com/stretchr/testify/mock"

	registry "github.com/edgexfoundry/go-mod-registry/v2/registry"

	time "time"
)

// ApplicationService is an autogenerated mock type for the ApplicationService type
type ApplicationService struct {
	mock.Mock
}

// AddBackgroundPublisher provides a mock function with given fields: capacity
func (_m *ApplicationService) AddBackgroundPublisher(capacity int) (interfaces.BackgroundPublisher, error) {
	ret := _m.Called(capacity)

	var r0 interfaces.BackgroundPublisher
	if rf, ok := ret.Get(0).(func(int) interfaces.BackgroundPublisher); ok {
		r0 = rf(capacity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.BackgroundPublisher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(capacity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddBackgroundPublisherWithTopic provides a mock function with given fields: capacity, topic
func (_m *ApplicationService) AddBackgroundPublisherWithTopic(capacity int, topic string) (interfaces.BackgroundPublisher, error) {
	ret := _m.Called(capacity, topic)

	var r0 interfaces.BackgroundPublisher
	if rf, ok := ret.Get(0).(func(int, string) interfaces.BackgroundPublisher); ok {
		r0 = rf(capacity, topic)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.BackgroundPublisher)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(capacity, topic)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddFunctionsPipelineForTopics provides a mock function with given fields: id, topic, transforms
func (_m *ApplicationService) AddFunctionsPipelineForTopics(id string, topic []string, transforms ...func(interfaces.AppFunctionContext, interface{}) (bool, interface{})) error {
	_va := make([]interface{}, len(transforms))
	for _i := range transforms {
		_va[_i] = transforms[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id, topic)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, ...func(interfaces.AppFunctionContext, interface{}) (bool, interface{})) error); ok {
		r0 = rf(id, topic, transforms...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddRoute provides a mock function with given fields: route, handler, methods
func (_m *ApplicationService) AddRoute(route string, handler func(http.ResponseWriter, *http.Request), methods ...string) error {
	_va := make([]interface{}, len(methods))
	for _i := range methods {
		_va[_i] = methods[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, route, handler)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, func(http.ResponseWriter, *http.Request), ...string) error); ok {
		r0 = rf(route, handler, methods...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ApplicationSettings provides a mock function with given fields:
func (_m *ApplicationService) ApplicationSettings() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// BuildContext provides a mock function with given fields: correlationId, contentType
func (_m *ApplicationService) BuildContext(correlationId string, contentType string) interfaces.AppFunctionContext {
	ret := _m.Called(correlationId, contentType)

	var r0 interfaces.AppFunctionContext
	if rf, ok := ret.Get(0).(func(string, string) interfaces.AppFunctionContext); ok {
		r0 = rf(correlationId, contentType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.AppFunctionContext)
		}
	}

	return r0
}

// CommandClient provides a mock function with given fields:
func (_m *ApplicationService) CommandClient() clientsinterfaces.CommandClient {
	ret := _m.Called()

	var r0 clientsinterfaces.CommandClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.CommandClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.CommandClient)
		}
	}

	return r0
}

// DeviceClient provides a mock function with given fields:
func (_m *ApplicationService) DeviceClient() clientsinterfaces.DeviceClient {
	ret := _m.Called()

	var r0 clientsinterfaces.DeviceClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.DeviceClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.DeviceClient)
		}
	}

	return r0
}

// DeviceProfileClient provides a mock function with given fields:
func (_m *ApplicationService) DeviceProfileClient() clientsinterfaces.DeviceProfileClient {
	ret := _m.Called()

	var r0 clientsinterfaces.DeviceProfileClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.DeviceProfileClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.DeviceProfileClient)
		}
	}

	return r0
}

// DeviceServiceClient provides a mock function with given fields:
func (_m *ApplicationService) DeviceServiceClient() clientsinterfaces.DeviceServiceClient {
	ret := _m.Called()

	var r0 clientsinterfaces.DeviceServiceClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.DeviceServiceClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.DeviceServiceClient)
		}
	}

	return r0
}

// EventClient provides a mock function with given fields:
func (_m *ApplicationService) EventClient() clientsinterfaces.EventClient {
	ret := _m.Called()

	var r0 clientsinterfaces.EventClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.EventClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.EventClient)
		}
	}

	return r0
}

// GetAppSetting provides a mock function with given fields: setting
func (_m *ApplicationService) GetAppSetting(setting string) (string, error) {
	ret := _m.Called(setting)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(setting)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(setting)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAppSettingStrings provides a mock function with given fields: setting
func (_m *ApplicationService) GetAppSettingStrings(setting string) ([]string, error) {
	ret := _m.Called(setting)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(setting)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(setting)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSecret provides a mock function with given fields: path, keys
func (_m *ApplicationService) GetSecret(path string, keys ...string) (map[string]string, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string, ...string) map[string]string); ok {
		r0 = rf(path, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(path, keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListenForCustomConfigChanges provides a mock function with given fields: configToWatch, sectionName, changedCallback
func (_m *ApplicationService) ListenForCustomConfigChanges(configToWatch interface{}, sectionName string, changedCallback func(interface{})) error {
	ret := _m.Called(configToWatch, sectionName, changedCallback)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, func(interface{})) error); ok {
		r0 = rf(configToWatch, sectionName, changedCallback)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadConfigurableFunctionPipelines provides a mock function with given fields:
func (_m *ApplicationService) LoadConfigurableFunctionPipelines() (map[string]interfaces.FunctionPipeline, error) {
	ret := _m.Called()

	var r0 map[string]interfaces.FunctionPipeline
	if rf, ok := ret.Get(0).(func() map[string]interfaces.FunctionPipeline); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interfaces.FunctionPipeline)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadConfigurablePipeline provides a mock function with given fields:
func (_m *ApplicationService) LoadConfigurablePipeline() ([]func(interfaces.AppFunctionContext, interface{}) (bool, interface{}), error) {
	ret := _m.Called()

	var r0 []func(interfaces.AppFunctionContext, interface{}) (bool, interface{})
	if rf, ok := ret.Get(0).(func() []func(interfaces.AppFunctionContext, interface{}) (bool, interface{})); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]func(interfaces.AppFunctionContext, interface{}) (bool, interface{}))
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadCustomConfig provides a mock function with given fields: _a0, sectionName
func (_m *ApplicationService) LoadCustomConfig(_a0 interfaces.UpdatableConfig, sectionName string) error {
	ret := _m.Called(_a0, sectionName)

	var r0 error
	if rf, ok := ret.Get(0).(func(interfaces.UpdatableConfig, string) error); ok {
		r0 = rf(_a0, sectionName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoggingClient provides a mock function with given fields:
func (_m *ApplicationService) LoggingClient() logger.LoggingClient {
	ret := _m.Called()

	var r0 logger.LoggingClient
	if rf, ok := ret.Get(0).(func() logger.LoggingClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.LoggingClient)
		}
	}

	return r0
}

// MakeItRun provides a mock function with given fields:
func (_m *ApplicationService) MakeItRun() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MakeItStop provides a mock function with given fields:
func (_m *ApplicationService) MakeItStop() {
	_m.Called()
}

// MetricsManager provides a mock function with given fields:
func (_m *ApplicationService) MetricsManager() bootstrapinterfaces.MetricsManager {
	ret := _m.Called()

	var r0 bootstrapinterfaces.MetricsManager
	if rf, ok := ret.Get(0).(func() bootstrapinterfaces.MetricsManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bootstrapinterfaces.MetricsManager)
		}
	}

	return r0
}

// NotificationClient provides a mock function with given fields:
func (_m *ApplicationService) NotificationClient() clientsinterfaces.NotificationClient {
	ret := _m.Called()

	var r0 clientsinterfaces.NotificationClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.NotificationClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.NotificationClient)
		}
	}

	return r0
}

// RegisterCustomStoreFactory provides a mock function with given fields: name, factory
func (_m *ApplicationService) RegisterCustomStoreFactory(name string, factory func(interfaces.DatabaseInfo, config.Credentials) (interfaces.StoreClient, error)) error {
	ret := _m.Called(name, factory)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, func(interfaces.DatabaseInfo, config.Credentials) (interfaces.StoreClient, error)) error); ok {
		r0 = rf(name, factory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterCustomTriggerFactory provides a mock function with given fields: name, factory
func (_m *ApplicationService) RegisterCustomTriggerFactory(name string, factory func(interfaces.TriggerConfig) (interfaces.Trigger, error)) error {
	ret := _m.Called(name, factory)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, func(interfaces.TriggerConfig) (interfaces.Trigger, error)) error); ok {
		r0 = rf(name, factory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegistryClient provides a mock function with given fields:
func (_m *ApplicationService) RegistryClient() registry.Client {
	ret := _m.Called()

	var r0 registry.Client
	if rf, ok := ret.Get(0).(func() registry.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(registry.Client)
		}
	}

	return r0
}

// RequestTimeout provides a mock function with given fields:
func (_m *ApplicationService) RequestTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// SecretProvider provides a mock function with given fields:
func (_m *ApplicationService) SecretProvider() bootstrapinterfaces.SecretProvider {
	ret := _m.Called()

	var r0 bootstrapinterfaces.SecretProvider
	if rf, ok := ret.Get(0).(func() bootstrapinterfaces.SecretProvider); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bootstrapinterfaces.SecretProvider)
		}
	}

	return r0
}

// SetDefaultFunctionsPipeline provides a mock function with given fields: transforms
func (_m *ApplicationService) SetDefaultFunctionsPipeline(transforms ...func(interfaces.AppFunctionContext, interface{}) (bool, interface{})) error {
	_va := make([]interface{}, len(transforms))
	for _i := range transforms {
		_va[_i] = transforms[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...func(interfaces.AppFunctionContext, interface{}) (bool, interface{})) error); ok {
		r0 = rf(transforms...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetFunctionsPipeline provides a mock function with given fields: transforms
func (_m *ApplicationService) SetFunctionsPipeline(transforms ...func(interfaces.AppFunctionContext, interface{}) (bool, interface{})) error {
	_va := make([]interface{}, len(transforms))
	for _i := range transforms {
		_va[_i] = transforms[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...func(interfaces.AppFunctionContext, interface{}) (bool, interface{})) error); ok {
		r0 = rf(transforms...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreSecret provides a mock function with given fields: path, secretData
func (_m *ApplicationService) StoreSecret(path string, secretData map[string]string) error {
	ret := _m.Called(path, secretData)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]string) error); ok {
		r0 = rf(path, secretData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionClient provides a mock function with given fields:
func (_m *ApplicationService) SubscriptionClient() clientsinterfaces.SubscriptionClient {
	ret := _m.Called()

	var r0 clientsinterfaces.SubscriptionClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.SubscriptionClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.SubscriptionClient)
		}
	}

	return r0
}

type mockConstructorTestingTNewApplicationService interface {
	mock.TestingT
	Cleanup(func())
}

// NewApplicationService creates a new instance of ApplicationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApplicationService(t mockConstructorTestingTNewApplicationService) *ApplicationService {
	mock := &ApplicationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
