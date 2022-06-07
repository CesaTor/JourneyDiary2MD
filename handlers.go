package main

import (
	"strconv"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func (j Journal) formatText() string {
	converter := md.NewConverter("", true, nil)
	var sb strings.Builder

	sb.WriteString("- #diary\n")

	txt, _ := converter.ConvertString(j.Text) // parse html to markdown

	txt = strings.ReplaceAll(txt, "\n", "\n\t- ")
	// TODO - only for your use case
	// txt = strings.ReplaceAll(txt, "- \\", "\t")
	txt = strings.ReplaceAll(txt, "\\-", "-")
	txt = strings.ReplaceAll(txt, "\\[\\[", "[[")
	txt = strings.ReplaceAll(txt, "\\]\\]", "]]")

	txt = strings.ReplaceAll(txt, " -", "\t")

	sb.WriteString("\t- " + txt)
	sb.WriteString("\n")

	return sb.String()
}

func (j Journal) formatPhotos(assetsFolder string) string {

	var sb strings.Builder
	sb.WriteString("\n")

	// add #photos
	sb.WriteString("- #photos\n")
	// add photos
	for _, photo := range j.Photos {
		sb.WriteString("\t- ![" + photo + "](" + assetsFolder + photo + ")\n")
	}
	sb.WriteString("\n")

	return sb.String()
}

func (j Journal) formatTags() string {

	var sb strings.Builder

	sb.WriteString("\n")
	// add #tags
	sb.WriteString("- Tags:\n\t")
	// add photos
	sb.WriteString("- ")

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
		sb.WriteString("- Location:\n\t" + j.Address + "\n\t" + j.Timezone + "\n\t" + float2str(j.Lat) + " - " + float2str(j.Lon))
	} else if j.Address != "" {
		sb.WriteString("- Location:\n\t" + j.Address + "\n\t" + float2str(j.Lat) + " - " + float2str(j.Lon))
	}
	sb.WriteString("\n")

	return sb.String()
}

func (j Journal) formatWeather() string {
	var sb strings.Builder

	sb.WriteString("\n")

	weather := j.Weather
	if weather.Description != "" && weather.Place != "" {
		sb.WriteString("- Weather:\n\t" + weather.Description + " - " + weather.Place + " - " + strconv.Itoa(weather.DegreeC))
	} else if weather.Description != "" {
		sb.WriteString("- Weather:\n\t" + weather.Description + " - " + strconv.Itoa(weather.DegreeC))
	} else if weather.Place != "" {
		sb.WriteString("- Weather:\n\t" + weather.Place + " - " + strconv.Itoa(weather.DegreeC))
	}
	sb.WriteString("\n")

	return sb.String()

}
