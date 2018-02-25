package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	url    string `json: "url"`
	port   int    `json: "port"`
	dburl  string `json: "dburl"`
	dbname string `json: "dbmame"`
}

//create a new global config:
var config configuration

func main() {

	//get the configuration file from disk:
	file, err := os.Open("./config/config.json")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.port)

}
