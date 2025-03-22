package scenes

import (
	"goPong/constants"
	"goPong/menu"
	"goPong/objects"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ComputerScene struct {
	loaded      bool
	paddle      *objects.Paddle
	enemyPaddle *objects.Paddle
	ball        *objects.Ball
	score       int
	highScore   int
}

func (c *ComputerScene) IsLoaded() bool {
	return c.loaded
}

func NewComputerScene() *ComputerScene {
	return &ComputerScene{
		loaded:      false,
		paddle:      nil,
		enemyPaddle: nil,
		ball:        nil,
		score:       0,
		highScore:   0,
	}
}

func (c *ComputerScene) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		float32(c.paddle.X), float32(c.paddle.Y),
		float32(c.paddle.W), float32(c.paddle.H),
		color.White, false,
	) // Draw the paddle
	vector.DrawFilledRect(screen,
		float32(c.enemyPaddle.X), float32(c.enemyPaddle.Y),
		float32(c.enemyPaddle.W), float32(c.enemyPaddle.H),
		color.White, false,
	) // Draw enemy paddle
	vector.DrawFilledRect(screen,
		float32(c.ball.X), float32(c.ball.Y),
		float32(c.ball.W), float32(c.ball.H),
		color.White, false,
	) // Draw the ball

	// Draw center lines

	for i := 0; i < constants.ScreenHeight; i += 24 {
		vector.DrawFilledRect(screen,
			float32(constants.ScreenWidth/2), float32(i),
			float32(3), float32(12),
			color.White, false,
		)
	}

	menu.ScreenDraw(constants.TextDimension, 10, 10, 1, 1, 1, 1, 1.5, screen, "Score: "+strconv.Itoa(c.score))
	menu.ScreenDraw(constants.TextDimension, 10, 30, 1, 1, 1, 1, 1.5, screen, "High Score: "+strconv.Itoa(c.highScore))
	menu.ScreenDraw(constants.TextDimension, 450, 10, 1, 1, 1, 1, 1.5, screen, "COMPUTER MODE")
}

// FirstLoad implements Scene.
func (c *ComputerScene) FirstLoad() {
	c.paddle = &objects.Paddle{
		Object: &objects.Object{
			X: constants.ScreenWidth - constants.PaddleDistanceFromWall,
			Y: constants.ScreenHeight / 2,
			W: 15,
			H: 100,
		},
	}
	c.enemyPaddle = &objects.Paddle{
		Object: &objects.Object{
			X: constants.PaddleDistanceFromWall,
			Y: constants.ScreenHeight / 2,
			W: 15,
			H: 100,
		},
	}
	c.ball = &objects.Ball{
		Object: &objects.Object{
			X: constants.ScreenWidth / 2,
			Y: constants.ScreenHeight / 2,
			W: 15,
			H: 15,
		},
		Dxdt: constants.BallSpeed,
		Dydt: constants.BallSpeed,
	}
}

func (c *ComputerScene) OnEnter() {

}

func (c *ComputerScene) OnExit() {

}

func (c *ComputerScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return PauseSceneId
	}

	if !c.paddle.MoveOnKeyPress(ebiten.KeyArrowUp, ebiten.KeyArrowDown) {
		c.paddle.MoveOnKeyPress(ebiten.KeyW, ebiten.KeyS)
	}

	c.ball.Move()
	c.CollideWithWall()

	if c.paddle.CollideWithPaddle(c.ball, true) {
		c.IncreaseScore()
		if c.score%5 == 0 {
			c.ball.IncreaseSpeed(2)
		}
		c.ball.Dxdt = -c.ball.Dxdt
	}

	return ComputerSceneId
}

var _ Scene = (*ComputerScene)(nil)

func (c *ComputerScene) CollideWithWall() { // Check if the ball collides with the wall
	if c.ball.X >= constants.ScreenWidth {
		c.Reset()
	} else if c.ball.X <= 0 {
		c.ball.Dxdt = constants.BallSpeed
	} else if c.ball.Y <= 0 {
		c.ball.Dydt = constants.BallSpeed
	} else if c.ball.Y >= constants.ScreenHeight {
		c.ball.Dydt = -constants.BallSpeed
	}
}

func (c *ComputerScene) IncreaseScore() {
	c.score++
	if c.score > c.highScore {
		c.highScore = c.score
	}
}

func (c *ComputerScene) Reset() { // Reset the game
	c.ball.X = constants.ScreenWidth / 2
	c.ball.Y = constants.ScreenHeight / 2
	c.score = 0
}
