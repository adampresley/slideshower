package main

import (
	// "image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	// "os"

	"github.com/disintegration/imaging"
	"github.com/hajimehoshi/ebiten/v2"
)

func mustLoadImage(name string) *ebiten.Image {
	// f, err := os.Open(name)
	//
	// if err != nil {
	// 	log.Fatalf("failed to open image: %v", name)
	// }
	//
	// defer f.Close()
	//
	// img, _, err := image.Decode(f)
	// if err != nil {
	// 	log.Fatalf("failed to decode image: %v", name)
	// }

	img, err := imaging.Open(name, imaging.AutoOrientation(true))

	if err != nil {
		log.Fatalf("failed to open image: %v", name)
	}

	return ebiten.NewImageFromImage(img)
}
