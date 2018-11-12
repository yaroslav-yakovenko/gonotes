package main

import (
	"bytes"
	"encoding/json"
	"gonotes/gonotesserver/pkg/model"
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

}

func TestAddCategoryHandler(t *testing.T) {

	client := &http.Client{}

	cat := model.Category{
		Name:        "Code Snippets",
		Description: "Small chunks of basic Go constructions",
	}

	bodyJSON, _ := json.Marshal(cat)
	body := bytes.NewReader([]byte(bodyJSON))

	request, err := http.NewRequest("POST", coreURL+"/api/v1/addCategory", body)
	if err != nil {
		t.Error(cross, err.Error())
	}

	resp, err := client.Do(request)
	if err != nil {
		t.Error(cross, err.Error())
	}

	if resp.StatusCode == http.StatusCreated {
		t.Log(check, "Document created. Response status code: ", resp.StatusCode)
	}

	if resp.StatusCode == http.StatusConflict {
		t.Log(check, "Document exists. Response status code: ", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusConflict {
		t.Error(cross, "Expected status code 201 or 409, but got: ", resp.StatusCode)
		r, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(cross, err.Error())
		}
		t.Error(cross, "Response body: ", string(r))
		return
	}

	r, err := ioutil.ReadAll(resp.Body)
	t.Log(check, "Response body: ", string(r))

}
