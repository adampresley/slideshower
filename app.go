package main

import "github.com/hajimehoshi/ebiten/v2"

type App struct {
	ScreenWidth  int
	ScreenHeight int
	Slideshow    *Slideshow
}

func (a *App) Update() error {
	_ = a.Slideshow.Update()
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	a.Slideshow.Draw(screen)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return a.ScreenWidth, a.ScreenHeight
}
