package test

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONError Writes error message to ResponseWriter
func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintf(w, `{"status":%q}`, err)
}

// DummyHandler Dummy Handler for testing golang-api-tester
func DummyHandler(w http.ResponseWriter, r *http.Request) {

	requestMap := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&requestMap)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")

	if requestMap["test case"] == "CORRECT CASE" {

		//test context
		ctx := r.Context()
		value := ctx.Value("requestID")
		valueInt, ok := value.(int)
		if !ok {
			panic("failed to find context value")
		}

		if valueInt != 123 {
			panic("context values does not match")
		}

		responseMap := map[string]interface{}{
			"result":     "CORRECT",
			"dummyValue": 1234,
		}
		jsonResponse, err := json.Marshal(responseMap)
		if err != nil {
			panic(err)
		}

		if _, err := w.Write(jsonResponse); err != nil {
			panic(err)
		}
	}

	if requestMap["test case"] == "INCORRECT CASE" {

		JSONError(w, "INCORRECT", http.StatusBadRequest)

	}
}
