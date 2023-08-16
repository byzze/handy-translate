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
		Appname      string               `toml:"appname"`
		Translate    map[string]Translate `toml:"translate"`
		TranslateWay string               `toml:"translateway"`
	}

	Translate struct {
		Name   string `toml:"name"`
		Key    string `toml:"key"`
		Secret string `toml:"secret"`
		Token  string `toml:"token"`
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
