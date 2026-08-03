package scenes

import (
	"encoding/json"
	"fmt"
	"goPong/config"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var highscoresFile string

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
	AIScore  int    `json:"ai_score"`
}

type MultiplayerScore struct {
	DateTime   string `json:"date_time"`
	Player1    string `json:"player1"`
	Player2    string `json:"player2"`
	Score      int    `json:"score"`
	EnemyScore int    `json:"enemy_score"`
}

type Highscores struct {
	Solo            []SoloScore        `json:"solo"`
	ComputerEasy    []ComputerScore    `json:"computer_easy"`
	ComputerDefault []ComputerScore    `json:"computer_default"`
	ComputerHard    []ComputerScore    `json:"computer_hard"`
	Multiplayer     []MultiplayerScore `json:"multiplayer"`
}

// Load all highscores from the JSON file
func loadHighscores() (*Highscores, error) {

	if highscoresFile == "" {
		highscoresFile = GetHighscoresPath()
	}

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

// Returns the score to beat to achieve a new Solo Highscore
func getTopSoloScore() (SoloScore, bool) {
	hs, _ := loadHighscores()
	if len(hs.Solo) == 0 {
		return SoloScore{}, false
	}
	if len(hs.Solo) == maxScores {
		return hs.Solo[len(hs.Solo)-1], true
	}
	return hs.Solo[len(hs.Solo)-1], false
}

// Returns the score to beat to achieve a new Computer Highscore
func getTopComputerScore() (ComputerScore, bool) {
	hs, _ := loadHighscores()
	var topScore []ComputerScore
	if config.GlobalConfig.Difficulty < 0.33 {
		topScore = hs.ComputerEasy
	} else if config.GlobalConfig.Difficulty >= 0.33 && config.GlobalConfig.Difficulty < 0.66 {
		topScore = hs.ComputerDefault
	} else {
		topScore = hs.ComputerHard
	}
	if len(topScore) == 0 {
		return ComputerScore{}, false
	}
	if len(topScore) == maxScores {
		return topScore[len(topScore)-1], true
	}
	return topScore[len(topScore)-1], false
}

// Returns the score to beat to achieve a new Multiplayer Highscore
func getTopMultiplayerScore() (MultiplayerScore, bool) {
	hs, _ := loadHighscores()
	if len(hs.Multiplayer) == 0 {
		return MultiplayerScore{}, false
	}
	if len(hs.Multiplayer) == maxScores {
		return hs.Multiplayer[len(hs.Multiplayer)-1], true
	}
	return hs.Multiplayer[len(hs.Multiplayer)-1], false
}

// Save highscores to the JSON file
func saveHighscores(hs *Highscores) error {

	if highscoresFile == "" {
		highscoresFile = GetHighscoresPath()
	}

	_ = os.MkdirAll(filepath.Dir(highscoresFile), 0755)
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
func AddSoloScore(score SoloScore) error {
	hs, err := loadHighscores()
	if err != nil {
		return err
	}
	// Store the date in RFC3339 for consistency, but display will format it
	hs.Solo = append(hs.Solo, score)
	sort.Slice(hs.Solo, func(i, j int) bool {
		return hs.Solo[i].Score > hs.Solo[j].Score
	})
	if len(hs.Solo) > maxScores {
		hs.Solo = hs.Solo[:maxScores]
	}
	return saveHighscores(hs)
}

// Add a computer score
func AddComputerScore(score ComputerScore) error {
	hs, err := loadHighscores()
	if err != nil {
		return err
	}
	if config.GlobalConfig.Difficulty < 0.33 {

		hs.ComputerEasy = append(hs.ComputerEasy, score)
		sort.Slice(hs.ComputerEasy, func(i, j int) bool {
			if hs.ComputerEasy[i].Score != hs.ComputerEasy[j].Score {
				return hs.ComputerEasy[i].Score > hs.ComputerEasy[j].Score
			}
			// If scores are equal, prefer the one who conceded fewer goals (lower AIScore)
			return hs.ComputerEasy[i].AIScore < hs.ComputerEasy[j].AIScore
		})
		if len(hs.ComputerEasy) > maxScores {
			hs.ComputerEasy = hs.ComputerEasy[:maxScores]
		}
		return saveHighscores(hs)

	} else if config.GlobalConfig.Difficulty >= 0.33 && config.GlobalConfig.Difficulty < 0.66 {
		hs.ComputerDefault = append(hs.ComputerDefault, score)
		sort.Slice(hs.ComputerDefault, func(i, j int) bool {
			if hs.ComputerDefault[i].Score != hs.ComputerDefault[j].Score {
				return hs.ComputerDefault[i].Score > hs.ComputerDefault[j].Score
			}
			return hs.ComputerDefault[i].AIScore < hs.ComputerDefault[j].AIScore
		})
		if len(hs.ComputerDefault) > maxScores {
			hs.ComputerDefault = hs.ComputerDefault[:maxScores]
		}
		return saveHighscores(hs)
	} else {
		hs.ComputerHard = append(hs.ComputerHard, score)
		sort.Slice(hs.ComputerHard, func(i, j int) bool {
			if hs.ComputerHard[i].Score != hs.ComputerHard[j].Score {
				return hs.ComputerHard[i].Score > hs.ComputerHard[j].Score
			}
			return hs.ComputerHard[i].AIScore < hs.ComputerHard[j].AIScore
		})
		if len(hs.ComputerHard) > maxScores {
			hs.ComputerHard = hs.ComputerHard[:maxScores]
		}
		return saveHighscores(hs)
	}
}

// Add a multiplayer score
func AddMultiplayerScore(score MultiplayerScore) error {
	hs, err := loadHighscores()
	if err != nil {
		return err
	}
	hs.Multiplayer = append(hs.Multiplayer, score)
	sort.Slice(hs.Multiplayer, func(i, j int) bool {
		if hs.Multiplayer[i].Score != hs.Multiplayer[j].Score {
			return hs.Multiplayer[i].Score > hs.Multiplayer[j].Score
		}
		// If scores are equal, prefer the match with fewer goals conceded (lower EnemyScore)
		return hs.Multiplayer[i].EnemyScore < hs.Multiplayer[j].EnemyScore
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
func GetComputerHighscoresStrings(selected int) []string {
	hs, err := loadHighscores()
	if err != nil {
		return []string{"Error loading highscores"}
	}
	var result []string
	var computer []ComputerScore
	if selected == 0 {
		computer = hs.ComputerEasy
	} else if selected == 1 {
		computer = hs.ComputerDefault
	} else {
		computer = hs.ComputerHard
	}

	// Find max lengths for alignment
	maxPlayerLen := 0
	maxAILevelLen := 0
	for _, s := range computer {
		if len(s.Player) > maxPlayerLen {
			maxPlayerLen = len(s.Player)
		}
		if len(s.AILevel) > maxAILevelLen {
			maxAILevelLen = len(s.AILevel)
		}
	}
	for i, s := range computer {
		t, err := time.Parse(time.RFC3339, s.DateTime)
		dateStr := s.DateTime
		if err == nil {
			dateStr = t.Format("02/01/2006 15:04")
		}
		result = append(result, fmt.Sprintf(
			"%2d. %-*s Score %-2d vs %-2d Mode: %-*s %s",
			i+1,
			maxPlayerLen, s.Player,
			s.Score,
			s.AIScore,
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
			"%2d. %-*s Score %d vs %d %-*s %s",
			i+1,
			maxPlayer1Len, s.Player1,
			s.Score,
			s.EnemyScore,
			maxPlayer2Len, s.Player2,
			dateStr,
		))
	}
	return result
}

// GetHighscoresPath returns the path to the highscores file based on the environment (production or development)
func GetHighscoresPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "./highscores.json"
	}
	gameDir := filepath.Join(configDir, "goPong")
	_ = os.MkdirAll(gameDir, 0755)
	return filepath.Join(gameDir, "highscores.json")
}
