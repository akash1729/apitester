package test

import (
	"testing"

	"github.com/akash1729/apitester"
)

func TestStatusCorrect(t *testing.T) {

	apiTest := apitester.TestCase{
		TestName:    "Testing",
		TestDetail:  "CORRECT CASE",
		Route:       "/test_it",
		Method:      "POST",
		HandlerFunc: DummyHandler,
		StatusCode:  200,
		AvoidKey:    []string{"dummyValue"},
		RequestHeader: map[string]string{
			"dummy_header": "header_value",
		},
		RequestMap: map[string]interface{}{
			"test case": "CORRECT CASE",
		},
		ResponseHeader: map[string]string{
			"dummy_header": "response_header_value",
		},
		ResponseMap: map[string]interface{}{
			"result": "CORRECT",
		},
		TypeCheck: map[string]interface{}{
			"result":     "abc",
			"dummyValue": 65,
		},
		RequestContextKey:   "requestID",
		RequestContextValue: 123,
	}

	apitester.RunTest(&apiTest, t)

}

func TestStatusIncorrect(t *testing.T) {

	apiTest := apitester.TestCase{
		TestName:    "Testing",
		TestDetail:  "INCORRECT CASE",
		Route:       "/test_it",
		Method:      "POST",
		HandlerFunc: DummyHandler,
		StatusCode:  400,
		RequestHeader: map[string]string{
			"dummy_header": "header_value",
		},
		RequestMap: map[string]interface{}{
			"test case": "INCORRECT CASE",
		},
		ResponseHeader: map[string]string{
			"dummy_header": "response_header_value",
		},
		ResponseMap: map[string]interface{}{
			"status": "INCORRECT",
		},
		TypeCheck: map[string]interface{}{
			"status": "abc",
		},
	}

	apitester.RunTest(&apiTest, t)

}
func TestErrorCall(t *testing.T) {

	apiTest := apitester.TestCase{
		TestName:    "Testing",
		TestDetail:  "CORRECT CASE",
		Route:       "/test_it",
		Method:      "POST",
		HandlerFunc: DummyHandler,
		StatusCode:  200,
		AvoidKey:    []string{"dummyValue"},
		RequestHeader: map[string]string{
			"dummy_header": "header_value",
		},
		RequestMap: map[string]interface{}{
			"test case": "CORRECT CASE",
		},
		ResponseHeader: map[string]string{
			"dummy_header": "response_header_value",
		},
		ResponseMap: map[string]interface{}{
			"error": "CORRECT",
		},
		TypeCheck: map[string]interface{}{
			"result":     "abc",
			"dummyValue": 65,
		},
		RequestContextKey:   "requestID",
		RequestContextValue: 123,
	}

	apitester.RunTest(&apiTest, t)

}
