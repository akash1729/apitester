package utils

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

// RemoveKey Remove Values from the map based on the list of keys
func RemoveKey(input map[string]interface{}, keys []string) (map[string]interface{}, error) {

	for _, key := range keys {

		if _, ok := input[key]; ok {
			delete(input, key)
		} else {
			return input, errors.New("Obtained response does not have the key : " + key)
		}
	}
	return input, nil
}

// CompareMaps Compare content of two maps by converting it into json since api handles json data
func CompareMaps(t *testing.T, expected map[string]interface{}, obtained map[string]interface{}, message string) {

	expectedResultJSON, _ := json.Marshal(expected)
	obtainedResultJSON, _ := json.Marshal(obtained)
	require.JSONEq(t, string(expectedResultJSON), string(obtainedResultJSON), message)

}

// CheckEqual Compare two values of any type
func CheckEqual(t *testing.T, expected interface{}, obtained interface{}, message string) {
	require.Equal(t, expected, obtained, message)
}

// CompareTypeMap Compares type of elements in a map recursively
func CompareTypeMap(t *testing.T, mapA map[string]interface{}, mapB map[string]interface{}, message string) {

	for key, value := range mapA {

		require.IsType(t, value, mapB[key], message)

		switch value.(type) {

		case map[string]interface{}:
			CompareTypeMap(t, value.(map[string]interface{}), mapB[key].(map[string]interface{}), message)

		case []interface{}:
			CompareTypeArray(t, value.([]interface{}), mapB[key].([]interface{}), message)

		}
	}

}

// CompareTypeArray Compare type of elements in an array recursively
func CompareTypeArray(t *testing.T, arrA []interface{}, arrB []interface{}, message string) {

	for index, value := range arrA {

		require.IsType(t, value, arrB[index])
		switch value.(type) {

		case map[string]interface{}:
			CompareTypeMap(t, value.(map[string]interface{}), arrB[index].(map[string]interface{}), message)

		case []interface{}:
			CompareTypeArray(t, value.([]interface{}), arrB[index].([]interface{}), message)
		}
	}

}
