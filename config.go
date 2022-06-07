package main

import (
	"strconv"
)

type Config struct {
	ScanFolder        string `json:"scanFolder"`
	DestFolder        string `json:"destFolder"`
	AssetsFolder      string `json:"assetsFolder"`
	FileNameFormat    string `json:"fileNameFormat"`
	IsLogseq          bool   `json:"isLogseq"`
	SplitSameDayNotes bool   `json:"splitSameDayNotes"`
}

func (c Config) printable() string {

	return "\n{\n\t" +
		"ScanFolder: " + c.ScanFolder + ",\n\t" +
		"DestFolder: " + c.DestFolder + ",\n\t" +
		"AssetsFolder: " + c.AssetsFolder + ",\n\t" +
		"FileNameFormat: " + c.FileNameFormat + ",\n\t" +
		"IsLogseq: " + strconv.FormatBool(c.IsLogseq) + ",\n\t" +
		"SplitSameDayNotes: " + strconv.FormatBool(c.SplitSameDayNotes) + ",\n}"
}
