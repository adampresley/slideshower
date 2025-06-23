package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// AlternatingSlideEffect implements a slide-in transition that alternates between left and right
type AlternatingSlideEffect struct {
	duration     int
	currentFrame int
	screenWidth  float64
	screenHeight float64
	slideRight   bool // true if sliding from left to right, false if sliding from right to left
}

func NewAlternatingSlideEffect(durationInFrames int, screenWidth, screenHeight float64) *AlternatingSlideEffect {
	return &AlternatingSlideEffect{
		duration:     durationInFrames,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		slideRight:   true, // Start with sliding from left to right
	}
}

func (as *AlternatingSlideEffect) Update() {
	as.currentFrame++
}

func (as *AlternatingSlideEffect) Draw(screen *ebiten.Image, current, next *ebiten.Image) {
	// Draw current image
	as.drawScaledImage(screen, current, 0)

	// Calculate progress and position for the next image
	progress := float64(as.currentFrame) / float64(as.duration)
	var xOffset float64
	if as.slideRight {
		xOffset = as.screenWidth * (progress - 1) // Slide from left
	} else {
		xOffset = as.screenWidth * (1 - progress) // Slide from right
	}

	// Draw next image sliding in
	as.drawScaledImage(screen, next, xOffset)
}

func (as *AlternatingSlideEffect) drawScaledImage(screen *ebiten.Image, img *ebiten.Image, xOffset float64) {
	b := img.Bounds()
	imageWidth := float64(b.Dx())
	imageHeight := float64(b.Dy())

	scaleX := as.screenWidth / imageWidth
	scaleY := as.screenHeight / imageHeight
	scale := min(scaleX, scaleY)

	newWidth := imageWidth * scale
	newHeight := imageHeight * scale

	x := (as.screenWidth - newWidth) / 2
	y := (as.screenHeight - newHeight) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x+xOffset, y)
	screen.DrawImage(img, op)
}

func (as *AlternatingSlideEffect) IsComplete() bool {
	return as.currentFrame >= as.duration
}

func (as *AlternatingSlideEffect) Reset() {
	as.currentFrame = 0
	as.slideRight = !as.slideRight // Toggle the slide direction for the next transition
}
