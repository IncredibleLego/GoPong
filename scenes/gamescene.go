package scenes

import (
	"goPong/config"
	"goPong/objects"
	"goPong/utils"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type GameScene struct {
	playerName string
	paddle     *objects.Paddle
	ball       *objects.Ball
	score      int
	increase   int
	bestScore  int
	showRecord bool
	recordTime time.Time
}

func (g *GameScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return reason == Unpause
}

func NewGameScene() *GameScene {
	return &GameScene{
		playerName: "",
		paddle:     nil,
		ball:       nil,
		score:      0,
		increase:   0,
	}
}

func (g *GameScene) Draw(screen *ebiten.Image) {

	// Draw the paddle
	g.paddle.Draw(screen)

	// Draw the ball
	g.ball.Draw(screen)

	// Draw the net
	utils.Net(screen)

	// Draw the points
	utils.PointsDraw(screen, float32(config.GlobalConfig.ScreenWidth)/6+float32(70*config.DefaultConfig.Scale)/2, float32(config.GlobalConfig.ScreenHeight)/14, g.score)

	X := float64(config.GlobalConfig.ScreenWidth/4*3) - ((config.GlobalConfig.TextDimension - 3) * float64(len(g.playerName)/2))

	// Player name
	utils.ScreenDraw(-3, X, float64(config.GlobalConfig.ScreenHeight)/72, "white", screen, g.playerName)

	// Draw Wall on left
	vector.DrawFilledRect(screen,
		0, 0,
		float32(config.GlobalConfig.PaddleWidth), float32(config.GlobalConfig.ScreenHeight),
		color.White, false,
	)

	// Print message once if a new record is set
	if g.showRecord && time.Now().Before(g.recordTime) {
		utils.NewHighscore(screen)
	}
}

// FirstLoad implements Scene.
func (g *GameScene) FirstLoad() {
	g.playerName = config.GlobalConfig.Player1Name
	g.paddle = &objects.Paddle{
		Object: &objects.Object{
			X: config.GlobalConfig.ScreenWidth - config.GlobalConfig.PaddleDistanceFromWall,
			Y: config.GlobalConfig.ScreenHeight/2 - config.GlobalConfig.PaddleHeight/2,
			W: config.GlobalConfig.PaddleWidth,
			H: config.GlobalConfig.PaddleHeight,
		},
	}
	g.ball = &objects.Ball{
		Object: &objects.Object{
			X: config.GlobalConfig.ScreenWidth / 2,
			Y: config.GlobalConfig.ScreenHeight / 2,
			W: config.GlobalConfig.BallSize,
			H: config.GlobalConfig.BallSize,
		},
		Dxdt: config.GlobalConfig.BallSpeed,
		Dydt: config.GlobalConfig.BallSpeed,
	}
	g.ball.GenerateRandomDirection()
	g.score = 0
	hs, _ := loadHighscores()
	if len(hs.Solo) == 0 {
		g.bestScore = 0
	} else {
		g.bestScore = hs.Solo[0].Score
	}
	g.ball.Reset(false)
}

func (g *GameScene) OnEnter() {}

func (g *GameScene) OnExit() {}

func (g *GameScene) updateDimensions() {
	g.ball.W = config.GlobalConfig.BallSize
	g.ball.H = config.GlobalConfig.BallSize
	g.paddle.H = config.GlobalConfig.PaddleHeight
}

func (g *GameScene) Update() SceneId {

	g.updateDimensions()

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		highScore, maxAdded := getTopSoloScore()
		if maxAdded == false || g.score >= highScore.Score {
			DirtySoloScore = SoloScore{
				DateTime: time.Now().Format(time.RFC3339),
				Player:   g.playerName,
				Score:    g.score,
			}
		}
		return PauseSceneId
	}

	if !g.paddle.MoveOnKeyPress(ebiten.KeyArrowUp, ebiten.KeyArrowDown) {
		g.paddle.MoveOnKeyPress(ebiten.KeyW, ebiten.KeyS)
	}

	g.ball.Move()

	if g.ball.X+g.ball.W >= config.GlobalConfig.ScreenWidth {
		g.score = 0
		g.increase = 0
	}
	g.ball.CollideWithWall(false, true, config.GlobalConfig.PaddleWidth)

	if g.ball.CollideWithPaddle(g.paddle, true, g.increase) {
		g.score++
		if g.score > g.bestScore && g.showRecord == false {
			g.showRecord = true
			g.recordTime = time.Now().Add(3 * time.Second)
		}
		if g.score%5 == 0 {
			g.increase += 2
		}
	}

	return GameSceneId
}

var _ Scene = (*GameScene)(nil)
