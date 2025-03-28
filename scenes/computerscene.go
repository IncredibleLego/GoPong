package scenes

import (
	"goPong/config"
	"goPong/menu"
	"goPong/objects"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ComputerScene struct {
	playerName  string
	paddle      *objects.Paddle
	enemyPaddle *objects.Paddle
	ball        *objects.Ball
	score       int
	scoreEnemy  int
	highScore   int
}

func (c *ComputerScene) ShouldPreserveState(reason SceneChangeReason) bool {
	if reason == Unpause {
		return true
	}
	return false
}

func NewComputerScene() *ComputerScene {
	return &ComputerScene{
		playerName:  "",
		paddle:      nil,
		enemyPaddle: nil,
		ball:        nil,
		score:       0,
		scoreEnemy:  0,
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

	for i := 0; i < config.GlobalConfig.ScreenHeight; i += 24 {
		vector.DrawFilledRect(screen,
			float32(config.GlobalConfig.ScreenWidth/2), float32(i),
			float32(3), float32(12),
			color.White, false,
		)
	}

	menu.ScreenDraw(0, 10, 10, "white", screen, "COMPUTER")
	menu.ScreenDraw(0, 10, 25, "white", screen, "Score: "+strconv.Itoa(c.scoreEnemy))
	menu.ScreenDraw(0, 500, 10, "white", screen, c.playerName)
	menu.ScreenDraw(0, 500, 25, "white", screen, "Score: "+strconv.Itoa(c.score))
	menu.ScreenDraw(-3, 250, 10, "white", screen, "COMPUTER MODE")
}

// FirstLoad implements Scene.
func (c *ComputerScene) FirstLoad() {
	c.playerName = config.GlobalConfig.Player1Name
	c.paddle = &objects.Paddle{
		Object: &objects.Object{
			X: config.GlobalConfig.ScreenWidth - config.GlobalConfig.PaddleDistanceFromWall,
			Y: config.GlobalConfig.ScreenHeight/2 - config.GlobalConfig.PaddleHeight/2,
			W: 15,
			H: config.GlobalConfig.PaddleHeight,
		},
	}
	c.enemyPaddle = &objects.Paddle{
		Object: &objects.Object{
			X: config.GlobalConfig.PaddleDistanceFromWall,
			Y: config.GlobalConfig.ScreenHeight/2 - config.GlobalConfig.PaddleHeight/2,
			W: 15,
			H: config.GlobalConfig.PaddleHeight,
		},
	}
	c.ball = &objects.Ball{
		Object: &objects.Object{
			X: config.GlobalConfig.ScreenWidth / 2,
			Y: config.GlobalConfig.ScreenHeight / 2,
			W: config.GlobalConfig.BallSize,
			H: config.GlobalConfig.BallSize,
		},
		Dxdt: config.GlobalConfig.BallSpeed,
		Dydt: config.GlobalConfig.BallSpeed,
	}
	c.ball.GenerateRandomDirection()
	c.score = 0
	c.highScore = 0
}

func (c *ComputerScene) OnEnter() {

}

func (c *ComputerScene) OnExit() {

}

func (c *ComputerScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return PauseSceneId
	}

	c.paddle.MoveOnKeyPress(ebiten.KeyArrowUp, ebiten.KeyArrowDown)
	//c.enemyPaddle.MoveOnKeyPress(ebiten.KeyW, ebiten.KeyS)
	c.enemyPaddle.AiMovement(float64(c.ball.Y))

	c.ball.Move()
	c.ball.CollideWithWall(true, true)

	if c.ball.CollideWithWall(true, true) == 1 {
		c.score++
	} else if c.ball.CollideWithWall(true, true) == 2 {
		c.scoreEnemy++
	}

	c.ball.CollideWithPaddle(c.paddle, true)
	c.ball.CollideWithPaddle(c.enemyPaddle, false)

	return ComputerSceneId
}

var _ Scene = (*ComputerScene)(nil)

func (c *ComputerScene) CollideWithWall() { // Check if the ball collides with the wall
	if c.ball.X >= config.GlobalConfig.ScreenWidth {
		c.Reset()
	} else if c.ball.X <= 0 {
		c.ball.Dxdt = config.GlobalConfig.BallSpeed
	} else if c.ball.Y <= 0 {
		c.ball.Dydt = config.GlobalConfig.BallSpeed
	} else if c.ball.Y >= config.GlobalConfig.ScreenHeight {
		c.ball.Dydt = -config.GlobalConfig.BallSpeed
	}
}

func (c *ComputerScene) IncreaseScore() {
	c.score++
	if c.score > c.highScore {
		c.highScore = c.score
	}
}

func (c *ComputerScene) Reset() { // Reset the game
	c.ball.X = config.GlobalConfig.ScreenWidth / 2
	c.ball.Y = config.GlobalConfig.ScreenHeight / 2
	c.score = 0
}
