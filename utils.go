package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func parseHhMm(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}

func parseFloatDecimal(ff float64) string {
	return strconv.FormatFloat(ff, 'E', -1, 64)
}

func journalFromFilepath(path string) Journal {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	dayJson := string(content)
	var journal Journal
	json.Unmarshal([]byte(dayJson), &journal)
	return journal
}

func createFolder(folder string) {
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func copyFileTo(src string, dest string) {
	// Copy file to asset folder
	from, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}

	to.Close()
}

func newFileName(filename string, split bool) string {

	if !split {
		return filename
	}
	i := 0
	for {
		if i == 0 {
			if !fileExists(filename + ".md") {
				return filename
			}
			i++
		} else {
			if !fileExists(filename + "_" + strconv.Itoa(i) + ".md") {
				return filename + "_" + strconv.Itoa(i)
			}
			i++
		}

		log.Println(i)
	}

}
