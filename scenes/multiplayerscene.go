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

type MultiplayerScene struct {
	player1Name string
	player2Name string
	paddle1     *objects.Paddle
	paddle2     *objects.Paddle
	ball        *objects.Ball
	score1      int
	score2      int
	highScore   int
}

func (m *MultiplayerScene) ShouldPreserveState(reason SceneChangeReason) bool {
	if reason == Unpause {
		return true
	}
	return false
}

func NewMultiplayerScene() *MultiplayerScene {
	return &MultiplayerScene{
		player1Name: "",
		player2Name: "",
		paddle1:     nil,
		paddle2:     nil,
		ball:        nil,
		score1:      0,
		score2:      0,
		highScore:   0,
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

	for i := 0; i < config.GlobalConfig.ScreenHeight; i += 24 {
		vector.DrawFilledRect(screen,
			float32(config.GlobalConfig.ScreenWidth/2), float32(i),
			float32(3), float32(12),
			color.White, false,
		)
	}

	menu.ScreenDraw(config.GlobalConfig.TextDimension, 10, 10, 1, 1, 1, 1, 1.5, screen, m.player1Name)
	menu.ScreenDraw(config.GlobalConfig.TextDimension, 10, 25, 1, 1, 1, 1, 1.5, screen, "Score: "+strconv.Itoa(m.score1))
	menu.ScreenDraw(config.GlobalConfig.TextDimension, 500, 10, 1, 1, 1, 1, 1.5, screen, m.player2Name)
	menu.ScreenDraw(config.GlobalConfig.TextDimension, 500, 25, 1, 1, 1, 1, 1.5, screen, "Score: "+strconv.Itoa(m.score2))
	menu.ScreenDraw(config.GlobalConfig.TextDimension-3, 250, 10, 1, 1, 1, 1, 1.5, screen, "MULTIPLAYER MODE")
}

// FirstLoad implements Scene.
func (m *MultiplayerScene) FirstLoad() {
	m.player1Name = config.GlobalConfig.Player1Name
	m.player2Name = config.GlobalConfig.Player2Name
	m.paddle1 = &objects.Paddle{
		Object: &objects.Object{
			X: config.GlobalConfig.ScreenWidth - config.GlobalConfig.PaddleDistanceFromWall,
			Y: config.GlobalConfig.ScreenHeight/2 - config.GlobalConfig.PaddleHeight/2,
			W: 15,
			H: config.GlobalConfig.PaddleHeight,
		},
	}
	m.paddle2 = &objects.Paddle{
		Object: &objects.Object{
			X: config.GlobalConfig.PaddleDistanceFromWall,
			Y: config.GlobalConfig.ScreenHeight/2 - config.GlobalConfig.PaddleHeight/2,
			W: 15,
			H: config.GlobalConfig.PaddleHeight,
		},
	}
	m.ball = &objects.Ball{
		Object: &objects.Object{
			X: config.GlobalConfig.ScreenWidth / 2,
			Y: config.GlobalConfig.ScreenHeight / 2,
			W: config.GlobalConfig.BallSize,
			H: config.GlobalConfig.BallSize,
		},
		Dxdt: config.GlobalConfig.BallSpeed,
		Dydt: config.GlobalConfig.BallSpeed,
	}
	m.ball.GenerateRandomDirection()
	m.score1 = 0
	m.score2 = 0
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

	m.ball.CollideWithPaddle(m.paddle1, true)
	m.ball.CollideWithPaddle(m.paddle2, false)

	return MultiplayerSceneId
}

var _ Scene = (*MultiplayerScene)(nil)
