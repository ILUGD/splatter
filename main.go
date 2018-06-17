package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ILUGD/splatter/painter"
	"github.com/ILUGD/splatter/readers"
)

var configFlag = flag.String("config", "NULL", "Path for the JSON config")

func main() {
	flag.Parse()

	if *configFlag == "NULL" {
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
