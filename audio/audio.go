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
	paddleBuffers [][]byte
	scoreBuffer   []byte
)

func Init() {
	audioContext = audio.NewContext(44100)
	rand.Seed(time.Now().UnixNano())

	// Decodifica e salva i buffer PCM una sola volta
	paddleBuffers = make([][]byte, 6)
	paddleBuffers[0] = decodeToPCM(pong1)
	paddleBuffers[1] = decodeToPCM(pong2)
	paddleBuffers[2] = decodeToPCM(pong3)
	paddleBuffers[3] = decodeToPCM(pong4)
	paddleBuffers[4] = decodeToPCM(pong5)
	paddleBuffers[5] = decodeToPCM(pong6)
	scoreBuffer = decodeToPCM(score)
}

func decodeToPCM(mp3data []byte) []byte {
	stream, err := mp3.DecodeWithSampleRate(44100, bytes.NewReader(mp3data))
	if err != nil {
		log.Println("Errore decoding mp3:", err)
		return nil
	}
	// stream non ha Close(), quindi non serve defer stream.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(stream)
	if err != nil {
		log.Println("Errore bufferizzazione stream:", err)
		return nil
	}
	return buf.Bytes()
}

func PlayPaddle() {
	playPCM(paddleBuffers[rand.Intn(len(paddleBuffers))])
}

func PlayScore() {
	playPCM(scoreBuffer)
}

func playPCM(pcm []byte) {
	if audioContext == nil || pcm == nil {
		return
	}
	player := audioContext.NewPlayerFromBytes(pcm)
	player.Play()
}
