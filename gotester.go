package golang-api-testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akash1729/golang-api-testing/utils"
)

type TestCase struct {
	TestName    string
	TestDetail  string
	Route       string
	Method      string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
	StatusCode  int
	AvoidKey    []string
	RequestMap  map[string]interface{}
	ResponseMap map[string]interface{}
	TypeCheck   map[string]interface{} //assign key and sample type
}

func RunTest(testCase *TestCase, t *testing.T) error {

	fmt.Printf("%s Testing : %s, STATUS : TESTING", testCase.TestName, testCase.TestDetail)

	requestBody, _ := json.Marshal(testCase.RequestMap)

	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest(testCase.Method, testCase.Route, bytes.NewReader(requestBody))

	http.HandlerFunc(testCase.HandlerFunc).ServeHTTP(recorder, request)

	// Check Status Code
	resultStatusCode := recorder.Result().StatusCode
	utils.CompareInt(t, testCase.StatusCode, resultStatusCode)

	obtainedResult := recorder.Body.String()
	obtainedValueMap := make(map[string]interface{})
	json.Unmarshal([]byte(obtainedResult), &obtainedValueMap)

	//do type check for required fields
	utils.CompareTypeMap(t, testCase.TypeCheck, obtainedValueMap)

	// remove keys where value is dynamic, eg: token
	obtainedValueMap, err := utils.RemoveKey(obtainedValueMap, testCase.AvoidKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	utils.CompareMaps(t, testCase.ResponseMap, obtainedValueMap)

	fmt.Printf("%s Testing : %s, STATUS : FINISHED", testCase.TestName, testCase.TestDetail)
	return nil
}
