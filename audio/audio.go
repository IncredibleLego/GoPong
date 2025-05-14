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
	audioContext  *audio.Context
	paddlePlayers []*audio.Player
	scorePlayer   *audio.Player
)

func Init() {
	audioContext = audio.NewContext(44100)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Decodifica e crea i player una sola volta
	paddlePlayers = make([]*audio.Player, 6)
	paddleBuffers := [][]byte{
		decodeToPCM(pong1),
		decodeToPCM(pong2),
		decodeToPCM(pong3),
		decodeToPCM(pong4),
		decodeToPCM(pong5),
		decodeToPCM(pong6),
	}
	for i, pcm := range paddleBuffers {
		if pcm != nil {
			paddlePlayers[i] = audioContext.NewPlayerFromBytes(pcm)
		}
	}
	scorePCM := decodeToPCM(score)
	if scorePCM != nil {
		scorePlayer = audioContext.NewPlayerFromBytes(scorePCM)
	}
}

func decodeToPCM(mp3data []byte) []byte {
	stream, err := mp3.DecodeWithSampleRate(44100, bytes.NewReader(mp3data))
	if err != nil {
		log.Println("Error decoding mp3:", err)
		return nil
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(stream)
	if err != nil {
		log.Println("Error buffer stream:", err)
		return nil
	}
	return buf.Bytes()
}

func PlayPaddle() {
	if len(paddlePlayers) == 0 {
		return
	}
	player := paddlePlayers[rand.Intn(len(paddlePlayers))]
	if player == nil {
		return
	}
	player.Rewind()
	player.Play()
}

func PlayScore() {
	if scorePlayer == nil {
		return
	}
	scorePlayer.Rewind()
	scorePlayer.Play()
}
