package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

//doesn't seem to matter if func, variable, type definitions are before or after
//the main() func. This looks similar to hoisting in javascript.
//even moving the Configuration type declaration below the LoadConfiguration() func still works.

func main() {

	fmt.Println("AppLoading...")

	config, err := LoadConfiguration("./config/config.json")
	Errorcheck(err)

	fmt.Println("Application Name: ", config.Application.Name)
	fmt.Println("Logging Dir: ", config.Logging.Dir)

	cmd := exec.Command("powershell.exe", `./scripts/hello-world.ps1 -someArg "jelly"`)
	out, err := cmd.CombinedOutput()
	Errorcheck(err)

	fmt.Print(string(out))
}

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
	Logging struct {
		Dir      string `json:"dir"`
		Filename string `json:"filename"`
	} `json:"logging"`
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
