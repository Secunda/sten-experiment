package steg

import (
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"log"
	"os"
)

// CustomImg - custom image structure
type CustomImg struct {
	image.Image
	setColor map[image.Point]color.Color
}

// InverseImage - method for inverting colors of provided image
func InverseImage(file io.Reader) image.Image {
	baseImage, _, err := image.Decode(file)

	if err != nil {
		log.Fatal(err)
	}

	bounds := baseImage.Bounds()

	customBaseImage := NewCustomImg(baseImage)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldPixel := baseImage.At(x, y)
			r, g, b, a := oldPixel.RGBA()
			r = 65535 - r
			g = 65535 - g
			b = 65535 - b
			pixel := color.RGBA{
				Uint32ToUint8(r),
				Uint32ToUint8(g),
				Uint32ToUint8(b),
				Uint32ToUint8(a),
			}
			customBaseImage.Set(x, y, pixel)
		}
	}

	return customBaseImage
}

// Save image to file
func SaveImageToFile(baseImage image.Image, destinationResult string) {
	newImage, err := os.Create(destinationResult)
	if err != nil {
		log.Fatal(err)
	}

	err = jpeg.Encode(newImage, baseImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		log.Fatal(err)
	}

	err = newImage.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// NewCustomImg - new custom image with ability to change pixel's color
func NewCustomImg(img image.Image) *CustomImg {
	return &CustomImg{img, map[image.Point]color.Color{}}
}

// Set - method for setting pixel's color
func (m *CustomImg) Set(x, y int, c color.Color) {
	m.setColor[image.Point{x, y}] = c
}

// At - method for getting pixel's color
func (m *CustomImg) At(x, y int) color.Color {
	// Explicitly changed part: custom colors of the changed pixels:
	if c := m.setColor[image.Point{x, y}]; c != nil {
		return c
	}
	// Unchanged part: colors of the original image:
	return m.Image.At(x, y)
}

// Uint32ToUint8 - convert uint32 to uint8
func Uint32ToUint8(value uint32) uint8 {
	return uint8(value >> 8)
}
