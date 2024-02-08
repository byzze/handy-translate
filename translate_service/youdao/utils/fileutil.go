package utils

import (
	base64util "encoding/base64"
	"io/ioutil"
	"log/slog"
	"os"
)

func SaveFile(path string, data []byte, needDecode bool) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if needDecode {
		base64 := string(data)
		data, _ = base64util.StdEncoding.DecodeString(base64)
	}
	if err != nil {
		slog.Error("file create failed. err: ", err)
		return
	}
	file.Write(data)
}

func ReadFileAsBase64(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		slog.Error("file read failed. err: ", err)
		return "", err
	}

	fd, err := ioutil.ReadAll(file)
	if err != nil {
		slog.Error("file read failed. err: ", err)
		return "", err
	}

	return base64util.StdEncoding.EncodeToString(fd), nil
}
