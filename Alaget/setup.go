package Alaget

import (
	"os/exec"
	"strings"
)

func FindAlacrittyApp() (string, error) {
	// Execute the command to find the Alacritty.app directory
	cmd := exec.Command("bash", "-c", "find $(brew --prefix) -name 'Alacritty.app'")

	// Get the command output
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Trim any trailing newline from the output
	appPath := strings.TrimSpace(string(output))

	return appPath, nil
}

func RemoveQuarantine(appPath string) error {
	// Execute the command to remove the quarantine attribute
	cmd := exec.Command("xattr", "-d", "com.apple.quarantine", appPath)

	// Run the command
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
