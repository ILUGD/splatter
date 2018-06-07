package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

//Date  Data Structure for containing the meetup date
type Date struct {
	D string `json:"d"`
	M string `json:"m"`
	Y string `json:"y"`
}

//Time  Data Structure for containing timings
type Time struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

//Logos  Data Structure for containing paths for different logos
type Logos struct {
	Logo []string `json:"logo"`
}

//Websites  Data Structure for containing the different links
type Websites struct {
	Web []string `json:"web"`
}

//Document  Parent Data Structure for the Poster Details
type Document struct {
	Title         string   `json:"title"`
	EventDate     Date     `json:"date"`
	Venue         string   `json:"venue"`
	Timings       Time     `json:"time"`
	Background    string   `json:"background"`
	SponsorLogos  Logos    `json:"logos"`
	GroupWebsites Websites `json:"websites"`
}

//imageDate  Wrapper Data Structure for containing the Document
type imageData struct {
	PosterDetails Document `json:"document"`
}

var jsonFile = flag.String("config", "NULL", "Path for the JSON config")

func main() {
	flag.Parse()

	detailsFile, err := os.Open(*jsonFile)
	defer detailsFile.Close()

	must(err)

	var imageDetails imageData

	bytes, _ := ioutil.ReadAll(detailsFile)
	err = json.Unmarshal(bytes, &imageDetails)

	must(err)

}

func must(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}
