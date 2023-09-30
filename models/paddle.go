package models

import "time"

var Paddle1Direction = 1
var Paddle2Direction = -1
var Paddle1Y = ScreenHeight / 4
var Paddle2Y = (3 * ScreenHeight) / 4

// ... (resto del archivo sin cambios) ...

func MoveStaticPaddles() {
	speed := 5  

	for {
		Paddle1Y += Paddle1Direction * speed
		if Paddle1Y < 0 {
			Paddle1Y = 0
			Paddle1Direction = 1 
		}
		if Paddle1Y > ScreenHeight-paddleHeight {
			Paddle1Y = ScreenHeight - paddleHeight
			Paddle1Direction = -1 
		}

		Paddle2Y += Paddle2Direction * speed
		if Paddle2Y < 0 {
			Paddle2Y = 0
			Paddle2Direction = 1  
		}
		if Paddle2Y > ScreenHeight-paddleHeight {
			Paddle2Y = ScreenHeight - paddleHeight
			Paddle2Direction = -1  
		}

		time.Sleep(time.Millisecond * 50)
	}
}

