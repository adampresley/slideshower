package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Slideshow struct {
	imagePaths     []string
	imageIndex     int
	screenWidth    float64
	screenHeight   float64
	speedInSeconds int

	currentImage *ebiten.Image
	nextImage    *ebiten.Image

	preloadTimer    *Timer
	nextTimer       *Timer
	effect          Effect
	isTransitioning bool
}

type SlideshowConfig struct {
	ImagePaths     []string
	ScreenWidth    float64
	ScreenHeight   float64
	SpeedInSeconds int
	Effect         Effect
}

func NewSlideshow(config SlideshowConfig) *Slideshow {
	return &Slideshow{
		imagePaths:     config.ImagePaths,
		screenWidth:    config.ScreenWidth,
		screenHeight:   config.ScreenHeight,
		speedInSeconds: config.SpeedInSeconds,

		currentImage: mustLoadImage(config.ImagePaths[0]),

		preloadTimer: NewTimer(time.Second * time.Duration(config.SpeedInSeconds-1)),
		nextTimer:    NewTimer(time.Second * time.Duration(config.SpeedInSeconds)),

		effect: config.Effect,
	}
}

func (s *Slideshow) Update() error {
	s.preloadTimer.Update()
	s.nextTimer.Update()

	if s.isTransitioning {
		s.effect.Update()

		if s.effect.IsComplete() {
			s.currentImage = s.nextImage
			s.isTransitioning = false
			s.effect.Reset()
		}
	} else if s.nextTimer.IsReady() {
		if s.nextImage != nil {
			s.isTransitioning = true
		}

		s.nextTimer.Reset()
	}

	if s.preloadTimer.IsReady() {
		s.imageIndex++

		if s.imageIndex >= len(s.imagePaths) {
			s.imageIndex = 0
		}

		s.nextImage = mustLoadImage(s.imagePaths[s.imageIndex])
		s.preloadTimer.Reset()
	}

	return nil
}

func (s *Slideshow) Draw(screen *ebiten.Image) {
	if s.isTransitioning && s.nextImage != nil {
		s.effect.Draw(screen, s.currentImage, s.nextImage)
	} else {
		s.drawImage(screen, s.currentImage)
	}
}

func (s *Slideshow) drawImage(screen *ebiten.Image, img *ebiten.Image) {
	// Get the image dimensions
	b := img.Bounds()
	imageWidth := float64(b.Dx())
	imageHeight := float64(b.Dy())

	// Calculate scaling factor while maintaining aspect ratio
	scaleX := s.screenWidth / imageWidth
	scaleY := s.screenHeight / imageHeight
	scale := math.Min(scaleX, scaleY)

	// Calculate new dimensions
	newWidth := imageWidth * scale
	newHeight := imageHeight * scale

	// Calculate position to center the image
	x := (s.screenWidth - newWidth) / 2
	y := (s.screenHeight - newHeight) / 2

	// Draw the current image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	screen.DrawImage(img, op)
}
