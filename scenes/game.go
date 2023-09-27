package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"juego/models"
)

type Game struct{}

func (g *Game) Update() error {
	// No necesitamos lógica de movimiento aquí, se realiza en goroutines
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Dibujo del juego
	screen.DrawImage(backgroundImage, nil)
	drawImage(screen, paddleImage, 0, models.Player1Y)
	drawImage(screen, paddleImage, 520, models.Player2Y)
	drawImage(screen, ballImage, models.BallX, models.BallY)

	drawImage(screen, def1, models.ScreenWidth/4-models.PaddleWidth/2, models.Paddle1Y)
	drawImage(screen, def2, (3*models.ScreenWidth)/4-models.PaddleWidth/2-20, models.Paddle2Y)

	// Dibuja el marcador en la parte superior izquierda
	ebitenutil.DebugPrint(screen, "Player 1")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return models.ScreenWidth, models.ScreenHeight
}

func DrawImage(screen *ebiten.Image, img *ebiten.Image, x, y int) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, opts)
}