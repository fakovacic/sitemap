// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/fakovacic/sitemap/internal/sitemap"
	"sync"
)

// Ensure, that CollectorMock does implement sitemap.Collector.
// If this is not the case, regenerate this file with moq.
var _ sitemap.Collector = &CollectorMock{}

// CollectorMock is a mock implementation of sitemap.Collector.
//
// 	func TestSomethingThatUsesCollector(t *testing.T) {
//
// 		// make and configure a mocked sitemap.Collector
// 		mockedCollector := &CollectorMock{
// 			AddFunc: func(contextMoqParam context.Context, page sitemap.Page) bool {
// 				panic("mock out the Add method")
// 			},
// 			AllVisitedFunc: func(contextMoqParam context.Context) bool {
// 				panic("mock out the AllVisited method")
// 			},
// 			ListFunc: func(contextMoqParam context.Context) []sitemap.Page {
// 				panic("mock out the List method")
// 			},
// 			SetVisitFunc: func(contextMoqParam context.Context, s string)  {
// 				panic("mock out the SetVisit method")
// 			},
// 			TagFunc: func(contextMoqParam context.Context, s string, b bool)  {
// 				panic("mock out the Tag method")
// 			},
// 		}
//
// 		// use mockedCollector in code that requires sitemap.Collector
// 		// and then make assertions.
//
// 	}
type CollectorMock struct {
	// AddFunc mocks the Add method.
	AddFunc func(contextMoqParam context.Context, page sitemap.Page) bool

	// AllVisitedFunc mocks the AllVisited method.
	AllVisitedFunc func(contextMoqParam context.Context) bool

	// ListFunc mocks the List method.
	ListFunc func(contextMoqParam context.Context) []sitemap.Page

	// SetVisitFunc mocks the SetVisit method.
	SetVisitFunc func(contextMoqParam context.Context, s string)

	// TagFunc mocks the Tag method.
	TagFunc func(contextMoqParam context.Context, s string, b bool)

	// calls tracks calls to the methods.
	calls struct {
		// Add holds details about calls to the Add method.
		Add []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Page is the page argument value.
			Page sitemap.Page
		}
		// AllVisited holds details about calls to the AllVisited method.
		AllVisited []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
		// List holds details about calls to the List method.
		List []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
		// SetVisit holds details about calls to the SetVisit method.
		SetVisit []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// S is the s argument value.
			S string
		}
		// Tag holds details about calls to the Tag method.
		Tag []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// S is the s argument value.
			S string
			// B is the b argument value.
			B bool
		}
	}
	lockAdd        sync.RWMutex
	lockAllVisited sync.RWMutex
	lockList       sync.RWMutex
	lockSetVisit   sync.RWMutex
	lockTag        sync.RWMutex
}

// Add calls AddFunc.
func (mock *CollectorMock) Add(contextMoqParam context.Context, page sitemap.Page) bool {
	if mock.AddFunc == nil {
		panic("CollectorMock.AddFunc: method is nil but Collector.Add was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Page            sitemap.Page
	}{
		ContextMoqParam: contextMoqParam,
		Page:            page,
	}
	mock.lockAdd.Lock()
	mock.calls.Add = append(mock.calls.Add, callInfo)
	mock.lockAdd.Unlock()
	return mock.AddFunc(contextMoqParam, page)
}

// AddCalls gets all the calls that were made to Add.
// Check the length with:
//     len(mockedCollector.AddCalls())
func (mock *CollectorMock) AddCalls() []struct {
	ContextMoqParam context.Context
	Page            sitemap.Page
} {
	var calls []struct {
		ContextMoqParam context.Context
		Page            sitemap.Page
	}
	mock.lockAdd.RLock()
	calls = mock.calls.Add
	mock.lockAdd.RUnlock()
	return calls
}

// AllVisited calls AllVisitedFunc.
func (mock *CollectorMock) AllVisited(contextMoqParam context.Context) bool {
	if mock.AllVisitedFunc == nil {
		panic("CollectorMock.AllVisitedFunc: method is nil but Collector.AllVisited was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockAllVisited.Lock()
	mock.calls.AllVisited = append(mock.calls.AllVisited, callInfo)
	mock.lockAllVisited.Unlock()
	return mock.AllVisitedFunc(contextMoqParam)
}

// AllVisitedCalls gets all the calls that were made to AllVisited.
// Check the length with:
//     len(mockedCollector.AllVisitedCalls())
func (mock *CollectorMock) AllVisitedCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockAllVisited.RLock()
	calls = mock.calls.AllVisited
	mock.lockAllVisited.RUnlock()
	return calls
}

// List calls ListFunc.
func (mock *CollectorMock) List(contextMoqParam context.Context) []sitemap.Page {
	if mock.ListFunc == nil {
		panic("CollectorMock.ListFunc: method is nil but Collector.List was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc(contextMoqParam)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//     len(mockedCollector.ListCalls())
func (mock *CollectorMock) ListCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}

// SetVisit calls SetVisitFunc.
func (mock *CollectorMock) SetVisit(contextMoqParam context.Context, s string) {
	if mock.SetVisitFunc == nil {
		panic("CollectorMock.SetVisitFunc: method is nil but Collector.SetVisit was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		S               string
	}{
		ContextMoqParam: contextMoqParam,
		S:               s,
	}
	mock.lockSetVisit.Lock()
	mock.calls.SetVisit = append(mock.calls.SetVisit, callInfo)
	mock.lockSetVisit.Unlock()
	mock.SetVisitFunc(contextMoqParam, s)
}

// SetVisitCalls gets all the calls that were made to SetVisit.
// Check the length with:
//     len(mockedCollector.SetVisitCalls())
func (mock *CollectorMock) SetVisitCalls() []struct {
	ContextMoqParam context.Context
	S               string
} {
	var calls []struct {
		ContextMoqParam context.Context
		S               string
	}
	mock.lockSetVisit.RLock()
	calls = mock.calls.SetVisit
	mock.lockSetVisit.RUnlock()
	return calls
}

// Tag calls TagFunc.
func (mock *CollectorMock) Tag(contextMoqParam context.Context, s string, b bool) {
	if mock.TagFunc == nil {
		panic("CollectorMock.TagFunc: method is nil but Collector.Tag was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		S               string
		B               bool
	}{
		ContextMoqParam: contextMoqParam,
		S:               s,
		B:               b,
	}
	mock.lockTag.Lock()
	mock.calls.Tag = append(mock.calls.Tag, callInfo)
	mock.lockTag.Unlock()
	mock.TagFunc(contextMoqParam, s, b)
}

// TagCalls gets all the calls that were made to Tag.
// Check the length with:
//     len(mockedCollector.TagCalls())
func (mock *CollectorMock) TagCalls() []struct {
	ContextMoqParam context.Context
	S               string
	B               bool
} {
	var calls []struct {
		ContextMoqParam context.Context
		S               string
		B               bool
	}
	mock.lockTag.RLock()
	calls = mock.calls.Tag
	mock.lockTag.RUnlock()
	return calls
}
