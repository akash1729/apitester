package usage

import (
	"testing"

	"github.com/akash1729/apitester"
)

func GetUserTests(t *testing.T) {

	testCase := apitester.TestCase{
		TestName:    "CreateUser",
		TestDetail:  "Proper Implemetaion",
		Route:       "/User",
		Method:      "POST",
		HandlerFunc: CreateUser, //import and assign corresponding handler
		StatusCode:  200,
		AvoidKey:    []string{"token"},
		RequestHeader: map[string]string{
			"Authorization": "authorization key",
		},
		RequestMap: map[string]interface{}{
			"username": "nitya",
			"password": "password"},
		ResponseHeader: map[string]string{
			"Content-Type": "application/json",
		},
		ResponseMap: map[string]interface{}{
			"status":   "User Created succesfully",
			"userID":   88,
			"username": "nitya"},
		TypeCheck: map[string]interface{}{
			"token":    "stringType",
			"userID":   0,
			"username": "stringType",
		},
		RequestContextKey:   "requestID",
		RequestContextValue: 123,
	}

	apitester.RunTest(&testCase, t)

}
