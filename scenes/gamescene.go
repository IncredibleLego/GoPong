package scenes

import (
	"goPong/config"
	"goPong/objects"
	"goPong/utils"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	vector.DrawFilledRect(screen,
		float32(g.paddle.X), float32(g.paddle.Y),
		float32(g.paddle.W), float32(g.paddle.H),
		color.White, false,
	) // Draw the paddle
	vector.DrawFilledRect(screen,
		float32(g.ball.X), float32(g.ball.Y),
		float32(g.ball.W), float32(g.ball.H),
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

	//utils.PointsDraw(screen, 170, 100, g.score%10) // Draw the first digit of the score
	//utils.PointsDraw(screen, 100, 100, g.score/10) // Draw the second digit of the score

	utils.PointsDraw(screen, 170, 100, 2)

	//Array of all characters of the points, to draw
	//Max points 99

	utils.ScreenDraw(0, 10, 10, "white", screen, "Score "+g.playerName+":"+strconv.Itoa(g.score))
	utils.ScreenDraw(0, 10, 30, "white", screen, "High Score: "+strconv.Itoa(g.highScore))
	utils.ScreenDraw(0, 500, 10, "white", screen, "SOLO MODE")
}

// FirstLoad implements Scene.
func (g *GameScene) FirstLoad() {
	g.playerName = config.GlobalConfig.Player1Name
	g.paddle = &objects.Paddle{
		Object: &objects.Object{
			X: config.GlobalConfig.ScreenWidth - config.GlobalConfig.PaddleDistanceFromWall,
			Y: config.GlobalConfig.ScreenHeight/2 - config.GlobalConfig.PaddleHeight/2,
			W: 15,
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
