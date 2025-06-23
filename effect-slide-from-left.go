package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// SlideFromLeftEffect implements a slide-in transition from the left
type SlideFromLeftEffect struct {
	duration     int
	currentFrame int
	screenWidth  float64
	screenHeight float64
}

func NewSlideFromLeftEffect(durationInFrames int, screenWidth, screenHeight float64) *SlideFromLeftEffect {
	return &SlideFromLeftEffect{
		duration:     durationInFrames,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
}

func (sf *SlideFromLeftEffect) Update() {
	sf.currentFrame++
}

func (sf *SlideFromLeftEffect) Draw(screen *ebiten.Image, current, next *ebiten.Image) {
	// Draw current image
	sf.drawScaledImage(screen, current, 0)

	// Calculate progress and position for the next image
	progress := float64(sf.currentFrame) / float64(sf.duration)
	xOffset := sf.screenWidth * (progress - 1)

	// Draw next image sliding in from left
	sf.drawScaledImage(screen, next, xOffset)
}

func (sf *SlideFromLeftEffect) drawScaledImage(screen *ebiten.Image, img *ebiten.Image, xOffset float64) {
	b := img.Bounds()
	imageWidth := float64(b.Dx())
	imageHeight := float64(b.Dy())

	scaleX := sf.screenWidth / imageWidth
	scaleY := sf.screenHeight / imageHeight
	scale := min(scaleX, scaleY)

	newWidth := imageWidth * scale
	newHeight := imageHeight * scale

	x := (sf.screenWidth - newWidth) / 2
	y := (sf.screenHeight - newHeight) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x+xOffset, y)
	screen.DrawImage(img, op)
}

func (sf *SlideFromLeftEffect) IsComplete() bool {
	return sf.currentFrame >= sf.duration
}

func (sf *SlideFromLeftEffect) Reset() {
	sf.currentFrame = 0
}
