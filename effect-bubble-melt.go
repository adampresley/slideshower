package main

import (
	_ "embed"
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed bubble_melt.kage
var bubbleMeltShaderSrc []byte

type Bubble struct {
	x, y, radius float32
}

type BubbleMeltShaderEffect struct {
	duration      int
	currentFrame  int
	width, height int
	bubbles       []Bubble
	numBubbles    int
	shader        *ebiten.Shader
}

func NewBubbleMeltShaderEffect(duration, width, height, numBubbles int) *BubbleMeltShaderEffect {
	shader, err := ebiten.NewShader(bubbleMeltShaderSrc)
	if err != nil {
		log.Fatalf("Failed to compile shader: %v", err)
	}

	effect := &BubbleMeltShaderEffect{
		duration:   duration,
		width:      width,
		height:     height,
		numBubbles: numBubbles,
		shader:     shader,
		bubbles:    make([]Bubble, numBubbles),
	}
	effect.generateBubbles()
	return effect
}

func (e *BubbleMeltShaderEffect) generateBubbles() {
	for i := range e.bubbles {
		e.bubbles[i] = Bubble{
			x:      rand.Float32() * float32(e.width),
			y:      rand.Float32() * float32(e.height),
			radius: float32(rand.Float64()*40 + 20),
		}
	}
}

func (e *BubbleMeltShaderEffect) Update() {
	e.currentFrame++
}

func (e *BubbleMeltShaderEffect) Reset() {
	e.currentFrame = 0
	e.generateBubbles()
}

func (e *BubbleMeltShaderEffect) IsComplete() bool {
	return e.currentFrame >= e.duration
}

func (e *BubbleMeltShaderEffect) Draw(screen, current, next *ebiten.Image) {
	progress := min(float32(e.currentFrame)/float32(e.duration), 1)

	// Prepare bubble uniform array
	bubbleData := [100][3]float32{}

	for i, b := range e.bubbles {
		sizeProgress := progress
		if progress < 0.5 {
			sizeProgress = progress / 0.5 // erase phase
		} else {
			sizeProgress = (1.0 - progress) / 0.5 // reveal phase
		}
		if i < len(bubbleData) {
			bubbleData[i] = [3]float32{b.x, b.y, b.radius * sizeProgress}
		}
	}

	op := &ebiten.DrawRectShaderOptions{}
	op.Images[0] = e.fillImageToScreen(current)
	op.Images[1] = e.fillImageToScreen(next)
	op.Uniforms = map[string]any{
		"Progress":   progress,
		"Resolution": [2]float32{float32(e.width), float32(e.height)},
		"Bubbles":    bubbleData,
	}

	screen.DrawRectShader(e.width, e.height, e.shader, op)
}

func (e *BubbleMeltShaderEffect) fillImageToScreen(img *ebiten.Image) *ebiten.Image {
	// Create a new black image with screen dimensions
	screenImg := ebiten.NewImage(e.width, e.height)
	screenImg.Fill(color.Black)

	// Get the dimensions of the input image
	imgWidth, imgHeight := img.Bounds().Dx(), img.Bounds().Dy()

	// Calculate scaling factors
	scaleX := float64(e.width) / float64(imgWidth)
	scaleY := float64(e.height) / float64(imgHeight)
	scale := math.Min(scaleX, scaleY)

	// Calculate new dimensions
	newWidth := int(float64(imgWidth) * scale)
	newHeight := int(float64(imgHeight) * scale)

	// Calculate position to center the image
	x := (e.width - newWidth) / 2
	y := (e.height - newHeight) / 2

	// Create draw options
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(x), float64(y))

	// Draw the scaled and centered image onto the black background
	screenImg.DrawImage(img, op)

	return screenImg
}
