package main

import (
	"github.com/signintech/gopdf"
)

//GeneratePoster  Function to generate a poster based on the given details
func GeneratePoster(posterDetails Document) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595, H: 842}}) //A4
	pdf.AddPage()

	//Draw Background
	rect := gopdf.Rect{W: 595, H: 842}
	err := pdf.Image(posterDetails.Background, 0, 0, &rect)
	must(err)

	pdf.WritePdf(posterDetails.Title + ".pdf")
}
