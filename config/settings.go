package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Player1Name            string
	Player2Name            string
	ScreenWidth            int
	ScreenHeight           int
	BallSpeed              int
	PaddleSpeed            int
	MenuOptionsPerSecond   time.Duration
	TextDimension          float64
	PaddleDistanceFromWall int
	MaxBounceAngle         float64
	Difficulty             float64
	PaddleHeight           int
	BallSize               int
}

var GlobalConfig = &Config{
	Player1Name:            "Player 1",
	Player2Name:            "Player 2",
	ScreenWidth:            640,
	ScreenHeight:           480,
	BallSpeed:              6,
	PaddleSpeed:            6,
	MenuOptionsPerSecond:   4,
	TextDimension:          20,
	PaddleDistanceFromWall: 40,
	MaxBounceAngle:         0.7853975, //45.0 * (3.14159 / 180.0)
	Difficulty:             0.5,
	PaddleHeight:           100,
	BallSize:               15,
}

var DefaultConfig = &Config{
	Player1Name:            "Player 1",
	Player2Name:            "Player 2",
	ScreenWidth:            640,
	ScreenHeight:           480,
	BallSpeed:              6,
	PaddleSpeed:            6,
	MenuOptionsPerSecond:   4,
	TextDimension:          13,
	PaddleDistanceFromWall: 40,
	MaxBounceAngle:         0.7853975, //45.0 * (3.14159 / 180.0)
	Difficulty:             0.5,
	PaddleHeight:           100,
	BallSize:               15,
}

const configFilePath = "settings.json"

// SaveConfig salva la configurazione corrente in un file JSON.
func SaveConfig(config *Config) error {
	// Converte la configurazione in JSON formattato
	data, err := json.MarshalIndent(config, "", "  ") // "" è il prefisso e "  " è l'indentazione
	if err != nil {
		return err
	}

	// Scrive il JSON formattato nel file
	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

// LoadConfig carica la configurazione da un file JSON.
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// UpdateConfig aggiorna un campo della configurazione e salva il file JSON.
func UpdateConfig(updateFunc func(*Config)) error {
	updateFunc(GlobalConfig)
	return SaveConfig(GlobalConfig)
}

// InitConfig inizializza la configurazione caricandola dal file o usando quella di default.
func InitConfig() {
	config, err := LoadConfig(configFilePath)
	if err != nil {
		fmt.Println("Impossibile caricare il file di configurazione, uso la configurazione di default:", err)
		//GlobalConfig = GlobalConfig
		_ = SaveConfig(GlobalConfig) // Salva la configurazione di default
	} else {
		GlobalConfig = config
	}
}
