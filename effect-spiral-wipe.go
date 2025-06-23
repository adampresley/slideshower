package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpiralWipeEffect struct {
	duration     int
	currentFrame int
	screenWidth  float64
	screenHeight float64
	spiralMask   *ebiten.Image
	blockSize    int
}

func NewSpiralWipeEffect(durationInFrames int, screenWidth, screenHeight float64, blockSize int) *SpiralWipeEffect {
	effect := &SpiralWipeEffect{
		duration:     durationInFrames,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		blockSize:    blockSize,
	}
	effect.generateSpiralMask()
	return effect
}

func (sw *SpiralWipeEffect) Update() {
	sw.currentFrame++
}

func (sw *SpiralWipeEffect) IsComplete() bool {
	return sw.currentFrame >= sw.duration
}

func (sw *SpiralWipeEffect) Reset() {
	sw.currentFrame = 0
}

func (sw *SpiralWipeEffect) Draw(screen *ebiten.Image, current, next *ebiten.Image) {
	progress := float64(sw.currentFrame) / float64(sw.duration)
	if progress > 1 {
		progress = 1
	}

	// First half: fade out current image
	if progress < 0.5 {
		sw.drawScaledImage(screen, current)
		sw.drawSpiralMask(screen, progress*2, color.Black)
	} else {
		// Second half: fade in next image
		sw.drawScaledImage(screen, next)
		sw.drawSpiralMask(screen, 2*(1.0-progress), color.Black)
	}
}

func (sw *SpiralWipeEffect) drawScaledImage(screen *ebiten.Image, img *ebiten.Image) {
	b := img.Bounds()
	imageWidth := float64(b.Dx())
	imageHeight := float64(b.Dy())

	scaleX := sw.screenWidth / imageWidth
	scaleY := sw.screenHeight / imageHeight
	scale := math.Min(scaleX, scaleY)

	newWidth := imageWidth * scale
	newHeight := imageHeight * scale

	x := (sw.screenWidth - newWidth) / 2
	y := (sw.screenHeight - newHeight) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	screen.DrawImage(img, op)
}

func (sw *SpiralWipeEffect) generateSpiralMask() {
	w := int(sw.screenWidth)
	h := int(sw.screenHeight)
	sw.spiralMask = ebiten.NewImage(w, h)

	centerX := sw.screenWidth / 2
	centerY := sw.screenHeight / 2
	maxRadius := math.Sqrt(centerX*centerX + centerY*centerY)

	for y := 0; y < h; y += sw.blockSize {
		for x := 0; x < w; x += sw.blockSize {
			// Use the center of the block
			cx := float64(x + sw.blockSize/2)
			cy := float64(y + sw.blockSize/2)

			dx := cx - centerX
			dy := cy - centerY
			distance := math.Sqrt(dx*dx + dy*dy)
			angle := math.Atan2(dy, dx)
			if angle < 0 {
				angle += 2 * math.Pi
			}

			value := (distance/maxRadius + 5*angle/(2*math.Pi)) * 255
			value = math.Mod(value, 255)
			gray := uint8(value)

			// Fill entire block with gray value
			for dy := 0; dy < sw.blockSize; dy++ {
				for dx := 0; dx < sw.blockSize; dx++ {
					ix := x + dx
					iy := y + dy
					if ix < w && iy < h {
						sw.spiralMask.Set(ix, iy, color.RGBA{gray, gray, gray, 255})
					}
				}
			}
		}
	}
}

func (sw *SpiralWipeEffect) drawSpiralMask(screen *ebiten.Image, progress float64, c color.Color) {
	threshold := uint8(255 * progress)
	mask := ebiten.NewImage(int(sw.screenWidth), int(sw.screenHeight))

	for y := 0; y < int(sw.screenHeight); y++ {
		for x := 0; x < int(sw.screenWidth); x++ {
			r, _, _, _ := sw.spiralMask.At(x, y).RGBA()
			r8 := uint8(r >> 8)
			if r8 < threshold {
				mask.Set(x, y, c)
			}
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleWithColor(c)
	op.Blend = ebiten.BlendSourceOver
	screen.DrawImage(mask, op)
}
