package config

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

//go:embed config.toml
var configData []byte

// Data config
var Data config

type (
	config struct {
		Appname   string      `toml:"appname"`
		Keyboard  []string    `toml:"keyboard"`
		Translate []Translate `toml:"translate"`
	}

	Translate struct {
		Name   string `toml:"name" json:"name"`
		Key    string `toml:"key" json:"key"`
		Secret string `toml:"secret" json:"secret"`
		Token  string `toml:"token" json:"token"`
		AppID  string `toml:"appID" json:"appID"`
	}
)

// Init  config
func Init() {
	f := configData
	err := toml.Unmarshal(f, &Data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(Data)
}
