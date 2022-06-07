package main

type Journal struct {
	Text            string  `json:"text"`
	DateModified    int64   `json:"date_modified"` // not_used
	DateJournal     int64   `json:"date_journal"`  // title
	Id              string  `json:"id"`            // not_used
	PreviewText     string  `json:"preview_text"`  // not_used
	Address         string  `json:"address"`
	MusicArtist     string  `json:"music_artist"` // no_idea not_used
	MusicTitle      string  `json:"music_title"`  // no_idea not_used
	Lat             float64 `json:"lat"`
	Lon             float64 `json:"lon"`
	Mood            int     `json:"mood"`
	Label           string  `json:"label"`  // no_idea not_used
	Folder          string  `json:"folder"` // no_idea not_used
	Sentiment       int     `json:"sentiment"`
	Timezone        string  `json:"timezone"`
	Favourite       bool    `json:"favourite"`         // no_idea not_used
	Type            string  `json:"type"`              // not_used
	LinkedAccountId string  `json:"linked_account_id"` // not_used

	Weather Weather `json:"weather"`

	Photos []string `json:"photos"`
	Tags   []string `json:"tags"`
}

type Weather struct {
	Id          int    `json:"id"`
	DegreeC     int    `json:"degree_c"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Place       string `json:"place"`
}
