package scenes

import (
	"goPong/config"
	"goPong/objects"
	"goPong/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameScene struct {
	playerName string
	paddle     *objects.Paddle
	ball       *objects.Ball
	score      int
	highScore  int
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
		highScore:  0,
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
	utils.PointsDraw(screen, 100, 50, g.score)

	utils.ScreenDraw(0, 400, 10, "white", screen, g.playerName)

	/*
		utils.ScreenDraw(0, 10, 10, "white", screen, "Score "+g.playerName+":"+strconv.Itoa(g.score))
		utils.ScreenDraw(0, 10, 30, "white", screen, "High Score: "+strconv.Itoa(g.highScore))
		utils.ScreenDraw(0, 500, 10, "white", screen, "SOLO MODE")
	*/
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
	g.highScore = 0
	g.ball.Reset(false)
}

func (g *GameScene) OnEnter() {

}

func (g *GameScene) OnExit() {

}

func (g *GameScene) updateDimensions() {
	g.ball.W = config.GlobalConfig.BallSize
	g.ball.H = config.GlobalConfig.BallSize
	g.paddle.H = config.GlobalConfig.PaddleHeight
}

func (g *GameScene) Update() SceneId {

	g.updateDimensions()

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return PauseSceneId
	}

	if !g.paddle.MoveOnKeyPress(ebiten.KeyArrowUp, ebiten.KeyArrowDown) {
		g.paddle.MoveOnKeyPress(ebiten.KeyW, ebiten.KeyS)
	}

	g.ball.Move()

	if g.ball.X+g.ball.W >= config.GlobalConfig.ScreenWidth {
		g.score = 0
	}
	g.ball.CollideWithWall(false, true)

	if g.ball.CollideWithPaddle(g.paddle, true) {
		g.IncreaseScore()
		if g.score%5 == 0 {
			g.ball.IncreaseSpeed(2)
		}
	}

	return GameSceneId
}

var _ Scene = (*GameScene)(nil)

func (g *GameScene) IncreaseScore() {
	g.score++
	if g.score > g.highScore {
		g.highScore = g.score
	}
}
