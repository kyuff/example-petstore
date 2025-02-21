// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package pets_crud_test

import (
	"context"
	"github.com/kyuff/es-commands"
	"github.com/kyuff/example-petstore/internal/features/pets_crud"
	"sync"
)

// Ensure, that DispatcherMock does implement pets_crud.Dispatcher.
// If this is not the case, regenerate this file with moq.
var _ pets_crud.Dispatcher = &DispatcherMock{}

// DispatcherMock is a mock implementation of pets_crud.Dispatcher.
//
//	func TestSomethingThatUsesDispatcher(t *testing.T) {
//
//		// make and configure a mocked pets_crud.Dispatcher
//		mockedDispatcher := &DispatcherMock{
//			DispatchFunc: func(ctx context.Context, entityID string, cmd commands.Command) error {
//				panic("mock out the Dispatch method")
//			},
//		}
//
//		// use mockedDispatcher in code that requires pets_crud.Dispatcher
//		// and then make assertions.
//
//	}
type DispatcherMock struct {
	// DispatchFunc mocks the Dispatch method.
	DispatchFunc func(ctx context.Context, entityID string, cmd commands.Command) error

	// calls tracks calls to the methods.
	calls struct {
		// Dispatch holds details about calls to the Dispatch method.
		Dispatch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// EntityID is the entityID argument value.
			EntityID string
			// Cmd is the cmd argument value.
			Cmd commands.Command
		}
	}
	lockDispatch sync.RWMutex
}

// Dispatch calls DispatchFunc.
func (mock *DispatcherMock) Dispatch(ctx context.Context, entityID string, cmd commands.Command) error {
	if mock.DispatchFunc == nil {
		panic("DispatcherMock.DispatchFunc: method is nil but Dispatcher.Dispatch was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		EntityID string
		Cmd      commands.Command
	}{
		Ctx:      ctx,
		EntityID: entityID,
		Cmd:      cmd,
	}
	mock.lockDispatch.Lock()
	mock.calls.Dispatch = append(mock.calls.Dispatch, callInfo)
	mock.lockDispatch.Unlock()
	return mock.DispatchFunc(ctx, entityID, cmd)
}

// DispatchCalls gets all the calls that were made to Dispatch.
// Check the length with:
//
//	len(mockedDispatcher.DispatchCalls())
func (mock *DispatcherMock) DispatchCalls() []struct {
	Ctx      context.Context
	EntityID string
	Cmd      commands.Command
} {
	var calls []struct {
		Ctx      context.Context
		EntityID string
		Cmd      commands.Command
	}
	mock.lockDispatch.RLock()
	calls = mock.calls.Dispatch
	mock.lockDispatch.RUnlock()
	return calls
}
