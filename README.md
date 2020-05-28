# apitester : API testing package

## Overview [![GoDoc](https://godoc.org/github.com/akash1729/apitester?status.svg)](https://godoc.org/github.com/akash1729/apitester)

Package to unittest test golang API with json request and responses. 

## Install

```
go get github.com/akash1729/apitester
```

## Example

```
package usage

import (
	"github.com/akash1729/apitester"
)

func GetUserTests() {

	testCase := apitester.TestCase{
		TestName:    "CreateUser",
		TestDetail:  "Proper Implemetaion",
		Route:       "/User",
		Method:      "POST",
		HandlerFunc: CreateUser, //import and assign corresponding handler func
		StatusCode:  200,
		AvoidKey:    []string{"token"},
		RequestMap: map[string]interface{}{
			"username": "nitya",
			"password": "password"},
		ResponseMap: map[string]interface{}{
			"status":   "User Created succesfully",
			"userID":   88,
			"username": "nitya"},
		TypeCheck: map[string]interface{}{
			"token":    "stringType",
			"userID":   0,
			"username": "stringType",
		},
	}

	apitester.RunTest(testCase)

}

```

## Author

Akash Thomas

## License

MIT.

