package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

var allEffectsForRandom = []string{
	"fade",
	"crossfade",
	"slide-from-left",
	"slide-from-right",
	"spiral-wipe",
	"bubble-melt",
}

func main() {
	config := LoadConfig()
	imagePaths := getImagePaths()

	if len(imagePaths) == 0 {
		log.Fatal("No images found in the current directory")
	}

	if config.SpeedInSeconds < 2 {
		config.SpeedInSeconds = 5
	}

	config.ScreenWidth, config.ScreenHeight = ebiten.Monitor().Size()
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Slide Shower")

	if config.Fullscreen {
		ebiten.SetFullscreen(true)
	}

	app := &App{
		ScreenWidth:  config.ScreenWidth,
		ScreenHeight: config.ScreenHeight,
		Slideshow: NewSlideshow(SlideshowConfig{
			ImagePaths:     imagePaths,
			ScreenWidth:    float64(config.ScreenWidth),
			ScreenHeight:   float64(config.ScreenHeight),
			SpeedInSeconds: config.SpeedInSeconds,
			Effect:         getEffectByName(config.Effect, float64(config.ScreenWidth), float64(config.ScreenHeight)),
		}),
	}

	if err := ebiten.RunGame(app); err != nil {
		log.Fatalf("failed to run game: %v", err)
	}
}

func getImagePaths() []string {
	imagePaths := []string{}
	dirs, err := os.ReadDir("./")

	if err != nil {
		log.Fatalf("failed to read directory: %v", err)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}

		// Only support images (jpeg, png)
		ext := strings.ToLower(filepath.Ext(dir.Name()))

		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
			imagePaths = append(imagePaths, "./"+dir.Name())
		}
	}

	return imagePaths
}

func getEffectByName(name string, screenWidth, screenHeight float64) Effect {
	var effect Effect

	switch name {
	case "fade":
		effect = NewFadeEffect(60, screenWidth, screenHeight)
	case "crossfade":
		effect = NewCrossFadeEffect(60, screenWidth, screenHeight)
	case "slide-from-left":
		effect = NewSlideFromLeftEffect(60, screenWidth, screenHeight)
	case "slide-from-right":
		effect = NewSlideFromRightEffect(60, screenWidth, screenHeight)
	case "alternating-slide":
		effect = NewAlternatingSlideEffect(60, screenWidth, screenHeight)
	case "spiral-wipe":
		effect = NewSpiralWipeEffect(40, screenWidth, screenHeight, 4)
	case "bubble-melt":
		effect = NewBubbleMeltShaderEffect(120, int(screenWidth), int(screenHeight), 50)
	case "random":
		effect = NewRandomEffect(60, screenWidth, screenHeight, allEffectsForRandom)
	default:
		effect = NewFadeEffect(60, screenWidth, screenHeight)
	}

	return effect
}
