package apitester

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akash1729/apitester/utils"
)

// TestCase Create a test case by creating instance of this struct.
// This structure is currently based on only json requests and responses
type TestCase struct {
	TestName            string                                       // name of the test, eg: GET Person
	TestDetail          string                                       // case which is being tested, eg: Person with invalid DOB
	Route               string                                       // Route, eg: /Person
	Method              string                                       // HTTP method, eg: POST
	HandlerFunc         func(w http.ResponseWriter, r *http.Request) // handler function
	StatusCode          int                                          // expected return status code
	AvoidKey            []string                                     // Keys with dynamic values like token or timestamp, eg: ["token"]
	RequestMap          map[string]interface{}                       // Request data that can be marshaled into json
	ResponseMap         map[string]interface{}                       // Response map that is unmarshaled from a json
	TypeCheck           map[string]interface{}                       // Values for type check, only the types of values are compared. For testing values like token
	RequestContextKey   interface{}                                  // Key value to assign if request has context with values
	RequestContextValue interface{}                                  // Value for context for the key

}

// RunTest To run the test call RunTest with the test case and testing package pointer
func RunTest(testCase *TestCase, t *testing.T) error {

	fmt.Printf("%s Testing : %s, STATUS : TESTING\n", testCase.TestName, testCase.TestDetail)

	requestBody, _ := json.Marshal(testCase.RequestMap)

	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest(testCase.Method, testCase.Route, bytes.NewReader(requestBody))

	if testCase.RequestContextKey != nil {
		// add context to the request
		ctx := request.Context()
		ctx = context.WithValue(ctx, testCase.RequestContextKey, testCase.RequestContextValue)
		request = request.WithContext(ctx)
	}

	http.HandlerFunc(testCase.HandlerFunc).ServeHTTP(recorder, request)

	// Check Status Code
	resultStatusCode := recorder.Result().StatusCode
	utils.CompareInt(t, testCase.StatusCode, resultStatusCode, "Status code does Not Match")

	obtainedValueMap := make(map[string]interface{})

	json.Unmarshal([]byte(recorder.Body.String()), &obtainedValueMap)

	//marshal and decode our expected map to json
	typeCheckMap := make(map[string]interface{})
	typeCheckJSON, _ := json.Marshal(testCase.TypeCheck)

	json.Unmarshal(typeCheckJSON, &typeCheckMap)

	//do type check for required fields
	utils.CompareTypeMap(t, typeCheckMap, obtainedValueMap, "Types does not match")

	// remove keys where value is dynamic, eg: token
	obtainedValueMap, err := utils.RemoveKey(obtainedValueMap, testCase.AvoidKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	utils.CompareMaps(t, testCase.ResponseMap, obtainedValueMap, "Response data does not match")

	fmt.Printf("%s Testing : %s, STATUS : FINISHED\n", testCase.TestName, testCase.TestDetail)
	return nil
}
