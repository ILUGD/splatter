package main

import (
	"strings"

	"github.com/ungerik/go-cairo"
)

//GeneratePoster  Function to generate a poster based on the given details
func GeneratePoster(posterDetails Document) {
	surface, _ := cairo.NewSurfaceFromPNG(posterDetails.Background)
	DrawTitle(posterDetails.Title, surface)
	surface.WriteToPNG(posterDetails.Title + ".png")
}

//DrawTitle  Function to draw the Title of the Poster
func DrawTitle(title string, surface *cairo.Surface) {
	surface.SelectFontFace("sans-serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surface.SetFontSize(50.0)
	surface.SetSourceRGB(0.086, 0.141, 0.173)
	x, y := 150.00, 198.00
	parsedTitle := breakTitle(title)

	width := surface.GetWidth()
	center := width / 2

	for _, text := range parsedTitle {
		extents := surface.TextExtents(text)
		x = float64(center) - (extents.Width/2 + extents.Xbearing)
		surface.MoveTo(x, y)
		surface.ShowText(text)
		x = 150
		y += extents.Height + 10
	}
}

func breakTitle(title string) []string {
	stopWords := []string{"And", "and", "&", "+"}
	tokenizedTitle := strings.Split(title, " ")
	var organizedTitle []string

	count := 0

	for i, text := range tokenizedTitle {
		count++
		if count == 2 {
			organizedTitle = append(organizedTitle, tokenizedTitle[i-1]+
				" "+tokenizedTitle[i])
			count = 0
		}

		for _, word := range stopWords {
			if text == word {
				organizedTitle = append(organizedTitle, word)
				count = 0
			}
		}
	}

	if count == 1 {
		organizedTitle = append(organizedTitle,
			tokenizedTitle[len(tokenizedTitle)-1])
	}

	return organizedTitle
}
