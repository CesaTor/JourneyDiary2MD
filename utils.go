package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

func float2str(ff float64) string {
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
