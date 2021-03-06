// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/fakovacic/sitemap/internal/sitemap"
	"net/http"
	"sync"
)

// Ensure, that ClientMock does implement sitemap.Client.
// If this is not the case, regenerate this file with moq.
var _ sitemap.Client = &ClientMock{}

// ClientMock is a mock implementation of sitemap.Client.
//
// 	func TestSomethingThatUsesClient(t *testing.T) {
//
// 		// make and configure a mocked sitemap.Client
// 		mockedClient := &ClientMock{
// 			DoFunc: func(req *http.Request) (*http.Response, error) {
// 				panic("mock out the Do method")
// 			},
// 		}
//
// 		// use mockedClient in code that requires sitemap.Client
// 		// and then make assertions.
//
// 	}
type ClientMock struct {
	// DoFunc mocks the Do method.
	DoFunc func(req *http.Request) (*http.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// Do holds details about calls to the Do method.
		Do []struct {
			// Req is the req argument value.
			Req *http.Request
		}
	}
	lockDo sync.RWMutex
}

// Do calls DoFunc.
func (mock *ClientMock) Do(req *http.Request) (*http.Response, error) {
	if mock.DoFunc == nil {
		panic("ClientMock.DoFunc: method is nil but Client.Do was just called")
	}
	callInfo := struct {
		Req *http.Request
	}{
		Req: req,
	}
	mock.lockDo.Lock()
	mock.calls.Do = append(mock.calls.Do, callInfo)
	mock.lockDo.Unlock()
	return mock.DoFunc(req)
}

// DoCalls gets all the calls that were made to Do.
// Check the length with:
//     len(mockedClient.DoCalls())
func (mock *ClientMock) DoCalls() []struct {
	Req *http.Request
} {
	var calls []struct {
		Req *http.Request
	}
	mock.lockDo.RLock()
	calls = mock.calls.Do
	mock.lockDo.RUnlock()
	return calls
}
