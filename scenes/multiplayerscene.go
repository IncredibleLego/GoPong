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

type MultiplayerScene struct {
	loaded    bool
	paddle1   *objects.Paddle
	paddle2   *objects.Paddle
	ball      *objects.Ball
	score1    int
	score2    int
	highScore int
}

func (m *MultiplayerScene) IsLoaded() bool {
	return m.loaded
}

func NewMultiplayerScene() *MultiplayerScene {
	return &MultiplayerScene{
		loaded:    false,
		paddle1:   nil,
		paddle2:   nil,
		ball:      nil,
		score1:    0,
		score2:    0,
		highScore: 0,
	}
}

func (m *MultiplayerScene) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		float32(m.paddle1.X), float32(m.paddle1.Y),
		float32(m.paddle1.W), float32(m.paddle1.H),
		color.White, false,
	) // Draw the paddle
	vector.DrawFilledRect(screen,
		float32(m.paddle2.X), float32(m.paddle2.Y),
		float32(m.paddle2.W), float32(m.paddle2.H),
		color.White, false,
	) // Draw enemy paddle
	vector.DrawFilledRect(screen,
		float32(m.ball.X), float32(m.ball.Y),
		float32(m.ball.W), float32(m.ball.H),
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

	menu.ScreenDraw(constants.TextDimension, 10, 10, 1, 1, 1, 1, 1.5, screen, "Score: "+strconv.Itoa(m.score1))
	menu.ScreenDraw(constants.TextDimension, 500, 10, 1, 1, 1, 1, 1.5, screen, "Score: "+strconv.Itoa(m.score2))
	menu.ScreenDraw(constants.TextDimension-3, 250, 10, 1, 1, 1, 1, 1.5, screen, "MULTIPLAYER MODE")
}

// FirstLoad implements Scene.
func (m *MultiplayerScene) FirstLoad() {
	m.paddle1 = &objects.Paddle{
		Object: &objects.Object{
			X: constants.ScreenWidth - constants.PaddleDistanceFromWall,
			Y: constants.ScreenHeight / 2,
			W: 15,
			H: 100,
		},
	}
	m.paddle2 = &objects.Paddle{
		Object: &objects.Object{
			X: constants.PaddleDistanceFromWall,
			Y: constants.ScreenHeight / 2,
			W: 15,
			H: 100,
		},
	}
	m.ball = &objects.Ball{
		Object: &objects.Object{
			X: constants.ScreenWidth / 2,
			Y: constants.ScreenHeight / 2,
			W: 15,
			H: 15,
		},
		Dxdt: constants.BallSpeed,
		Dydt: constants.BallSpeed,
	}
	m.ball.GenerateRandomDirection()
}

func (g *MultiplayerScene) OnEnter() {

}

func (g *MultiplayerScene) OnExit() {

}

func (m *MultiplayerScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return PauseSceneId
	}
	m.paddle1.MoveOnKeyPress(ebiten.KeyArrowUp, ebiten.KeyArrowDown)
	m.paddle2.MoveOnKeyPress(ebiten.KeyW, ebiten.KeyS)
	m.ball.Move()
	m.ball.CollideWithWall(true, true)

	if m.ball.CollideWithWall(true, true) == 1 {
		m.score1++
	} else if m.ball.CollideWithWall(true, true) == 2 {
		m.score2++
	}

	m.paddle1.CollideWithPaddle(m.ball, true)
	m.paddle2.CollideWithPaddle(m.ball, false)

	return MultiplayerSceneId
}

var _ Scene = (*MultiplayerScene)(nil)
