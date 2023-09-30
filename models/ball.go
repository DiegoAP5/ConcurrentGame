package models

import (
	"time"
)

var (
	BallX, BallY = ScreenWidth / 2, ScreenHeight / 2
	BallDX = 2
	BallDY = 2
	ballSize     = 10
)

// ... (resto del archivo sin cambios) ...

func MoveBall() {
	for {
		BallX += BallDX 
		BallY += BallDY 

		if BallY < 0 || BallY > ScreenHeight-ballSize {
			BallDY *= -1
		}

		if BallX < paddleWidth && BallY > Player1Y && BallY < Player1Y+paddleHeight {
			BallDX *= -1
		}

		if BallX > ScreenWidth-paddleWidth && BallY > Player2Y && BallY < Player2Y+paddleHeight {
			if BallX < ScreenWidth-paddleWidth/2 {
				BallDX *= -1
			}
		}

		if BallX < 0 || BallX > ScreenWidth {
			BallX, BallY = ScreenWidth/2, ScreenHeight/2
			BallDX *= -1
		}

		if BallDX < 0 && BallX < ScreenWidth/4+paddleWidth/2 && BallY > Paddle1Y && BallY < Paddle1Y+paddleHeight {
			BallDX *= -1 
		}

		if BallDX > 0 && BallX > (3*ScreenWidth)/4-paddleWidth/2 && BallY > Paddle2Y && BallY < Paddle2Y+paddleHeight {
			BallDX *= -1
		}

		time.Sleep(time.Millisecond * 1) 
	}
}

