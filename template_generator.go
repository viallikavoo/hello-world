package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

func main() {
	const url = "https://raw.githubusercontent.com/viallikavoo/docker-templates/main/templates/go/Dockerfile.template"

	if len(os.Args) < 2 {
		log.Fatalf("Please provide app name as an argument to the binary. Usage: %s app_name", os.Args[0])
	}
	args := os.Args[1]
	data, err := http.Get(url)
	if err != nil || data.StatusCode != 200 {
		log.Fatalf("Could not read data from %s due to error %s", url, err)
	}
	defer data.Body.Close()
	template, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now() // current local time
	err1 := ioutil.WriteFile("Dockerfile", []byte("# "+now.String()+"\n"+string(bytes.Replace(template, []byte("{APP_NAME}"), []byte(args), -1))), 0777)
	if err1 != nil {
		log.Fatal(err1)
	}

	result, err := ioutil.ReadFile("Dockerfile")
	if err != nil {
		log.Fatal(err)
	}

	yellow := color.New(color.FgYellow)
	boldYellow := yellow.Add(color.Bold)
	boldYellow.Println(string(result))

}
