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

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ILUGD/splatter/lib/painter"
	"github.com/ILUGD/splatter/lib/readers"
)

var configFlag = flag.String("config", "", "Path for the JSON config")

func main() {
	flag.Parse()

	if *configFlag == "" {
		must(errors.New("Forgot to Enter the File Location?" +
			"Hint : Use the -config flag"))
	}

	configReader := readers.NewConfigReader(*configFlag)

	imageDetails, err := configReader.ReadConfig()

	must(err)

	painter.GeneratePoster(imageDetails)
}

//must  Function for handling errors
func must(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}
