package main

import (
	"log"
	"sync"
	"time"
	"github.com/hajimehoshi/ebiten/v2"
	"juego/scenes"
	"juego/models"
)

var lastUpdate time.Time
var mutex sync.Mutex

const (
	ScreenHeight = 480
	ScreenWidth  = 640
)

func main() {
	lastUpdate = time.Now()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Futbolito")
	
	scenes.Images()

	go models.MovePlayer1()
	go models.MovePlayer2()
	go models.MoveBall()
	go models.MoveStaticPaddles() 

	game := &scenes.Gameplay{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
