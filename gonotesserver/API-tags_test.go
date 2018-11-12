package main

import (
	"bytes"
	"encoding/json"
	"gonotes/gonotesserver/pkg/model"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetTagsHandler(t *testing.T) {

	client := &http.Client{}

	request, err := http.NewRequest("GET", coreURL+"/api/v1/getTags", nil)
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

func TestAddTagHandler(t *testing.T) {

	client := &http.Client{}

	tag := model.Tag{
		Name:        "Data Strcuctures",
		Description: "Lists, trees and so on",
	}

	bodyJSON, _ := json.Marshal(tag)
	body := bytes.NewReader([]byte(bodyJSON))

	request, err := http.NewRequest("POST", coreURL+"/api/v1/addTag", body)
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
