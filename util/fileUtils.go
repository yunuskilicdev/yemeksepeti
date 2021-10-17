package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

const basePath = "/tmp/kvStore"

func DeleteAllFiles() {
	dir, _ := ioutil.ReadDir(basePath)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{basePath, d.Name()}...))
	}
}

func CreateFile(data []byte) {
	filename := fmt.Sprintf("/tmp/kvStore/%v-data.json", time.Now().Unix())
	os.WriteFile(filename, data, 0644)
}

func InitializeBasePath() {
	path := basePath
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ReadJsonFromFile(output *map[string]string) {
	var fileInfo os.FileInfo
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if fileInfo == nil {
			fileInfo = f
			continue
		}
		if f.ModTime().After(fileInfo.ModTime()) {
			fileInfo = f
		}
	}

	if fileInfo != nil {
		data, err := ioutil.ReadFile("/tmp/kvStore/" + fileInfo.Name())
		if err != nil {
			log.Fatal(err)
		}
		jsonErr := json.Unmarshal(data, output)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
	}
}
