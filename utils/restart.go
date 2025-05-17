package utils

import (
	"os"
	"os/exec"
)

// RestartGame reloads the current process with the same arguments.
func RestartGame() {
	exe, err := os.Executable()
	if err != nil {
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
		os.Exit(1)
	}
	// Terminate the current process
	os.Exit(0)
}
