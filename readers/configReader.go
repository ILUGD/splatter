package readers

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//Date  Data Structure for containing the Meetup Date
type Date struct {
	D int `json:"d"`
	M int `json:"m"`
	Y int `json:"y"`
}

//Time  Data Structure for containing Timings
type Time struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

//Document  Parent Data Structure for the Poster Details
type Document struct {
	Title         string   `json:"title"`
	EventDate     Date     `json:"date"`
	Venue         string   `json:"venue"`
	Timings       Time     `json:"time"`
	Background    string   `json:"background"`
	SponsorLogos  []string `json:"logos"`
	GroupWebsites []string `json:"websites"`
}

type configReader struct {
	filename string
}

func (reader *configReader) ReadConfig() (Document, error) {

	configFile, err := os.Open(reader.filename)
	defer configFile.Close()

	if err != nil {
		return Document{}, err
	}

	var imageDetails Document

	bytes, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(bytes, &imageDetails)

	if err != nil {
		return Document{}, err
	}

	return imageDetails, nil
}

//NewConfigReader  Function for creating a new reader for the config file
func NewConfigReader(filepath string) *configReader {

	reader := configReader{
		filename: filepath,
	}

	return &reader
}
