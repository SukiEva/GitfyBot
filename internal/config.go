package internal

import (
	"encoding/json"
	"io/ioutil"
)

type App struct {
	Github struct {
		Token string `json:"token"`
		URL   string `json:"url"`
	} `json:"github"`
	Database struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"database"`
	Telegram struct {
		Token string `json:"token"`
		Proxy string `json:"proxy"`
	} `json:"telegram"`
	Admin          string `json:"admin"`
	UpdateInterval int    `json:"updateInterval"`
	LogPath        string `json:"logPath"`
}

var Config App

func init() {
	fileData, err := ioutil.ReadFile("configs/config.json")
	dropErr(err)
	err = json.Unmarshal(fileData, &Config)
	dropErr(err)
}

func dropErr(e error) {
	if e != nil {
		panic(e)
	}
}
