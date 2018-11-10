package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetCategoriesHandler(t *testing.T) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", coreURL+"/api/v1/getCategories", nil)
	if err != nil {
		t.Error(cross, err.Error())
	}

	resp, err := client.Do(request)
	if err != nil {
		t.Error(cross, err.Error())
	}

	if resp.StatusCode != 200 {
		t.Error(cross, "Expected status code 200, but got: ", resp.StatusCode)
		r, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(cross, err.Error())
		}
		t.Error(cross, "Response body: ", string(r))
		return
	}

	t.Log(check, "Response status code: ", resp.StatusCode)
	if err != nil {
		t.Error(cross, err.Error())
	}

	r, err := ioutil.ReadAll(resp.Body)
	t.Log(check, "Response body: ", string(r))

}
