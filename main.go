package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func importConfig() Config {
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	dayJson := string(content)
	var config Config
	json.Unmarshal([]byte(dayJson), &config)

	return config
}

func main() {

	// Import Configurations from file
	config := importConfig()
	fmt.Println("running config:", config.printable())

	// Read directory
	files, _ := ioutil.ReadDir(config.ScanFolder)

	// Create Folders
	createFolder(config.DestFolder)
	createFolder(config.AssetsFolder)

	for _, file := range files {

		// Handle .json files in folder
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {

			var sb strings.Builder

			// Get journal data
			journal := journalFromFilepath(config.ScanFolder + file.Name())
			journal.isLogseq = config.IsLogseq
			journal.date = time.UnixMilli(journal.DateJournal)

			// TEXT
			if journal.Text != "" {
				sb.WriteString(journal.formatText())
			}

			// PHOTOS (actuallu, ASSETS)
			if len(journal.Photos) != 0 {
				sb.WriteString(journal.formatPhotos(config.AssetsFolder))
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
				newFileName(config.DestFolder+journal.date.Format(config.FileNameFormat), config.SplitSameDayNotes)+".md",
				os.O_APPEND|os.O_CREATE, 0666)
			f.WriteString(sb.String())
			f.Close()

		} else if !file.IsDir() {
			copyFileTo(config.ScanFolder+file.Name(), config.AssetsFolder+file.Name())
		}

	}

	fmt.Println("Done!")
}
