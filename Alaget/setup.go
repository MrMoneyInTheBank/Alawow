package Alaget

import (
	"os"
	"os/exec"
	"path/filepath"
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

func InstallNerdFont() error {
	cmd := exec.Command("brew", "tap", "homebrew/cask-fonts")
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("brew", "install", "font-bigblue-terminal-nerd-font")

	err = cmd.Run()
	if err != nil {
		return err
	}
	return err
}

func GenerateAlacrittyConfig() error {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Construct the destination file path
	destDirPath := filepath.Join(homeDir, ".config", "alacritty")
	destFilePath := filepath.Join(destDirPath, "alacritty.toml")

	// Create the destination directory if it doesn't exist
	err = os.MkdirAll(destDirPath, 0755)
	if err != nil {
		return err
	}

	// Define the contents of the alacritty.toml file
	config := `# Default Alacritty configuration

[font]
normal.family = "BigBlueTerm437 Nerd Font"
size = 20
`

	// Create the destination file for writing
	destFile, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Write the contents to the destination file
	_, err = destFile.WriteString(config)
	if err != nil {
		return err
	}

	return nil
}
