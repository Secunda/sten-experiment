package steg

import (
	"image"
)

// Pixel struct example
type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

// GetRGBAPixels - get the bi-dimensional pixel array
func GetRGBAPixels(img image.Image) ([][]Pixel, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{
		Uint32ToUint8(r),
		Uint32ToUint8(g),
		Uint32ToUint8(b),
		Uint32ToUint8(a),
	}
}
