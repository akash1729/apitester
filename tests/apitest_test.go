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
		RequestMap: map[string]interface{}{
			"test case": "CORRECT CASE",
		},
		ResponseMap: map[string]interface{}{
			"result": "CORRECT",
		},
		TypeCheck: map[string]interface{}{
			"result":     "abc",
			"dummyValue": 65,
		},
	}

	apitester.RunTest(&apiTest, t)

}

func TestStatusIncorrect(t *testing.T) {

	apiTest := apitester.TestCase{
		TestName:    "Testing",
		TestDetail:  "CORRECT CASE",
		Route:       "/test_it",
		Method:      "POST",
		HandlerFunc: DummyHandler,
		StatusCode:  400,
		RequestMap: map[string]interface{}{
			"test case": "INCORRECT CASE",
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
		RequestMap: map[string]interface{}{
			"test case": "CORRECT CASE",
		},
		ResponseMap: map[string]interface{}{
			"error": "CORRECT",
		},
		TypeCheck: map[string]interface{}{
			"result":     "abc",
			"dummyValue": 65,
		},
	}

	apitester.RunTest(&apiTest, t)

}
