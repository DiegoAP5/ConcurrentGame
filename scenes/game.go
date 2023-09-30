package scenes

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"juego/models"
)

const (
	screenWidth  = 640
	screenHeight = 480
	paddleWidth  = 15
)

type Gameplay struct{}

func (g *Gameplay) Update() error {
	return nil
}

func (g *Gameplay) Draw(screen *ebiten.Image) {
	screen.DrawImage(models.BackgroundImage, nil)
	
	models.DrawImage(screen, models.PaddleImage, 0, models.Player1Y)
	models.DrawImage(screen, models.PaddleImage, 520, models.Player2Y)
	models.DrawImage(screen, models.BallImage, models.BallX, models.BallY)
	models.DrawImage(screen, models.Def1, screenWidth/4-paddleWidth/2, models.Paddle1Y)
	models.DrawImage(screen, models.Def2, (3*screenWidth)/4-paddleWidth/2-20, models.Paddle2Y)
	
	ebitenutil.DebugPrint(screen, "Player 1")

	
}

func Images(){
	
	var err error
	models.PaddleImage, _, err = ebitenutil.NewImageFromFile("assets/portero.png")
	if err != nil {
		log.Fatal(err)
	}

	models.BallImage, _, err = ebitenutil.NewImageFromFile("assets/balon.png")
	if err != nil {
		log.Fatal(err)
	}

	models.BackgroundImage, _, err = ebitenutil.NewImageFromFile("assets/fondo.png")
	if err != nil {
		log.Fatal(err)
	}

	models.Def1, _, err = ebitenutil.NewImageFromFile("assets/defensa.png")
	if err != nil {
		log.Fatal(err)
	}

	models.Def2, _, err = ebitenutil.NewImageFromFile("assets/defensa.png")
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Gameplay) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
