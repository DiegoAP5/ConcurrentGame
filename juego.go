package main

import (
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	paddleWidth  = 15
	paddleHeight = 80
	ballSize     = 10
)

var (
	player1Y     = screenHeight / 2
	player2Y     = screenHeight / 2
	ballX, ballY = screenWidth / 2, screenHeight / 2
	ballDX       = 2
	ballDY       = 2

	paddle1Direction = 1  
    paddle2Direction = -1

	paddle1Y = screenHeight / 4
	paddle2Y = (3 * screenHeight) / 4

	mutex      sync.Mutex
	lastUpdate time.Time

	paddleImage     *ebiten.Image
	ballImage       *ebiten.Image
	backgroundImage *ebiten.Image
	def1            *ebiten.Image
	def2            *ebiten.Image
)

type Game struct{}

func (g *Game) Update() error {
	// No necesitamos lógica de movimiento aquí, se realiza en goroutines
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Dibuja el fondo
	screen.DrawImage(backgroundImage, nil)

	// Dibuja las paletas y la pelota utilizando las imágenes
	drawImage(screen, paddleImage, 0, player1Y)
	drawImage(screen, paddleImage, 520, player2Y)
	drawImage(screen, ballImage, ballX, ballY)

	drawImage(screen, def1, screenWidth/4-paddleWidth/2, paddle1Y)
	drawImage(screen, def2, (3*screenWidth)/4-paddleWidth/2-20, paddle2Y)

	// Dibuja el marcador en la parte superior izquierda
	ebitenutil.DebugPrint(screen, "Player 1")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func drawImage(screen *ebiten.Image, img *ebiten.Image, x, y int) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, opts)
}

func main() {
	lastUpdate = time.Now()
	ebiten.SetWindowSize(screenWidth, screenHeight)
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

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func movePlayer1() {
	for {
		// Lógica de movimiento del jugador 1 (paleta izquierda)
		if ebiten.IsKeyPressed(ebiten.KeyW) && player1Y > 0 {
			player1Y -= 5 // Ajusta la velocidad de movimiento
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) && player1Y < screenHeight-paddleHeight {
			player1Y += 5 // Ajusta la velocidad de movimiento
		}
		time.Sleep(time.Millisecond * 16) // Ajusta el tiempo según la velocidad deseada
	}
}

func movePlayer2() {
	for {
		// Lógica de movimiento del jugador 2 (paleta derecha)
		if ebiten.IsKeyPressed(ebiten.KeyUp) && player2Y > 0 {
			player2Y -= 5 // Ajusta la velocidad de movimiento
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) && player2Y < screenHeight-paddleHeight {
			player2Y += 5 // Ajusta la velocidad de movimiento
		}
		time.Sleep(time.Millisecond * 16) // Ajusta el tiempo según la velocidad deseada
	}
}

func moveBall() {
	for {
		// Lógica de movimiento de la pelota
		ballX += ballDX // Ajusta la velocidad de movimiento
		ballY += ballDY // Ajusta la velocidad de movimiento

		if ballY < 0 || ballY > screenHeight-ballSize {
			ballDY *= -1
		}

		// Colisión con paleta del jugador 1 (paleta izquierda)
		if ballX < paddleWidth && ballY > player1Y && ballY < player1Y+paddleHeight {
			ballDX *= -1
		}

		// Colisión con paleta del jugador 2 (paleta derecha)
		if ballX > screenWidth-paddleWidth && ballY > player2Y && ballY < player2Y+paddleHeight {
			if ballX < screenWidth-paddleWidth/2 {
				ballDX *= -1
			}
		}

		if ballX < 0 || ballX > screenWidth {
			ballX, ballY = screenWidth/2, screenHeight/2
			ballDX *= -1
		}

		// Colisión con la primera paleta estática
		if ballDX < 0 && ballX < screenWidth/4+paddleWidth/2 && ballY > paddle1Y && ballY < paddle1Y+paddleHeight {
			ballDX *= -1 // Cambia la dirección horizontal de la pelota
		}

		// Colisión con la segunda paleta estática
		if ballDX > 0 && ballX > (3*screenWidth)/4-paddleWidth/2 && ballY > paddle2Y && ballY < paddle2Y+paddleHeight {
			ballDX *= -1 // Cambia la dirección horizontal de la pelota
		}

		time.Sleep(time.Millisecond * 1) // Ajusta el tiempo según la velocidad deseada
	}
}

func moveStaticPaddles() {
	speed := 5  // Ajusta según la velocidad que desees

	for {
		// Mover paddle1Y
		paddle1Y += paddle1Direction * speed
		if paddle1Y < 0 {
			paddle1Y = 0
			paddle1Direction = 1  // Cambia la dirección a hacia abajo
		}
		if paddle1Y > screenHeight-paddleHeight {
			paddle1Y = screenHeight - paddleHeight
			paddle1Direction = -1  // Cambia la dirección a hacia arriba
		}

		// Mover paddle2Y
		paddle2Y += paddle2Direction * speed
		if paddle2Y < 0 {
			paddle2Y = 0
			paddle2Direction = 1  // Cambia la dirección a hacia abajo
		}
		if paddle2Y > screenHeight-paddleHeight {
			paddle2Y = screenHeight - paddleHeight
			paddle2Direction = -1  // Cambia la dirección a hacia arriba
		}

		time.Sleep(time.Millisecond * 50) // Ajusta este valor para controlar cuán rápido se actualizan las posiciones
	}
}
