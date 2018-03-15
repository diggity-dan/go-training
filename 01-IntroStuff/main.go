package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

//Configuration Exported configuration struct
type Configuration struct {
	Database struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"database"`
	Application struct {
		Name string `json:"name"`
		URL  string `json:"url"`
		Port int    `json:"port"`
	} `json:"application"`
}

//Errorcheck Exported DRY error checker
func Errorcheck(e error) {
	if e != nil {
		fmt.Println("ERROR!", e)
		panic(e)
	}
}

//LoadConfiguration Exported JSON configuration loader
func LoadConfiguration(filename string) (Configuration, error) {
	var config Configuration
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return config, err
	}

	jsonparser := json.NewDecoder(file)
	err = jsonparser.Decode(&config)

	return config, err
}

func main() {

	fmt.Println("AppLoading...")

	config, err := LoadConfiguration("./config/config.json")
	Errorcheck(err)

	fmt.Println("Application Name: ", config.Application.Name)

	cmd := exec.Command("powershell.exe", `./scripts/hello-world.ps1 -someArg "jelly"`)
	out, err := cmd.CombinedOutput()

	fmt.Print(string(out))

}
