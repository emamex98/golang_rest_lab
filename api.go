package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var cities = createCities()

type city struct {
	Name string
	Code string
	Tmzn string
}

type cityCode struct {
	Code string
}

type timeStruct struct {
	CityName        string
	CurrentDateTime string `json:"currentDateTime"`
}

func createCities() (cities [6]city) {

	cities[0] = city{
		Name: "Guadalajara",
		Code: "GDL",
		Tmzn: "CST",
	}

	cities[1] = city{
		Name: "Mexico City",
		Code: "MEX",
		Tmzn: "CST",
	}

	cities[2] = city{
		Name: "Monterrey",
		Code: "MTY",
		Tmzn: "CST",
	}

	cities[3] = city{
		Name: "Cancun",
		Code: "CUN",
		Tmzn: "EST",
	}

	cities[4] = city{
		Name: "Chihuahua",
		Code: "CUU",
		Tmzn: "MST",
	}

	cities[5] = city{
		Name: "Tijuana",
		Code: "TIJ",
		Tmzn: "PST",
	}

	return
}

func getCityIndex(code string) (index int) {
	index = -1
	for i, v := range cities {
		if v.Code == code {
			index = i
		}
	}
	return
}

func getCityCodes() (codes []cityCode) {
	for _, v := range cities {
		codes = append(codes, cityCode{v.Code})
	}
	return
}

func getTime(code string) (data timeStruct) {

	index := getCityIndex(code)
	timezone := cities[index].Tmzn

	resp, err := http.Get("http://worldclockapi.com/api/json/" + timezone + "/now")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &data)
	data.CityName = cities[index].Name

	return data
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{ "Message": "Hello, world!" }`)
	fmt.Println("Endpoint reached: helloWorld")
}

func getCitiesList(w http.ResponseWriter, r *http.Request) {
	codes := getCityCodes()
	jsonObj, _ := json.Marshal(codes)
	fmt.Fprintf(w, string(jsonObj))
	fmt.Println("Endpoint reached: getCitiesList")
}

func getTimeRest(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	code := strings.ToUpper(args["code"])
	resp := getTime(code)
	fmt.Fprintf(w, `{ "Code" : "`+code+`", "Name": "`+resp.CityName+`", "Date & Time": "`+resp.CurrentDateTime+`" }`)
	fmt.Println("Endpoint reached: getTimeRest")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", helloWorld)
	router.HandleFunc("/api/cities", getCitiesList)
	router.HandleFunc("/api/cities/{code}", getTimeRest)
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}
