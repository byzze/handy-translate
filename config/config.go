package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

// Data config
var Data config

type (
	config struct {
		Appname      string               `toml:"appname"`
		Translate    map[string]Translate `toml:"translate"`
		TranslateWay string               `toml:"translateway"`
	}

	Translate struct {
		Key    string `toml:"key"`
		Secret string `toml:"secret"`
		Token  string `toml:"token"`
	}
)

// Init  config
func Init() {
	f := "config/config.toml"
	if _, err := os.Stat(f); err != nil {
		logrus.Panic("read file error")
	}
	_, err := toml.DecodeFile(f, &Data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(Data)
}
