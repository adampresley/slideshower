package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// CrossFadeEffect implements a cross-fade transition
type CrossFadeEffect struct {
	duration     int
	currentFrame int
	currentAlpha float64
	nextAlpha    float64
	screenWidth  float64
	screenHeight float64
}

func NewCrossFadeEffect(durationInFrames int, screenWidth, screenHeight float64) *CrossFadeEffect {
	return &CrossFadeEffect{
		duration:     durationInFrames,
		currentAlpha: 1,
		nextAlpha:    0,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
}

func (cf *CrossFadeEffect) Update() {
	progress := min(float64(cf.currentFrame)/float64(cf.duration), 1)

	cf.currentAlpha = 1.0 - progress
	cf.nextAlpha = progress
	cf.currentFrame++
}

func (cf *CrossFadeEffect) Draw(screen *ebiten.Image, current, next *ebiten.Image) {
	// Draw current image with decreasing alpha
	cf.drawScaledImage(screen, current, cf.currentAlpha)

	// Draw next image with increasing alpha
	cf.drawScaledImage(screen, next, cf.nextAlpha)
}

func (cf *CrossFadeEffect) drawScaledImage(screen *ebiten.Image, img *ebiten.Image, alpha float64) {
	b := img.Bounds()
	imageWidth := float64(b.Dx())
	imageHeight := float64(b.Dy())

	scaleX := cf.screenWidth / imageWidth
	scaleY := cf.screenHeight / imageHeight
	scale := min(scaleX, scaleY)

	newWidth := imageWidth * scale
	newHeight := imageHeight * scale

	x := (cf.screenWidth - newWidth) / 2
	y := (cf.screenHeight - newHeight) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleAlpha(float32(alpha))
	screen.DrawImage(img, op)
}

func (cf *CrossFadeEffect) IsComplete() bool {
	return cf.currentFrame >= cf.duration
}

func (cf *CrossFadeEffect) Reset() {
	cf.currentFrame = 0
	cf.currentAlpha = 1
	cf.nextAlpha = 0
}
