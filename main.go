package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var baseUrl = "http://localhost:8001"

type mahasiswa struct {
	NPM   string
	Nama  string
	Grade int
}

func fetchUsers() ([]mahasiswa, error) {
	var err error
	var client = &http.Client{}
	var data []mahasiswa

	request, err := http.NewRequest("POST", baseUrl+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	var users, err = fetchUsers()
	if err != nil {
		fmt.Println("error!!!", err.Error())
	}

	for _, each := range users {
		fmt.Printf("NPM : %s, Nama : %s, Grades : %d\n", each.NPM, each.Nama, each.Grade)
	}

}
