package models

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const (
	ScreenHeight = 480
	paddleHeight = 80
	ScreenWidth  = 640
	paddleWidth  = 15
)

var Player1Y = ScreenHeight / 2
var Player2Y = ScreenHeight / 2

func MovePlayer1() {
	for {
		if ebiten.IsKeyPressed(ebiten.KeyW) && Player1Y > 0 {
			Player1Y -= 5
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) && Player1Y < ScreenHeight-paddleHeight {
			Player1Y += 5
		}
		time.Sleep(time.Millisecond * 16)
	}
}

func MovePlayer2() {
	for {
		if ebiten.IsKeyPressed(ebiten.KeyUp) && Player2Y > 0 {
			Player2Y -= 5
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) && Player2Y < ScreenHeight-paddleHeight {
			Player2Y += 5
		}
		time.Sleep(time.Millisecond * 16)
	}
}
