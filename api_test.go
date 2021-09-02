package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
)

var host = "http://localhost:10000"

func TestHelloWorld(t *testing.T) {

	resp, err := http.Get(host + "/api")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	expected := `{ "Message": "Hello, world!" }`
	actual := string(body)

	if actual != expected {
		t.Error("helloWorld failed. Expected: " + expected + ", got: " + actual)
	} else {
		t.Logf("Expected: " + expected + ", got: " + actual)
	}

}

func TestCitiesListEndpoint_NoEmptyList(t *testing.T) {

	resp, err := http.Get(host + "/api/cities")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type cityCode struct {
		Code string
	}

	var citiesList []cityCode
	uerr := json.Unmarshal(body, &citiesList)

	if uerr != nil {
		t.Error("getCitiesList failed. Wrong JSON response.")
	}

	if len(citiesList) <= 0 {
		t.Error("getCitiesList failed. Expected: list of cities, got: empty list")
	} else {
		t.Logf("Expected: list of cities, got: " + string(body))
	}

}

func TestCitiesListEndpoint_NoEmptyCodes(t *testing.T) {

	resp, err := http.Get(host + "/api/cities")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type cityCode struct {
		Code string
	}

	var citiesList []cityCode
	uerr := json.Unmarshal(body, &citiesList)

	if uerr != nil {
		t.Error("getCitiesList failed. Wrong JSON response.")
	}

	for i, v := range citiesList {
		if v.Code == "" {
			t.Error("getCitiesList failed. Expected: all cities to have a code, got: no code for index " + strconv.Itoa(i))
		} else {
			t.Logf("Expected: all cities to have a code, got: " + v.Code)
		}
	}
}

func TestSingleCityEndpoint_ValidArg(t *testing.T) {

	validCode := "GDL"

	resp, err := http.Get(host + "/api/cities/" + validCode)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type City struct {
		Code string
		Name string
		Time string `json:"DateTime"`
	}

	var city City
	uerr := json.Unmarshal(body, &city)

	if uerr != nil {
		t.Error("getTimeRest failed. Wrong JSON response.")
	}

	if city.Code == "" || city.Name == "" || city.Time == "" {
		t.Error("getTimeRest failed. Expected: no empty fields, got: " + string(body))
	} else {
		t.Logf("Expected: no empty fields, got: " + string(body))
	}

}

func TestSingleCityEndpoint_InvalidArg(t *testing.T) {

	invalidCode := "QRO"

	resp, err := http.Get(host + "/api/cities/" + invalidCode)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type Err struct {
		Error string
	}

	var er Err
	uerr := json.Unmarshal(body, &er)

	if uerr != nil {
		t.Error("getTimeRest failed. Wrong JSON response.")
	}

	if er.Error != "invalid city code" {
		t.Error("getTimeRest failed. Expected: no empty fields, got: " + string(body))
	} else {
		t.Logf("Expected: no empty fields, got: " + string(body))
	}

}
