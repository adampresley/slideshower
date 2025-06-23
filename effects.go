package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Effect is the interface for transition effects
type Effect interface {
	Update()
	Draw(screen *ebiten.Image, current, next *ebiten.Image)
	IsComplete() bool
	Reset()
}
