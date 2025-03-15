package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func OpenImage(fileName string) image.Image {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("err opening file")
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("err decoding img")
	}

	return img
}

func PixelToAscii(c color.Color) string {
	const shades = " .:-=+*#%@"

	grayPixel := color.GrayModel.Convert(c).(color.Gray)
	pxl := int(grayPixel.Y) * (len(shades) - 1) / 255

	return string(shades[pxl])
}

func PixelToColoredAscii(c color.Color) string {
	const shadows = "@$B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,\"^`'. "
	r, g, b, _ := c.RGBA()
	r, g, b = r>>8, g>>8, b>>8

	grayPixel := color.GrayModel.Convert(c).(color.Gray)
	pxl := int(grayPixel.Y) * (len(shadows) - 1) / 255

	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, string(shadows[len(shadows)-pxl-1]))
}

func PixelToAnsi(c color.Color) string {
	r, g, b, _ := c.RGBA()
	r, g, b = r>>8, g>>8, b>>8

	return fmt.Sprintf("\033[38;2;%d;%d;%dm\033[48;2;%d;%d;%dm%s\033[0m",
		r, g, b, r, g, b, string('A'))
}

func ConvertImage(img image.Image) []string {
	width := 75
	height := int(float64(img.Bounds().Dy()) / float64(img.Bounds().Dx()) * float64(width) * 0.55)
	resized := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	var ansiImage []string

	for y := range resized.Bounds().Dy() {
		var row string
		for x := range resized.Bounds().Dx() {
			row += PixelToAnsi(resized.At(x, y))
		}
		ansiImage = append(ansiImage, row)
	}

	return ansiImage
}
