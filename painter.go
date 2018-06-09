package main

import (
	"strings"

	"github.com/ungerik/go-cairo"
)

//GeneratePoster  Function to generate a poster based on the given details
func GeneratePoster(posterDetails Document) {
	surface, _ := cairo.NewSurfaceFromPNG(posterDetails.Background)

	//Draw the title
	surface.MoveTo(150, 150)
	surface.SelectFontFace("opensans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	DrawString(posterDetails.Title, surface, 47.0)

	//Draw a line
	_, y := surface.GetCurrentPoint()
	y += 20
	x := float64(surface.GetWidth() / 10)
	surface.SetLineWidth(3.0)
	surface.MoveTo(x, y)
	surface.SetSourceRGB(0.5, 0.5, 0.5)
	surface.RelLineTo(4*float64(surface.GetWidth()/5), 0)
	surface.Stroke()

	//Draw the Venue Address
	y += 30
	surface.MoveTo(x, y)
	surface.SelectFontFace("ubuntu", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	DrawString(posterDetails.Venue, surface, 20)

	//Saving the Poster
	surface.WriteToPNG(posterDetails.Title + ".png")
}

//DrawString  Function to draw Text Elements
func DrawString(rawtext string, surface *cairo.Surface, fontSize float64) {
	surface.SetFontSize(fontSize)
	surface.SetSourceRGB(0.086, 0.141, 0.173)
	x, y := surface.GetCurrentPoint()
	parsedText := breakText(rawtext)

	width := surface.GetWidth()
	center := width / 2

	for _, text := range parsedText {
		extents := surface.TextExtents(text)
		x = float64(center) - (extents.Width/2 + extents.Xbearing)
		surface.MoveTo(x, y)
		surface.ShowText(text)
		x = 150
		y += extents.Height + 10
	}
}

//breakText  Function to break the Text into proper set of words
func breakText(rawstring string) []string {
	stopWords := []string{"And", "and", "&", "+"}
	tokenizedText := strings.Split(rawstring, " ")
	var organizedText []string

	count := 0

	for i, text := range tokenizedText {
		count++
		if count == 2 {
			organizedText = append(organizedText, tokenizedText[i-1]+
				" "+tokenizedText[i])
			count = 0
		}

		for _, word := range stopWords {
			if text == word {
				organizedText = append(organizedText, word)
				count = 0
			}
		}
	}

	if count == 1 {
		organizedText = append(organizedText,
			tokenizedText[len(tokenizedText)-1])
	}

	return organizedText
}
