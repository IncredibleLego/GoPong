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

type GameScene struct {
	loaded    bool
	paddle    *objects.Paddle
	ball      *objects.Ball
	score     int
	highScore int
}

func (g *GameScene) IsLoaded() bool {
	return g.loaded
}

func NewGameScene() *GameScene {
	return &GameScene{
		loaded:    false,
		paddle:    nil,
		ball:      nil,
		score:     0,
		highScore: 0,
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

	for i := 0; i < constants.ScreenHeight; i += 24 {
		vector.DrawFilledRect(screen,
			float32(constants.ScreenWidth/2), float32(i),
			float32(3), float32(12),
			color.White, false,
		)
	}

	menu.ScreenDraw(constants.TextDimension, 10, 10, 1, 1, 1, 1, 1.5, screen, "Score: "+strconv.Itoa(g.score))
	menu.ScreenDraw(constants.TextDimension, 10, 30, 1, 1, 1, 1, 1.5, screen, "High Score: "+strconv.Itoa(g.highScore))
	menu.ScreenDraw(constants.TextDimension, 500, 10, 1, 1, 1, 1, 1.5, screen, "SOLO MODE")

	//Debug
	x := strconv.Itoa(g.ball.Dxdt)
	y := strconv.Itoa(g.ball.Dydt)

	menu.ScreenDraw(constants.TextDimension, 300, 10, 1, 1, 1, 1, 1.5, screen, x)
	menu.ScreenDraw(constants.TextDimension, 320, 10, 1, 1, 1, 1, 1.5, screen, y)
}

// FirstLoad implements Scene.
func (g *GameScene) FirstLoad() {
	g.paddle = &objects.Paddle{
		Object: &objects.Object{
			X: constants.ScreenWidth - constants.PaddleDistanceFromWall,
			Y: constants.ScreenHeight/2 - constants.PaddleHeight/2,
			W: 15,
			H: constants.PaddleHeight,
		},
	}
	g.ball = &objects.Ball{
		Object: &objects.Object{
			X: constants.ScreenWidth / 2,
			Y: constants.ScreenHeight / 2,
			W: constants.BallSize,
			H: constants.BallSize,
		},
		Dxdt: constants.BallSpeed,
		Dydt: constants.BallSpeed,
	}
	g.ball.GenerateRandomDirection()
}

func (g *GameScene) OnEnter() {

}

func (g *GameScene) OnExit() {

}

func (g *GameScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return PauseSceneId
	}

	if !g.paddle.MoveOnKeyPress(ebiten.KeyArrowUp, ebiten.KeyArrowDown) {
		g.paddle.MoveOnKeyPress(ebiten.KeyW, ebiten.KeyS)
	}

	g.ball.Move()

	if g.ball.X+g.ball.W >= constants.ScreenWidth {
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
