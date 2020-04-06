package main

import (
	"fmt"
	"io/ioutil"
	"log"

	pigo "github.com/esimov/pigo/core"
)

func main() {
	cascadeFile, err := ioutil.ReadFile("cascade/facefinder")
	if err != nil {
		log.Fatalf("Error reading the cascade file: %v", err)
	}
	src, err := pigo.GetImage("1.jpg")
	if err != nil {
		log.Fatalf("Cannot open the image file: %v", err)
	}

	pixels := pigo.RgbToGrayscale(src)
	cols, rows := src.Bounds().Max.X, src.Bounds().Max.Y

	cParams := pigo.CascadeParams{
		MinSize:     20,
		MaxSize:     1000,
		ShiftFactor: 0.1,
		ScaleFactor: 1.1,

		ImageParams: pigo.ImageParams{
			Pixels: pixels,
			Rows:   rows,
			Cols:   cols,
			Dim:    cols,
		},
	}

	pigo := pigo.NewPigo()
	classifier, err := pigo.Unpack(cascadeFile)
	if err != nil {
		log.Fatalf("Error reading the cascade file: %s", err)
	}

	angle := 0.0
	dets := classifier.RunCascade(cParams, angle)
	fmt.Println(dets)
}
