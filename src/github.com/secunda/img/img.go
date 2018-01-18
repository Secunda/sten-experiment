package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/secunda/steg"
)

func main() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	file, err := os.Open("./111.jpg")

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	inversedFile := steg.InverseImage(file)

	fmt.Println(steg.GetRGBAPixels(inversedFile))

	steg.SaveImageToFile(inversedFile, "./inv.jpg")

	// baseImage, _, err := image.Decode(file)
	// oldPixel := baseImage.At(0, 0)
	// r, g, b, a := oldPixel.RGBA()
	// fmt.Println(r, g, b, a)
	// fmt.Println(r, steg.Uint32ToUint8(r))
	// fmt.Println(r+100, steg.Uint32ToUint8(r+100))
	// fmt.Println(uint8(r >> 8))
	// fmt.Println(uint8((r + 100) >> 8))
}
