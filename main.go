package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	scanFolder := "../journey/"
	destFolder := "../journals/"
	assetsFolder := "../assets/"
	fileNameFormat := "2006_01_02"
	isLogseq := false
	splitSameDayNotes := false

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
			journal.isLogseq = isLogseq
			journal.date = time.UnixMilli(journal.DateJournal)

			// TEXT
			if journal.Text != "" {
				sb.WriteString(journal.formatText())
			}

			// PHOTOS (actuallu, ASSETS)
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

			// multiple note separator
			sb.WriteString("\n\n---\n\n")

			// Create/Open file and write
			f, _ := os.OpenFile(
				newFileName(destFolder+journal.date.Format(fileNameFormat), splitSameDayNotes)+".md",
				os.O_APPEND|os.O_CREATE, 0666)
			f.WriteString(sb.String())
			f.Close()

		} else if !file.IsDir() {
			copyFileTo(scanFolder+file.Name(), assetsFolder+file.Name())
		}

	}
}
