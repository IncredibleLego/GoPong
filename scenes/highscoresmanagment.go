package scenes

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

const highscoresFile = "./scenes/highscores.json"
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

// Load all highscores from the JSON file
func loadHighscores() (*Highscores, error) {
	file, err := os.Open(highscoresFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &Highscores{}, nil // No highscores file exists, return empty highscores
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

// Save highscores to the JSON file
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

// Add a solo score
func AddSoloScore(player string, score int) error {
	hs, err := loadHighscores()
	if err != nil {
		return err
	}
	// Store the date in RFC3339 for consistency, but display will format it
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

// Add a computer score
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

// Add a multiplayer score
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

// Get formatted solo highscores as a slice of strings (error handled internally)
func GetSoloHighscoresStrings() []string {
	hs, err := loadHighscores()
	if err != nil {
		return []string{"Error loading highscores"}
	}
	var result []string

	maxPlayerLen := 0
	for _, s := range hs.Solo {
		if len(s.Player) > maxPlayerLen {
			maxPlayerLen = len(s.Player)
		}
	}
	for i, s := range hs.Solo {
		// Parse the RFC3339 date to time.Time
		t, err := time.Parse(time.RFC3339, s.DateTime)
		dateStr := s.DateTime
		if err == nil {
			dateStr = t.Format("02/01/2006 15:04")
		}
		result = append(result, fmt.Sprintf(
			"%2d. %-*s  Score %-4d    %s",
			i+1,
			maxPlayerLen, s.Player,
			s.Score,
			dateStr,
		))
	}
	return result
}

// Get formatted computer highscores as a slice of strings (error handled internally)
func GetComputerHighscoresStrings() []string {
	hs, err := loadHighscores()
	if err != nil {
		return []string{"Error loading highscores"}
	}
	var result []string

	// Find max lengths for alignment
	maxPlayerLen := 0
	maxAILevelLen := 0
	for _, s := range hs.Computer {
		if len(s.Player) > maxPlayerLen {
			maxPlayerLen = len(s.Player)
		}
		if len(s.AILevel) > maxAILevelLen {
			maxAILevelLen = len(s.AILevel)
		}
	}
	for i, s := range hs.Computer {
		t, err := time.Parse(time.RFC3339, s.DateTime)
		dateStr := s.DateTime
		if err == nil {
			dateStr = t.Format("02/01/2006 15:04")
		}
		result = append(result, fmt.Sprintf(
			"%2d. %-*s  Score %-4d  Difficulty: %-*s  %s",
			i+1,
			maxPlayerLen, s.Player,
			s.Score,
			maxAILevelLen, s.AILevel,
			dateStr,
		))
	}
	return result
}

// Get formatted multiplayer highscores as a slice of strings (error handled internally)
func GetMultiplayerHighscoresStrings() []string {
	hs, err := loadHighscores()
	if err != nil {
		return []string{"Error loading highscores"}
	}
	var result []string

	// Find max lengths for alignment
	maxPlayer1Len := 0
	maxPlayer2Len := 0
	for _, s := range hs.Multiplayer {
		if len(s.Player1) > maxPlayer1Len {
			maxPlayer1Len = len(s.Player1)
		}
		if len(s.Player2) > maxPlayer2Len {
			maxPlayer2Len = len(s.Player2)
		}
	}

	for i, s := range hs.Multiplayer {
		t, err := time.Parse(time.RFC3339, s.DateTime)
		dateStr := s.DateTime
		if err == nil {
			dateStr = t.Format("02/01/2006 15:04")
		}
		result = append(result, fmt.Sprintf(
			"%2d. %-*s  Score %-4d vs %-*s  %s",
			i+1,
			maxPlayer1Len, s.Player1,
			s.Score,
			maxPlayer2Len, s.Player2,
			dateStr,
		))
	}
	return result
}
