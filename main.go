package main

import (
	"log"
	"time"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"juego/models"
	"juego/scenes"
)

var (
	mutex      sync.Mutex
	lastUpdate time.Time
	paddleImage     *ebiten.Image
	ballImage       *ebiten.Image
	backgroundImage *ebiten.Image
	def1            *ebiten.Image
	def2            *ebiten.Image
)

func main() {
	lastUpdate = time.Now()
	ebiten.SetWindowSize(models.ScreenWidth, models.ScreenHeight)
	ebiten.SetWindowTitle("Ping Pong Game")

	// Carga las imágenes desde archivos
	var err error
	paddleImage, _, err = ebitenutil.NewImageFromFile("assets/portero.png")
	if err != nil {
		log.Fatal(err)
	}

	ballImage, _, err = ebitenutil.NewImageFromFile("assets/balon.png")
	if err != nil {
		log.Fatal(err)
	}

	backgroundImage, _, err = ebitenutil.NewImageFromFile("assets/fondo.png")
	if err != nil {
		log.Fatal(err)
	}

	def1, _, err = ebitenutil.NewImageFromFile("assets/defensa.png")
	if err != nil {
		log.Fatal(err)
	}

	def2, _, err = ebitenutil.NewImageFromFile("assets/defensa.png")
	if err != nil {
		log.Fatal(err)
	}

	// Inicia goroutines para la lógica de movimiento
	go movePlayer1()
	go movePlayer2()
	go moveBall()
	go moveStaticPaddles()

	game := &scenes.Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
