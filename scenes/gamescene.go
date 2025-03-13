package scenes

import (
	"bytes"
	_ "embed"
	"image/color"
	"log"
	"main/constants"
	"main/objects"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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

	// Text Options
	s, err := text.NewGoTextFaceSource(bytes.NewReader(pressStart2P))
	if err != nil {
		log.Fatal(err)
	}
	pressStart2PFaceSource = s

	scoreTextOptions := &text.DrawOptions{}
	scoreTextOptions.GeoM.Translate(10, 10)
	scoreTextOptions.ColorScale.Scale(1, 1, 1, 1)
	scoreTextOptions.LineSpacing = 1.5

	highScoreTextOptions := &text.DrawOptions{}
	highScoreTextOptions.GeoM.Translate(10, 30)
	highScoreTextOptions.ColorScale.Scale(1, 1, 1, 1)
	highScoreTextOptions.LineSpacing = 1.5

	textFace := &text.GoTextFace{
		Source: pressStart2PFaceSource,
		Size:   13,
	}

	scoreStr := "Score: " + strconv.Itoa(g.score)
	text.Draw(screen, scoreStr, textFace, scoreTextOptions)

	HighScoreStr := "High Score: " + strconv.Itoa(g.highScore)
	text.Draw(screen, HighScoreStr, textFace, highScoreTextOptions)

}

// FirstLoad implements Scene.
func (g *GameScene) FirstLoad() {
	g.paddle = &objects.Paddle{
		Ojbect: &objects.Ojbect{
			X: 600,
			Y: 200,
			W: 15,
			H: 100,
		},
	}
	g.enemyPaddle = &objects.Paddle{
		Ojbect: &objects.Ojbect{
			X: 40,
			Y: 200,
			W: 15,
			H: 100,
		},
	}
	g.ball = &objects.Ball{
		Ojbect: &objects.Ojbect{
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
	g.CollideWithPaddle()
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

func (g *GameScene) CollideWithPaddle() { // Check if the ball collides with the paddle
	if g.ball.X+g.ball.W >= g.paddle.X && g.ball.Y+g.ball.W >= g.paddle.Y && g.ball.Y+g.ball.W <= g.paddle.Y+g.paddle.H {
		g.ball.Dxdt = -g.ball.Dxdt
		g.IncreaseScore()
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
