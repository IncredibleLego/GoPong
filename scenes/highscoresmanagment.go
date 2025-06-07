package scenes

import (
	"encoding/json"
	"os"
	"sort"
	"time"
)

const highscoresFile = "highscores.json"
const maxScores = 10

type SoloScore struct {
	DateTime string `json:"date_time"`
	Player   string `json:"player"`
	Score    int    `json:"score"`
}

type ComputerScore struct {
	DateTime string `json:"date_time"`
	Player   string `json:"player"`
	AILevel  string `json:"ai_level"`
	Score    int    `json:"score"`
}

type MultiplayerScore struct {
	DateTime string `json:"date_time"`
	Player1  string `json:"player1"`
	Player2  string `json:"player2"`
	Score    int    `json:"score"`
}

type Highscores struct {
	Solo        []SoloScore        `json:"solo"`
	Computer    []ComputerScore    `json:"computer"`
	Multiplayer []MultiplayerScore `json:"multiplayer"`
}

// Carica gli highscores dal file JSON
func loadHighscores() (*Highscores, error) {
	file, err := os.Open(highscoresFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &Highscores{}, nil // Nessun file, ritorna vuoto
		}
		return nil, err
	}
	defer file.Close()
	var hs Highscores
	err = json.NewDecoder(file).Decode(&hs)
	if err != nil {
		return nil, err
	}
	return &hs, nil
}

// Salva gli highscores nel file JSON
func saveHighscores(hs *Highscores) error {
	file, err := os.Create(highscoresFile)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(hs)
}

// Aggiungi un punteggio solo mode
func AddSoloScore(player string, score int) error {
	hs, err := loadHighscores()
	if err != nil {
		return err
	}
	hs.Solo = append(hs.Solo, SoloScore{
		DateTime: time.Now().Format(time.RFC3339),
		Player:   player,
		Score:    score,
	})
	sort.Slice(hs.Solo, func(i, j int) bool {
		return hs.Solo[i].Score > hs.Solo[j].Score
	})
	if len(hs.Solo) > maxScores {
		hs.Solo = hs.Solo[:maxScores]
	}
	return saveHighscores(hs)
}

// Aggiungi un punteggio computer mode
func AddComputerScore(player, aiLevel string, score int) error {
	hs, err := loadHighscores()
	if err != nil {
		return err
	}
	hs.Computer = append(hs.Computer, ComputerScore{
		DateTime: time.Now().Format(time.RFC3339),
		Player:   player,
		AILevel:  aiLevel,
		Score:    score,
	})
	sort.Slice(hs.Computer, func(i, j int) bool {
		return hs.Computer[i].Score > hs.Computer[j].Score
	})
	if len(hs.Computer) > maxScores {
		hs.Computer = hs.Computer[:maxScores]
	}
	return saveHighscores(hs)
}

// Aggiungi un punteggio multiplayer mode
func AddMultiplayerScore(player1, player2 string, score int) error {
	hs, err := loadHighscores()
	if err != nil {
		return err
	}
	hs.Multiplayer = append(hs.Multiplayer, MultiplayerScore{
		DateTime: time.Now().Format(time.RFC3339),
		Player1:  player1,
		Player2:  player2,
		Score:    score,
	})
	sort.Slice(hs.Multiplayer, func(i, j int) bool {
		return hs.Multiplayer[i].Score > hs.Multiplayer[j].Score
	})
	if len(hs.Multiplayer) > maxScores {
		hs.Multiplayer = hs.Multiplayer[:maxScores]
	}
	return saveHighscores(hs)
}

// Ottieni tutti gli highscores
func GetHighscores() (*Highscores, error) {
	return loadHighscores()
}
