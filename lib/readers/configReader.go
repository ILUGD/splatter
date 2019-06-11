/*
 *  Splatter
 *  Copyright (C) 2019 Kuntal Majumder
 *
 *  This program is free software; you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation; either version 2 of the License, or (at
 *  your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful, but
 *  WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 *  General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program; if not, write to the Free Software
 *  Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA 02111-1307
 *  USA.
 */

package readers

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ILUGD/splatter/lib/logger"
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

type ConfigReader struct {
	*logger.Logger
	filename string
}

func (reader *ConfigReader) ReadConfig() Document {

	configFile, err := os.Open(reader.filename)
	defer configFile.Close()

	reader.Must(err, "Read the config file")

	var imageDetails Document

	bytes, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(bytes, &imageDetails)

	reader.Must(err, "Unmarshalled JSON")

	return imageDetails
}

//NewConfigReader  Function for creating a new reader for the config file
func NewConfigReader(filepath string, l *logger.Logger) *ConfigReader {

	reader := ConfigReader{
		l, filepath,
	}

	return &reader
}
