// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/ONSdigital/dp-authorisation/v2/authorisation"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	"net/http"
	"sync"
)

// Ensure, that MiddlewareMock does implement authorisation.Middleware.
// If this is not the case, regenerate this file with moq.
var _ authorisation.Middleware = &MiddlewareMock{}

// MiddlewareMock is a mock implementation of authorisation.Middleware.
//
// 	func TestSomethingThatUsesMiddleware(t *testing.T) {
//
// 		// make and configure a mocked authorisation.Middleware
// 		mockedMiddleware := &MiddlewareMock{
// 			CloseFunc: func(ctx context.Context) error {
// 				panic("mock out the Close method")
// 			},
// 			HealthCheckFunc: func(ctx context.Context, state *health.CheckState) error {
// 				panic("mock out the HealthCheck method")
// 			},
// 			RequireFunc: func(permission string, handlerFunc http.HandlerFunc) http.HandlerFunc {
// 				panic("mock out the Require method")
// 			},
// 		}
//
// 		// use mockedMiddleware in code that requires authorisation.Middleware
// 		// and then make assertions.
//
// 	}
type MiddlewareMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func(ctx context.Context) error

	// HealthCheckFunc mocks the HealthCheck method.
	HealthCheckFunc func(ctx context.Context, state *health.CheckState) error

	// RequireFunc mocks the Require method.
	RequireFunc func(permission string, handlerFunc http.HandlerFunc) http.HandlerFunc

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// HealthCheck holds details about calls to the HealthCheck method.
		HealthCheck []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// State is the state argument value.
			State *health.CheckState
		}
		// Require holds details about calls to the Require method.
		Require []struct {
			// Permission is the permission argument value.
			Permission string
			// HandlerFunc is the handlerFunc argument value.
			HandlerFunc http.HandlerFunc
		}
	}
	lockClose       sync.RWMutex
	lockHealthCheck sync.RWMutex
	lockRequire     sync.RWMutex
}

// Close calls CloseFunc.
func (mock *MiddlewareMock) Close(ctx context.Context) error {
	if mock.CloseFunc == nil {
		panic("MiddlewareMock.CloseFunc: method is nil but Middleware.Close was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	return mock.CloseFunc(ctx)
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedMiddleware.CloseCalls())
func (mock *MiddlewareMock) CloseCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// HealthCheck calls HealthCheckFunc.
func (mock *MiddlewareMock) HealthCheck(ctx context.Context, state *health.CheckState) error {
	if mock.HealthCheckFunc == nil {
		panic("MiddlewareMock.HealthCheckFunc: method is nil but Middleware.HealthCheck was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		State *health.CheckState
	}{
		Ctx:   ctx,
		State: state,
	}
	mock.lockHealthCheck.Lock()
	mock.calls.HealthCheck = append(mock.calls.HealthCheck, callInfo)
	mock.lockHealthCheck.Unlock()
	return mock.HealthCheckFunc(ctx, state)
}

// HealthCheckCalls gets all the calls that were made to HealthCheck.
// Check the length with:
//     len(mockedMiddleware.HealthCheckCalls())
func (mock *MiddlewareMock) HealthCheckCalls() []struct {
	Ctx   context.Context
	State *health.CheckState
} {
	var calls []struct {
		Ctx   context.Context
		State *health.CheckState
	}
	mock.lockHealthCheck.RLock()
	calls = mock.calls.HealthCheck
	mock.lockHealthCheck.RUnlock()
	return calls
}

// Require calls RequireFunc.
func (mock *MiddlewareMock) Require(permission string, handlerFunc http.HandlerFunc) http.HandlerFunc {
	if mock.RequireFunc == nil {
		panic("MiddlewareMock.RequireFunc: method is nil but Middleware.Require was just called")
	}
	callInfo := struct {
		Permission  string
		HandlerFunc http.HandlerFunc
	}{
		Permission:  permission,
		HandlerFunc: handlerFunc,
	}
	mock.lockRequire.Lock()
	mock.calls.Require = append(mock.calls.Require, callInfo)
	mock.lockRequire.Unlock()
	return mock.RequireFunc(permission, handlerFunc)
}

// RequireCalls gets all the calls that were made to Require.
// Check the length with:
//     len(mockedMiddleware.RequireCalls())
func (mock *MiddlewareMock) RequireCalls() []struct {
	Permission  string
	HandlerFunc http.HandlerFunc
} {
	var calls []struct {
		Permission  string
		HandlerFunc http.HandlerFunc
	}
	mock.lockRequire.RLock()
	calls = mock.calls.Require
	mock.lockRequire.RUnlock()
	return calls
}
