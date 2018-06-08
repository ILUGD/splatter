package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
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

var configFlag = flag.String("config", "NULL", "Path for the JSON config")

func main() {
	flag.Parse()

	if *configFlag == "NULL" {
		must(errors.New("Forgot to Enter the File Location?" +
			"Hint : Use the -config flag"))
	}

	configFile, err := os.Open(*configFlag)
	defer configFile.Close()

	must(err)

	var imageDetails Document

	bytes, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(bytes, &imageDetails)

	must(err)

	GeneratePoster(imageDetails)
}

//must  Function for handling errors
func must(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}
