package scenes

import (
	"goPong/config"
	"goPong/objects"
	"goPong/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	return reason == Unpause
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

	// Draw the paddle
	m.paddle1.Draw(screen)
	// Draw enemy
	m.paddle2.Draw(screen)

	// Draw the ball
	m.ball.Draw(screen)

	// Draw the net
	utils.Net(screen)

	// Draw the points
	utils.PointsDraw(screen, float32(config.GlobalConfig.ScreenWidth)/6, 50, m.score1)
	utils.PointsDraw(screen, (float32(config.GlobalConfig.ScreenWidth)/6)*4, 50, m.score2)

	utils.ScreenDraw(0, float64(config.GlobalConfig.ScreenWidth)/6, 10, "white", screen, m.player1Name)
	utils.ScreenDraw(0, (float64(config.GlobalConfig.ScreenWidth)/6)*4, 10, "white", screen, m.player2Name)
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
	m.ball.Reset(false)
}

func (m *MultiplayerScene) OnEnter() {

}

func (m *MultiplayerScene) OnExit() {

}

func (m *MultiplayerScene) updateDimensions() {
	m.ball.W = config.GlobalConfig.BallSize
	m.ball.H = config.GlobalConfig.BallSize
	m.paddle1.H = config.GlobalConfig.PaddleHeight
	m.paddle2.H = config.GlobalConfig.PaddleHeight
}

func (m *MultiplayerScene) Update() SceneId {

	m.updateDimensions()

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return PauseSceneId
	}
	m.paddle1.MoveOnKeyPress(ebiten.KeyArrowUp, ebiten.KeyArrowDown)
	m.paddle2.MoveOnKeyPress(ebiten.KeyW, ebiten.KeyS)
	m.ball.Move()

	test := m.ball.CollideWithWall(true, true)
	if test == 1 {
		m.score2++
	} else if test == 2 {
		m.score1++
	}

	m.ball.CollideWithPaddle(m.paddle1, true)
	m.ball.CollideWithPaddle(m.paddle2, false)

	return MultiplayerSceneId
}

var _ Scene = (*MultiplayerScene)(nil)
