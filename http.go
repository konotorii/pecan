package main

import (
	"fmt"
	"net/http"
)

type CustomHeaders struct {
	label string
	value string
}

var client = http.Client{}

func getRequest(url string, headers []CustomHeaders) *http.Response {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
	}

	for _, value := range headers {
		req.Header.Set(value.label, value.value)
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}

	fmt.Printf(res.Status)

	return res
}

func postRequest() {

}

func putRequest() {

}
