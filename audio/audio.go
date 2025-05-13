package audio

import (
	"bytes"
	_ "embed"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

//go:embed paddle/pong1.mp3
var pong1 []byte

//go:embed paddle/pong2.mp3
var pong2 []byte

//go:embed paddle/pong3.mp3
var pong3 []byte

//go:embed paddle/pong4.mp3
var pong4 []byte

//go:embed paddle/pong5.mp3
var pong5 []byte

//go:embed paddle/pong6.mp3
var pong6 []byte

//go:embed score/score.mp3
var score []byte

var (
	audioContext *audio.Context
	paddleSounds [][]byte
)

func Init() {
	audioContext = audio.NewContext(44100)
	paddleSounds = [][]byte{pong1, pong2, pong3, pong4, pong5, pong6}
	rand.Seed(time.Now().UnixNano())
}

func PlayPaddle() {
	playMp3(paddleSounds[rand.Intn(len(paddleSounds))])
}

func PlayScore() {
	playMp3(score)
}

func playMp3(data []byte) {
	if audioContext == nil {
		log.Println("Audio context non inizializzato")
		return
	}
	stream, err := mp3.DecodeWithSampleRate(44100, bytes.NewReader(data))
	if err != nil {
		log.Println("Errore decoding mp3:", err)
		return
	}
	player, err := audioContext.NewPlayer(stream)
	if err != nil {
		log.Println("Errore creazione player:", err)
		return
	}
	player.Play()
}
