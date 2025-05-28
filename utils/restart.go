package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// RestartGame reloads the current process with the same arguments.
func RestartGame() {
	exe, err := os.Executable()
	if err != nil {
		log.Printf("Error in finding the executable: %v", err)
		os.Exit(1)
	}
	// Warning if the executable is in a temporary directory
	if strings.Contains(exe, os.TempDir()) {
		log.Printf("Restart not supported with 'go run'. Compile first with 'go build'")
		os.Exit(1)
	}
	args := os.Args
	env := os.Environ()
	cmd := exec.Command(exe, args[1:]...)
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Start the new process
	err = cmd.Start()
	if err != nil {
		log.Printf("Error in starting the new process: %v", err)
		os.Exit(1)
	}

	// Attend briefly to ensure the new process starts before exiting the old one
	time.Sleep(300 * time.Millisecond)

	os.Exit(0)
}
