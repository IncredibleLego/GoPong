package scenes

import (
	"goPong/config"
	"goPong/objects"
	"goPong/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ComputerScene struct {
	playerName  string
	enemyName   string
	paddle      *objects.Paddle
	enemyPaddle *objects.Paddle
	ball        *objects.Ball
	score       int
	scoreEnemy  int
	bestScore   int
	showRecord  bool
	recordTime  time.Time
}

func (c *ComputerScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return reason == Unpause
}

func NewComputerScene() *ComputerScene {
	return &ComputerScene{
		playerName:  "",
		enemyName:   "COMPUTER",
		paddle:      nil,
		enemyPaddle: nil,
		ball:        nil,
		score:       0,
		scoreEnemy:  0,
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

	X1 := float64(config.GlobalConfig.ScreenWidth/4) - ((config.GlobalConfig.TextDimension - 3) * float64(len(c.enemyName)/2))
	X2 := float64(config.GlobalConfig.ScreenWidth/4*3) - ((config.GlobalConfig.TextDimension - 3) * float64(len(c.playerName)/2))

	utils.ScreenDraw(-3, X1, float64(config.GlobalConfig.ScreenHeight)/72, "white", screen, c.enemyName)
	utils.ScreenDraw(-3, X2, float64(config.GlobalConfig.ScreenHeight)/72, "white", screen, c.playerName)

	// Print message once if new record is set
	if c.showRecord && time.Now().Before(c.recordTime) {
		utils.NewHighscore(screen)
	}
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
	c.enemyPaddle.InitAIStateFromCurrentY()

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
	c.scoreEnemy = 0
	hs, _ := loadHighscores()
	if config.GlobalConfig.Difficulty < 0.33 {
		if len(hs.ComputerEasy) == 0 {
			c.bestScore = 0
		} else {
			c.bestScore = hs.ComputerEasy[0].Score
		}
	} else if config.GlobalConfig.Difficulty < 0.66 {
		if len(hs.ComputerDefault) == 0 {
			c.bestScore = 0
		} else {
			c.bestScore = hs.ComputerDefault[0].Score
		}
	} else {
		if len(hs.ComputerHard) == 0 {
			c.bestScore = 0
		} else {
			c.bestScore = hs.ComputerHard[0].Score
		}
	}
	c.ball.Reset(false)
}

func (c *ComputerScene) OnEnter() {}

func (c *ComputerScene) updateDimensions() {
	c.ball.W = config.GlobalConfig.BallSize
	c.ball.H = config.GlobalConfig.BallSize
	c.paddle.H = config.GlobalConfig.PaddleHeight
	c.enemyPaddle.H = config.GlobalConfig.PaddleHeight
}

func (c *ComputerScene) Update() SceneId {

	c.updateDimensions()

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		highScore, maxAdded := getTopComputerScore()
		if maxAdded == false || c.score >= highScore.Score {
			DirtyComputerScore = ComputerScore{
				DateTime: time.Now().Format(time.RFC3339),
				Player:   c.playerName,
				AILevel:  config.DifficultyString(),
				Score:    c.score,
			}
		}
		return PauseSceneId
	}

	c.paddle.MoveOnKeyPress(ebiten.KeyArrowUp, ebiten.KeyArrowDown)
	c.enemyPaddle.AiMovement(c.ball)
	c.ball.Move()

	test := c.ball.CollideWithWall(true, true, 0)
	if test == 1 {
		c.score++
	} else if test == 2 {
		c.scoreEnemy++
	}

	if c.score > c.bestScore && c.showRecord == false {
		c.showRecord = true
		c.recordTime = time.Now().Add(3 * time.Second)
	}

	c.ball.CollideWithPaddle(c.paddle, true, 0)
	c.ball.CollideWithPaddle(c.enemyPaddle, false, 0)

	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		c.enemyName = utils.Kubrick(c.enemyName)
	}

	return ComputerSceneId
}

var _ Scene = (*ComputerScene)(nil)
