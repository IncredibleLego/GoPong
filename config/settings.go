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
	Difficulty             float64
	TextDimension          float64
	ScreenWidth            int
	ScreenHeight           int
	MenuOptionsPerSecond   time.Duration
	OptionsPerSecond       time.Duration
	MaxBounceAngle         float64
}

var GlobalConfig = &Config{
	Scale:                  1.0,
	Fullscreen:             true,
	Player1Name:            "Player 1",
	Player2Name:            "Player 2",
	BallSpeed:              6,
	BallSize:               15,
	PaddleSpeed:            6,
	PaddleHeight:           100,
	PaddleDistanceFromWall: 40,
	Difficulty:             0.5,
	TextDimension:          20,
	ScreenWidth:            960,
	ScreenHeight:           720,
	MenuOptionsPerSecond:   4,
	OptionsPerSecond:       90,
	MaxBounceAngle:         0.7853975, //45.0 * (3.14159 / 180.0)
}

var DefaultConfig = &Config{
	Scale:                  1.0,
	Fullscreen:             true,
	Player1Name:            "Player 1",
	Player2Name:            "Player 2",
	BallSpeed:              6,
	BallSize:               15,
	PaddleSpeed:            6,
	PaddleHeight:           100,
	PaddleDistanceFromWall: 40,
	Difficulty:             0.5,
	TextDimension:          20,
	ScreenWidth:            960,
	ScreenHeight:           720,
	MenuOptionsPerSecond:   4,
	OptionsPerSecond:       90,
	MaxBounceAngle:         0.7853975, //45.0 * (3.14159 / 180.0)
}

func ApplyScaling(cfg *Config) {
	cfg.BallSpeed = ScaledInt(cfg.BallSpeed)
	cfg.BallSize = ScaledInt(cfg.BallSize)
	cfg.PaddleSpeed = ScaledInt(cfg.PaddleSpeed)
	cfg.PaddleHeight = ScaledInt(cfg.PaddleHeight)
	cfg.PaddleDistanceFromWall = ScaledInt(cfg.PaddleDistanceFromWall)
	cfg.ScreenWidth = ScaledInt(cfg.ScreenWidth)
	cfg.ScreenHeight = ScaledInt(cfg.ScreenHeight)
}

func BaseScale(cfg *Config) {
	cfg.BallSpeed = DefaultConfig.BallSpeed
	cfg.BallSize = DefaultConfig.BallSize
	cfg.PaddleSpeed = DefaultConfig.PaddleSpeed
	cfg.PaddleHeight = DefaultConfig.PaddleHeight
	cfg.PaddleDistanceFromWall = DefaultConfig.PaddleDistanceFromWall
	cfg.ScreenWidth = DefaultConfig.ScreenWidth
	cfg.ScreenHeight = DefaultConfig.ScreenHeight
}

// ScaledInt and ScaledFloat are utility functions to scale values based on the global configuration.
// They are used to adjust the size of game elements based on the current scale factor.

func ScaledInt(value int) int {
	return int(math.Round(float64(value) * GlobalConfig.Scale))
}

func ScaledFloat(value float64) float64 {
	return value * GlobalConfig.Scale
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
