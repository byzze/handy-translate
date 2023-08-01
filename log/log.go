package log

import (
	"fmt"
	"os"
	"time"
)

// Init init log file
func Init() *os.File {
	var tnow = time.Now()
	tnowStr := tnow.Format("20060102")

	// check file stat
	filename := tnowStr + ".log"
	_, err := os.Stat(filename)
	var file *os.File

	if os.IsNotExist(err) {
		file, err = os.Create(filename)
		if err != nil {
			fmt.Println("create fail:", err)
			return nil
		}
		fmt.Println("create file success")
	} else {
		file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("create fail:", err)
			return nil
		}
	}

	return file
}
