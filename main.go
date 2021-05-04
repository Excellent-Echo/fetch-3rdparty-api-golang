package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Data struct {
	// Count    int    `json:"count"`
	// Next     string `json:"next"`
	// Previous string `json:"previous"`
	Results []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	// http.Get // set header
	// response, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=100&offset=200")

	// if err != nil || response.StatusCode != 200 {
	// 	fmt.Println("error fetch API")
	// 	return
	// }

	// if response.StatusCode != 200 { //400 / 500/ 401 dst
	// 	fmt.Println("error get fetch API")
	// 	return
	// }

	// // fmt.Println(response.StatusCode)

	// byteData, err := ioutil.ReadAll(response.Body)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// var datas Data

	// json.Unmarshal(byteData, &datas)

	// for _, data := range datas.Results {
	// 	fmt.Println("name : " + data.Name + " url : " + data.Url)
	// }

	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon?limit=100&offset=200", nil)

	// request.Header.Add("Accept", "application/json")
	// request.Header.Add("Content-Type", "application/json")
	// request.Header.Add("authorization", "secret_key")
	// request.Header.Add("access_token", "secret_key")

	if err != nil {
		fmt.Println("error create request API")
		return
	}

	resp, err := client.Do(request)

	if err != nil || resp.StatusCode != 200 {
		fmt.Println("error fetch API")
		return
	}

	byteData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var datas Data

	json.Unmarshal(byteData, &datas)

	for _, data := range datas.Results {
		fmt.Println("name : " + data.Name + " url : " + data.Url)
	}
}
