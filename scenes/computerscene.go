package scenes

import (
	"goPong/config"
	"goPong/objects"
	"goPong/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	return reason == Unpause
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

	// Draw the paddle
	c.paddle.Draw(screen)
	// Draw enemy paddle
	c.enemyPaddle.Draw(screen)

	// Draw the ball
	c.ball.Draw(screen)

	// Draw the net
	utils.Net(screen)

	measure := float32(70 * config.DefaultConfig.Scale)

	// Draw the points
	utils.PointsDraw(screen, float32(config.GlobalConfig.ScreenWidth)/6+measure/2, float32(config.GlobalConfig.ScreenHeight)/14, c.scoreEnemy)
	utils.PointsDraw(screen, (float32(config.GlobalConfig.ScreenWidth)/6)*4+measure/2, float32(config.GlobalConfig.ScreenHeight)/14, c.score)

	X1 := float64(config.GlobalConfig.ScreenWidth/4) - ((config.GlobalConfig.TextDimension - 3) * float64(len("COMPUTER")/2))
	X2 := float64(config.GlobalConfig.ScreenWidth/4*3) - ((config.GlobalConfig.TextDimension - 3) * float64(len(c.playerName)/2))

	utils.ScreenDraw(-3, X1, float64(config.GlobalConfig.ScreenHeight)/72, "white", screen, "COMPUTER")
	utils.ScreenDraw(-3, X2, float64(config.GlobalConfig.ScreenHeight)/72, "white", screen, c.playerName)

}

// FirstLoad implements Scene.
func (c *ComputerScene) FirstLoad() {
	c.playerName = config.GlobalConfig.Player1Name
	c.paddle = &objects.Paddle{
		Object: &objects.Object{
			X: config.GlobalConfig.ScreenWidth - config.GlobalConfig.PaddleDistanceFromWall,
			Y: config.GlobalConfig.ScreenHeight/2 - config.GlobalConfig.PaddleHeight/2,
			W: config.GlobalConfig.PaddleWidth,
			H: config.GlobalConfig.PaddleHeight,
		},
	}
	c.enemyPaddle = &objects.Paddle{
		Object: &objects.Object{
			X: config.GlobalConfig.PaddleDistanceFromWall,
			Y: config.GlobalConfig.ScreenHeight/2 - config.GlobalConfig.PaddleHeight/2,
			W: config.GlobalConfig.PaddleWidth,
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
	c.ball.Reset(false)
}

func (c *ComputerScene) OnEnter() {

}

func (c *ComputerScene) OnExit() {

}

func (c *ComputerScene) updateDimensions() {
	c.ball.W = config.GlobalConfig.BallSize
	c.ball.H = config.GlobalConfig.BallSize
	c.paddle.H = config.GlobalConfig.PaddleHeight
	c.enemyPaddle.H = config.GlobalConfig.PaddleHeight
}

func (c *ComputerScene) Update() SceneId {

	c.updateDimensions()

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return PauseSceneId
	}

	c.paddle.MoveOnKeyPress(ebiten.KeyArrowUp, ebiten.KeyArrowDown)
	//c.enemyPaddle.MoveOnKeyPress(ebiten.KeyW, ebiten.KeyS)
	c.enemyPaddle.AiMovement(c.ball)

	c.ball.Move()

	test := c.ball.CollideWithWall(true, true)
	if test == 1 {
		c.score++
	} else if test == 2 {
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
