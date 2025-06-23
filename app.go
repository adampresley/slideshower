package main

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ErrUserQuit = errors.New("user quit")
)

type App struct {
	ScreenWidth  int
	ScreenHeight int
	Slideshow    *Slideshow
	shouldQuit   bool
}

func (a *App) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		a.shouldQuit = true
		return ErrUserQuit
	}

	return a.Slideshow.Update()
}

func (a *App) Draw(screen *ebiten.Image) {
	a.Slideshow.Draw(screen)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return a.ScreenWidth, a.ScreenHeight
}
