package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type App struct {
	Appname    string `json:"appname"`
	Appversion string `json:"appversion"`
}

type Device struct {
	Name 	string `json:"name"`
	Type 	string `json:"interface"`
	Speed 	int    `json:"speed"`
	Port 	string `json:"port"`
}

type Service struct {
	Devices []Device
}

func Parse(path string, v interface{}) {

	jsonFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &v)

	if err != nil {
		log.Fatal(err)
	}
}
