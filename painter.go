package main

import (
	"github.com/ungerik/go-cairo"
)

//GeneratePoster  Function to generate a poster based on the given details
func GeneratePoster(posterDetails Document) {
	surface, _ := cairo.NewSurfaceFromPNG(posterDetails.Background)
	surface.WriteToPNG(posterDetails.Title + ".png")
}
