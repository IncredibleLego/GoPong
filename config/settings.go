package config

import "time"

type Config struct {
	Player1Name            string
	Player2Name            string
	ScreenWidth            int
	ScreenHeight           int
	BallSpeed              int
	PaddleSpeed            int
	MenuOptionsPerSecond   time.Duration
	TextDimension          int
	PaddleDistanceFromWall int
	MaxBounceAngle         float64
	Difficulty             float64
	PaddleHeight           int
	BallSize               int
}

var GlobalConfig = &Config{
	ScreenWidth:            640,
	ScreenHeight:           480,
	BallSpeed:              6,
	PaddleSpeed:            6,
	MenuOptionsPerSecond:   4,
	TextDimension:          13,
	PaddleDistanceFromWall: 40,
	MaxBounceAngle:         45.0 * (3.14159 / 180.0),
	Difficulty:             0.5,
	PaddleHeight:           100,
	BallSize:               15,
}
