package utils

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

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

func CompareInt(t *testing.T, expected int, obtained int) {

	require.Equal(t, expected, obtained, "Status code does Not Match")
}

func CompareMaps(t *testing.T, expected map[string]interface{}, obtained map[string]interface{}) {

	expectedResultJSON, _ := json.Marshal(expected)
	obtainedResultJSON, _ := json.Marshal(obtained)
	require.JSONEq(t, string(expectedResultJSON), string(obtainedResultJSON))

}

func CompareTypeMap(t *testing.T, mapA map[string]interface{}, mapB map[string]interface{}) {

	for key, value := range mapA {

		require.IsType(t, value, mapB[key])

		switch value.(type) {

		case map[string]interface{}:
			CompareTypeMap(t, value.(map[string]interface{}), mapB[key].(map[string]interface{}))

		case []interface{}:
			CompareTypeArray(t, value.([]interface{}), mapB[key].([]interface{}))

		}
	}

}

func CompareTypeArray(t *testing.T, arrA []interface{}, arrB []interface{}) {

	for index, value := range arrA {

		require.IsType(t, value, arrB[index])
		switch value.(type) {

		case map[string]interface{}:
			CompareTypeMap(t, value.(map[string]interface{}), arrB[index].(map[string]interface{}))

		case []interface{}:
			CompareTypeArray(t, value.([]interface{}), arrB[index].([]interface{}))
		}
	}

}
