package config

import (
	"encoding/json"
	"math"
	"os"
	"time"
)

type Config struct {
	Scale                  float64
	Fullscreen             bool
	Player1Name            string
	Player2Name            string
	BallSpeed              int
	BallSize               int
	PaddleSpeed            int
	PaddleHeight           int
	PaddleDistanceFromWall int
	PaddleWidth            int
	Difficulty             float64
	TextDimension          float64
	ScreenWidth            int
	ScreenHeight           int
	PopupWidth             int
	PopupHeight            int
	OptionsPerSecond       time.Duration
	MaxBounceAngle         float64
}

var GlobalConfig = &Config{
	Scale:                  1.0,
	Fullscreen:             true,
	Player1Name:            "Player 1",
	Player2Name:            "Player 2",
	BallSpeed:              9,
	BallSize:               22,
	PaddleSpeed:            9,
	PaddleHeight:           150,
	PaddleDistanceFromWall: 60,
	PaddleWidth:            22,
	Difficulty:             0.5,
	TextDimension:          30,
	ScreenWidth:            960,
	ScreenHeight:           720,
	PopupWidth:             528, // 55% of height
	PopupHeight:            216, // 30% of height
	OptionsPerSecond:       time.Duration(time.Second / 4),
	MaxBounceAngle:         0.7853975, //45.0 * (3.14159 / 180.0)
}

var DefaultConfig = &Config{
	Scale:                  1.0,
	Fullscreen:             true,
	Player1Name:            "Player 1",
	Player2Name:            "Player 2",
	BallSpeed:              9,
	BallSize:               22,
	PaddleSpeed:            9,
	PaddleHeight:           150,
	PaddleDistanceFromWall: 60,
	PaddleWidth:            22,
	Difficulty:             0.5,
	TextDimension:          30,
	ScreenWidth:            960,
	ScreenHeight:           720,
	PopupWidth:             528, // 55% of height
	PopupHeight:            216, // 30% of height
	OptionsPerSecond:       time.Duration(time.Second / 4),
	MaxBounceAngle:         0.7853975, //45.0 * (3.14159 / 180.0)
}

// Applica la scala ai valori di default e aggiorna la config
func ApplyScaleToConfig(cfg *Config, scale float64) {
	cfg.Scale = scale
	cfg.BallSpeed = int(math.Round(float64(DefaultConfig.BallSpeed) * scale))
	cfg.BallSize = int(math.Round(float64(DefaultConfig.BallSize) * scale))
	cfg.PaddleSpeed = int(math.Round(float64(DefaultConfig.PaddleSpeed) * scale))
	cfg.PaddleHeight = int(math.Round(float64(DefaultConfig.PaddleHeight) * scale))
	cfg.PaddleDistanceFromWall = int(math.Round(float64(DefaultConfig.PaddleDistanceFromWall) * scale))
	cfg.PaddleWidth = int(math.Round(float64(DefaultConfig.PaddleWidth) * scale))
	cfg.ScreenWidth = ((int(math.Round(float64(DefaultConfig.ScreenWidth)*scale)) + 5) / 10) * 10
	cfg.ScreenHeight = ((int(math.Round(float64(DefaultConfig.ScreenHeight)*scale)) + 5) / 10) * 10
	cfg.PopupWidth = int(math.Round(float64(DefaultConfig.PopupWidth) * scale))
	cfg.PopupHeight = int(math.Round(float64(DefaultConfig.PopupHeight) * scale))
	cfg.TextDimension = math.Round(DefaultConfig.TextDimension * scale)
}

// Cambia la scala e aggiorna la configurazione globale
func ChangeScale(newScale float64) error {
	ApplyScaleToConfig(GlobalConfig, newScale)
	return SaveConfig(GlobalConfig)
}

// DifficultyString returns a string representation of the difficulty level.
func DifficultyString() string {
	if GlobalConfig.Difficulty < 0.33 {
		return "Easy"
	}
	if GlobalConfig.Difficulty >= 0.33 && GlobalConfig.Difficulty < 0.66 {
		return "Medium"
	}
	return "Hard"
}

const configFilePath = "./config/settings.json" // Name of the configuration file

// SaveConfig saves the configuration to a JSON file.
func SaveConfig(config *Config) error {
	// Converts *config struct in JSON formatted with indentation ("" is prefix and "  " is indentation)
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	// Create the file if it doesn't exist, or replace it if it does
	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close() // Close the file when the function returns
	// Write the JSON data to the file
	_, err = file.Write(data)
	// If all went well err is nil
	return err
}

// LoadConfig loads the configuration from a JSON file.
func LoadConfig(filePath string) (*Config, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close() // Close the file when the function returns
	// Create a new Config struct that will hold the loaded data
	var config Config
	// Decode the JSON data into the Config struct
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	// Returns a pointer to the loaded Config struct
	return &config, nil
}

// UpdateConfig updates the configuration using a provided function and saves it to the file.
func UpdateConfig(updateFunc func(*Config)) error {
	// Call the update function to modify the configuration
	updateFunc(GlobalConfig)
	// Save the updated configuration to the file
	return SaveConfig(GlobalConfig)
}

// InitConfig initializes the configuration by loading it from a file or using the default configuration.
func InitConfig() {
	// Check if the configuration file exists
	config, err := LoadConfig(configFilePath)
	// If the file doesn't exist or there's an error loading it, use the default configuration
	if err != nil {
		// Create the file with the default configuration
		_ = SaveConfig(DefaultConfig)
		GlobalConfig = DefaultConfig
	} else {
		GlobalConfig = config
	}
}
