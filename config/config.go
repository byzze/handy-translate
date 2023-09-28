package config

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Data config
var Data config

type (
	config struct {
		Appname   string      `toml:"appname"`
		Keyboard  []string    `toml:"keyboard"`
		Translate []Translate `toml:"translate"`
	}

	Translate struct {
		Name  string `toml:"name" json:"name,omitempty"`
		AppID string `toml:"appID" json:"appID,omitempty"`
		Key   string `toml:"key" json:"key,omitempty"`
	}
)

// Init  config
func Init(ctx context.Context) {
	configFile, err := os.Open("./config.toml")
	if err != nil {
		logrus.WithError(err).Error("Open")
		runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "错误",
			Message: "配置文件打开失败:" + err.Error(),
			// DefaultButton: "No",
		})
		os.Exit(1)
	}
	defer configFile.Close()

	fd, err := io.ReadAll(configFile)
	if err != nil {
		logrus.WithError(err).Error("ReadAll")
		os.Exit(1)
	}

	err = toml.Unmarshal(fd, &Data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(Data)
}

func Save() {
	filePath := "./config.toml"
	// 使用Toml库编码并保存数据
	// 使用ioutil.WriteFile()函数将数据写入文件
	// err := ioutil.WriteFile("./config.toml", data, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	data, err := toml.Marshal(&Data)
	if err != nil {
		logrus.WithError(err).Error("Marshal")
	}

	// 使用 os.Remove() 函数删除文件
	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("删除文件时出错:", err)
		return
	}

	fmt.Println("文件已成功删除")

	// 打开文件，如果文件不存在则创建
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("无法写入文件:", err)
		return
	}

	fmt.Println("数据写入成功!")
}
