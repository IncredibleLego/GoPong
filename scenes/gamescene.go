package scenes

import (
	_ "embed"
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
	loaded      bool
	paddle      *objects.Paddle
	enemyPaddle *objects.Paddle
	ball        *objects.Ball
	score       int
	highScore   int
}

func (g *GameScene) IsLoaded() bool {
	return g.loaded
}

func NewGameScene() *GameScene {
	return &GameScene{
		loaded:      false,
		paddle:      nil,
		enemyPaddle: nil,
		ball:        nil,
		score:       0,
		highScore:   0,
	}
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		float32(g.paddle.X), float32(g.paddle.Y),
		float32(g.paddle.W), float32(g.paddle.H),
		color.White, false,
	) // Draw the paddle
	vector.DrawFilledRect(screen,
		float32(g.enemyPaddle.X), float32(g.enemyPaddle.Y),
		float32(g.enemyPaddle.W), float32(g.enemyPaddle.H),
		color.White, false,
	) // Draw enemy paddle
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

	menu.ScreenDraw(13, 10, 10, 1, 1, 1, 1, 1.5, screen, "Score: "+strconv.Itoa(g.score))
	menu.ScreenDraw(13, 10, 30, 1, 1, 1, 1, 1.5, screen, "High Score: "+strconv.Itoa(g.highScore))

}

// FirstLoad implements Scene.
func (g *GameScene) FirstLoad() {
	g.paddle = &objects.Paddle{
		Object: &objects.Object{
			X: 600,
			Y: 200,
			W: 15,
			H: 100,
		},
	}
	g.enemyPaddle = &objects.Paddle{
		Object: &objects.Object{
			X: 40,
			Y: 200,
			W: 15,
			H: 100,
		},
	}
	g.ball = &objects.Ball{
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

func (g *GameScene) OnEnter() {

}

func (g *GameScene) OnExit() {

}

func (g *GameScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return PauseSceneId
	}
	g.paddle.MoveOnKeyPress()
	g.ball.Move()
	g.CollideWithWall()

	if g.paddle.CollideWithPaddle(g.ball) {
		g.ball.Dxdt = -g.ball.Dxdt
		g.IncreaseScore()
	}

	return GameSceneId
}

var _ Scene = (*GameScene)(nil)

func (g *GameScene) CollideWithWall() { // Check if the ball collides with the wall
	if g.ball.X >= constants.ScreenWidth {
		g.Reset()
	} else if g.ball.X <= 0 {
		g.ball.Dxdt = constants.BallSpeed
	} else if g.ball.Y <= 0 {
		g.ball.Dydt = constants.BallSpeed
	} else if g.ball.Y >= constants.ScreenHeight {
		g.ball.Dydt = -constants.BallSpeed
	}
}

func (g *GameScene) IncreaseScore() {
	g.score++
	if g.score > g.highScore {
		g.highScore = g.score
	}
}

func (g *GameScene) Reset() { // Reset the game
	g.ball.X = constants.ScreenWidth / 2
	g.ball.Y = constants.ScreenHeight / 2
	g.score = 0
}
