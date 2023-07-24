package utils

import (
	base64util "encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func SaveFile(path string, data []byte, needDecode bool) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if needDecode {
		base64 := string(data)
		data, _ = base64util.StdEncoding.DecodeString(base64)
	}
	if err != nil {
		fmt.Print("file create failed. err: " + err.Error())
	} else {
		file.Write(data)
	}
}

func ReadFileAsBase64(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Print("file read failed. err: " + err.Error())
		return "", err
	} else {
		fd, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Print("file read failed. err: " + err.Error())
			return "", err
		} else {
			return base64util.StdEncoding.EncodeToString(fd), nil
		}
	}
}
