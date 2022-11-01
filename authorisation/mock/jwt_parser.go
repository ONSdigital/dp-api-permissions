// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/ONSdigital/dp-authorisation/v2/authorisation"
	permsdk "github.com/ONSdigital/dp-permissions-api/sdk"
	"sync"
)

// Ensure, that JWTParserMock does implement authorisation.JWTParser.
// If this is not the case, regenerate this file with moq.
var _ authorisation.JWTParser = &JWTParserMock{}

// JWTParserMock is a mock implementation of authorisation.JWTParser.
//
//	func TestSomethingThatUsesJWTParser(t *testing.T) {
//
//		// make and configure a mocked authorisation.JWTParser
//		mockedJWTParser := &JWTParserMock{
//			ParseFunc: func(tokenString string) (*permsdk.EntityData, error) {
//				panic("mock out the Parse method")
//			},
//		}
//
//		// use mockedJWTParser in code that requires authorisation.JWTParser
//		// and then make assertions.
//
//	}
type JWTParserMock struct {
	// ParseFunc mocks the Parse method.
	ParseFunc func(tokenString string) (*permsdk.EntityData, error)

	// calls tracks calls to the methods.
	calls struct {
		// Parse holds details about calls to the Parse method.
		Parse []struct {
			// TokenString is the tokenString argument value.
			TokenString string
		}
	}
	lockParse sync.RWMutex
}

// Parse calls ParseFunc.
func (mock *JWTParserMock) Parse(tokenString string) (*permsdk.EntityData, error) {
	if mock.ParseFunc == nil {
		panic("JWTParserMock.ParseFunc: method is nil but JWTParser.Parse was just called")
	}
	callInfo := struct {
		TokenString string
	}{
		TokenString: tokenString,
	}
	mock.lockParse.Lock()
	mock.calls.Parse = append(mock.calls.Parse, callInfo)
	mock.lockParse.Unlock()
	return mock.ParseFunc(tokenString)
}

// ParseCalls gets all the calls that were made to Parse.
// Check the length with:
//
//	len(mockedJWTParser.ParseCalls())
func (mock *JWTParserMock) ParseCalls() []struct {
	TokenString string
} {
	var calls []struct {
		TokenString string
	}
	mock.lockParse.RLock()
	calls = mock.calls.Parse
	mock.lockParse.RUnlock()
	return calls
}
