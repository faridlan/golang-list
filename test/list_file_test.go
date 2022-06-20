package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCreateFile(t *testing.T) {
	const (
		YYYYMMDD = "2006-01-02"
	)
	now := time.Now().UTC()
	nowFormat := now.Format(YYYYMMDD)
	noDash := strings.Replace(nowFormat, "-", "", -1)
	ext := ".json"

	_, err := os.Create("../json/" + noDash + ext)
	if err != nil {
		panic(err)
	}
}

func TestListAgain(t *testing.T) {
	files, err := ioutil.ReadDir("../json")
	if err != nil {
		panic(err)
	}
	var largernumber, temp int

	for _, file := range files {
		x := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		filename, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}

		if filename > temp {
			temp = filename
			largernumber = temp
		}
	}

	fmt.Println(largernumber)
}
