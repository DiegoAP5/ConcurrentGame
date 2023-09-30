package models

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	PaddleImage     *ebiten.Image
	BallImage       *ebiten.Image
	BackgroundImage *ebiten.Image
	Def1            *ebiten.Image
	Def2            *ebiten.Image
)

func DrawImage(screen, img *ebiten.Image, x, y int) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, opts)
}
