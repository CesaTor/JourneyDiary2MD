package main

import (
	"path/filepath"
	"strconv"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func (j Journal) prefix() string {
	if j.isLogseq {
		return "- "
	}
	return ""
}

func (j Journal) formatText() string {
	converter := md.NewConverter("", true, nil)
	var sb strings.Builder

	sb.WriteString(j.prefix() + "#diary " + parseHhMm(j.date.Hour()) + ":" + parseHhMm(j.date.Minute()) + "\n")

	txt, _ := converter.ConvertString(j.Text) // parse html to markdown

	txt = strings.ReplaceAll(txt, "\n", "\n\t"+j.prefix())

	txt = strings.ReplaceAll(txt, "\\-", "-")
	txt = strings.ReplaceAll(txt, "\\_", "_") //2019_10_06

	// For your use case, maybe not for everyone
	// backlinking fix
	txt = strings.ReplaceAll(txt, "\\[\\[", "[[")
	txt = strings.ReplaceAll(txt, "\\]\\]", "]]")
	// fix "-     text" case
	txt = strings.ReplaceAll(txt, "- -", "\t- ") // 2022_05_27

	sb.WriteString("\t" + j.prefix() + txt)
	sb.WriteString("\n")

	return sb.String()
}

func (j Journal) formatPhotos(assetsFolder string) string {

	var sb strings.Builder
	sb.WriteString("\n")

	// add #photos
	var tags strings.Builder
	tags.WriteString(j.prefix() + "#assets")

	var photos strings.Builder
	// add photos
	for _, photo := range j.Photos {

		switch filepath.Ext(photo) {
		case ".jpg", ".png", ".jpeg", ".webp":
			{
				if !strings.Contains(tags.String(), "photo") {
					tags.WriteString(" #photo")
				}
			}
		case ".mp3", ".aac", ".webm":
			{
				if !strings.Contains(tags.String(), "audio") {
					tags.WriteString(" #audio")
				}
			}
		}
		tags.WriteString("\n")

		photos.WriteString("\t" + j.prefix() + "![" + photo + "](" + assetsFolder + photo + ")\n")
	}

	sb.WriteString(tags.String())
	sb.WriteString(photos.String())
	sb.WriteString("\n")

	return sb.String()
}

func (j Journal) formatTags() string {

	var sb strings.Builder

	sb.WriteString("\n")
	// add #tags
	sb.WriteString(j.prefix() + "Tags:\n\t")
	// add photos
	sb.WriteString(j.prefix())

	for _, tag := range j.Tags {
		sb.WriteString("#" + tag + " ")
	}
	sb.WriteString("\n")

	return sb.String()
}

func (j Journal) formatPosition() string {

	var sb strings.Builder

	sb.WriteString("\n")
	if j.Timezone != "" && j.Address != "" {
		sb.WriteString(j.prefix() + "Location:\n\t" + j.Address + "\n\t" + j.Timezone + "\n\t" + parseFloatDecimal(j.Lat) + " - " + parseFloatDecimal(j.Lon))
	} else if j.Address != "" {
		sb.WriteString(j.prefix() + "Location:\n\t" + j.Address + "\n\t" + parseFloatDecimal(j.Lat) + " - " + parseFloatDecimal(j.Lon))
	}
	sb.WriteString("\n")

	return sb.String()
}

func (j Journal) formatWeather() string {
	var sb strings.Builder

	sb.WriteString("\n")

	weather := j.Weather
	if weather.Description != "" && weather.Place != "" {
		sb.WriteString(j.prefix() + "Weather:\n\t" + weather.Description + " - " + weather.Place + " - " + strconv.Itoa(weather.DegreeC))
	} else if weather.Description != "" {
		sb.WriteString(j.prefix() + "Weather:\n\t" + weather.Description + " - " + strconv.Itoa(weather.DegreeC))
	} else if weather.Place != "" {
		sb.WriteString(j.prefix() + "Weather:\n\t" + weather.Place + " - " + strconv.Itoa(weather.DegreeC))
	}
	sb.WriteString("\n")

	return sb.String()

}
