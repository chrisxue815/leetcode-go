package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"runtime"
)

var filePattern = regexp.MustCompile(`^p(\d+)`)

func LoadTestJson(testJson interface{}) {
	_, file, _, _ := runtime.Caller(1)
	LoadTestJsonFromFile(file, testJson)
}

func LoadTestJsonFromFile(file string, testJson interface{}) {
	filename := filepath.Base(filepath.Dir(file))
	id := filePattern.FindStringSubmatch(filename)[1]
	jsonFile := fmt.Sprintf("../../leetcode_test_cases/test_%s.json", id)

	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, testJson)
	if err != nil {
		panic(err)
	}
}
