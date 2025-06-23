package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// FadeEffect implements a fade transition through black
type FadeEffect struct {
	duration     int
	currentFrame int
	alpha        float64
	screenWidth  float64
	screenHeight float64
}

func NewFadeEffect(durationInFrames int, screenWidth, screenHeight float64) *FadeEffect {
	return &FadeEffect{
		duration:     durationInFrames,
		alpha:        0,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
}

func (f *FadeEffect) Update() {
	f.currentFrame++
	f.alpha = min(float64(f.currentFrame)/float64(f.duration), 1)
}

func (f *FadeEffect) Draw(screen *ebiten.Image, current, next *ebiten.Image) {
	if f.alpha < 0.5 {
		// First half: fade out current image to black
		f.drawScaledImage(screen, current, 1-2*f.alpha)
	} else {
		// Second half: fade in next image from black
		f.drawScaledImage(screen, next, 2*f.alpha-1)
	}

	// Draw black overlay
	blackAlpha := min(2*f.alpha, 2-2*f.alpha)
	f.drawBlackOverlay(screen, blackAlpha)
}

func (f *FadeEffect) drawScaledImage(screen *ebiten.Image, img *ebiten.Image, alpha float64) {
	b := img.Bounds()
	imageWidth := float64(b.Dx())
	imageHeight := float64(b.Dy())

	scaleX := f.screenWidth / imageWidth
	scaleY := f.screenHeight / imageHeight
	scale := min(scaleX, scaleY)

	newWidth := imageWidth * scale
	newHeight := imageHeight * scale

	x := (f.screenWidth - newWidth) / 2
	y := (f.screenHeight - newHeight) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	op.ColorScale.Scale(1, 1, 1, float32(alpha))
	screen.DrawImage(img, op)
}

func (f *FadeEffect) drawBlackOverlay(screen *ebiten.Image, alpha float64) {
	blackImg := ebiten.NewImage(int(f.screenWidth), int(f.screenHeight))
	blackImg.Fill(color.Black)

	op := &ebiten.DrawImageOptions{}
	op.ColorScale.Scale(1, 1, 1, float32(alpha))
	screen.DrawImage(blackImg, op)
}

func (f *FadeEffect) IsComplete() bool {
	return f.currentFrame >= f.duration
}

func (f *FadeEffect) Reset() {
	f.currentFrame = 0
	f.alpha = 0
}
