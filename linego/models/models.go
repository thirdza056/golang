package models

import (
	"../server"
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

/* files */

func SaveFile(path string, data []byte) (bool) {
	err := ioutil.WriteFile(path, data, 0644)
	if err!=nil {
		return false
	}
	return true
}

func DeleteFile(path string) (bool) {
	err := os.Remove(path)
	if err!=nil {
		return false
	}
	return true
}

func DownloadFile(url string) (string) {
	tmpfile, err := ioutil.TempFile("/tmp","LineAR-GOBOT-")
	if err!=nil {
		return "Failed."
	}
	res, err := server.GetContent(url, nil)
	if err!=nil {
		return "Failed."
	}
	if _, err := tmpfile.Write(res.Body); err!=nil {
		return "Failed."
	}
	return tmpfile.Name()
}