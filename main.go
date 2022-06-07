package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

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

func newFileName(filename string) string {

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

func main() {
	// TODO - error handling
	// TODO - add to cli
	scanFolder := "../journey/"
	destFolder := "../journals/"
	assetsFolder := "../assets/"

	// Read directory
	files, _ := ioutil.ReadDir(scanFolder)

	// Create Folders
	createFolder(destFolder)
	createFolder(assetsFolder)

	for _, file := range files {

		// Handle .json files in folder
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {

			var sb strings.Builder

			// Get journal data
			journal := journalFromFilepath(scanFolder + file.Name())

			// Create/Open File
			tm := time.UnixMilli(journal.DateJournal)

			f, _ := os.OpenFile(
				newFileName(destFolder+tm.Format("2006_01_02"))+".md", os.O_RDWR|os.O_CREATE, 0666)

			defer f.Close()

			// TEXT
			if journal.Text != "" {
				sb.WriteString(journal.formatText())
			}

			// PHOTOS
			if len(journal.Photos) != 0 {
				sb.WriteString(journal.formatPhotos(assetsFolder))
			}

			// TAGS
			if len(journal.Tags) != 0 {
				sb.WriteString(journal.formatTags())
			}

			// POSITION
			if journal.Address != "" || journal.Timezone != "" {
				sb.WriteString(journal.formatPosition())
			}

			// WEATHER
			if journal.Weather.Description != "" {
				sb.WriteString(journal.formatWeather())
			}

			// Write string
			f.WriteString(sb.String())

		} else if !file.IsDir() {
			// Move file to asset folder
			from, err := os.Open(scanFolder + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer from.Close()

			to, err := os.OpenFile(assetsFolder+file.Name(), os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				log.Fatal(err)
			}
			defer to.Close()

			_, err = io.Copy(to, from)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}
