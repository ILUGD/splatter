package painter

import (
	"strconv"
	"strings"

	"github.com/ILUGD/splatter/readers"
	"github.com/ungerik/go-cairo"
)

var months = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

//GeneratePoster  Function to generate a poster based on the given details
func GeneratePoster(posterDetails readers.Document) {
	surface, _ := cairo.NewSurfaceFromPNG(posterDetails.Background)

	//Draw the title
	surface.MoveTo(150, 150)
	surface.SelectFontFace("opensans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	drawString(posterDetails.Title, surface, 47.0)

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
	drawString(posterDetails.Venue, surface, 20)

	//Draw Another Line
	_, y = surface.GetCurrentPoint()
	y += 20
	x = float64(surface.GetWidth() / 4)
	surface.SetLineWidth(2.0)
	surface.MoveTo(x, y)
	surface.SetSourceRGB(0.7, 0.7, 0.7)
	surface.RelLineTo(float64(surface.GetWidth()/2), 0)
	surface.Stroke()

	//Draw the Timings
	text := posterDetails.Timings.Start + " - " + posterDetails.Timings.End
	y += 37
	surface.MoveTo(x, y)
	surface.SetFontSize(25)
	surface.SetSourceRGB(0.3, 0.3, 0.3)
	extents := surface.TextExtents(text)
	width := surface.GetWidth()
	center := width / 2
	x = float64(center) - (extents.Width/2 + extents.Xbearing)
	surface.MoveTo(x, y)
	surface.ShowText(text)

	//Draw the Date
	y += 40
	text = months[posterDetails.EventDate.M] + "  " +
		strconv.Itoa(posterDetails.EventDate.D)
	surface.MoveTo(x, y)
	surface.SetFontSize(25)
	surface.SetSourceRGB(0.3, 0.3, 0.3)
	extents = surface.TextExtents(text)
	width = surface.GetWidth()
	center = width / 2
	x = float64(center) - (extents.Width/2 + extents.Xbearing)
	surface.MoveTo(x, y)
	surface.ShowText(text)

	//Draw Another line
	_, y = surface.GetCurrentPoint()
	y += 20
	x = float64(surface.GetWidth() / 3)
	surface.SetLineWidth(2.0)
	surface.MoveTo(x, y)
	surface.SetSourceRGB(0.5, 0.5, 0.5)
	surface.RelLineTo(float64(surface.GetWidth()/3), 0)
	surface.Stroke()

	//Finally Drawing the Links
	y += 40
	surface.MoveTo(x, y)
	surface.SelectFontFace("roboto", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	drawLinks(posterDetails.GroupWebsites, surface, 17.0)
	//Saving the Poster
	surface.WriteToPNG(posterDetails.Title + ".png")
}

//drawLinks  Function for drawing the hyperlinks
func drawLinks(link []string, surface *cairo.Surface, fontSize float64) {
	surface.SetFontSize(fontSize)
	surface.SetSourceRGB(0, 0, 0)
	x, y := surface.GetCurrentPoint()

	width := surface.GetWidth()
	center := width / 2
	for _, text := range link {
		extents := surface.TextExtents(text)
		x = float64(center) - (extents.Width/2 + extents.Xbearing)
		surface.MoveTo(x, y)
		surface.ShowText(text)
		y += extents.Height + 10
	}
}

//drawString  Function to draw Text Elements
func drawString(rawtext string, surface *cairo.Surface, fontSize float64) {
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
