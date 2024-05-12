package Dependency

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func CheckOS() {
	if runtime.GOOS != "darwin" {
		fmt.Println("Alawow is only supported on macOS. Exiting...")
		os.Exit(1)
	}
}

func CheckDep(dependency string) bool {
	_, err := exec.LookPath(dependency)
	return err == nil
}

func InstallHomebrew() {
	cmd := exec.Command("/bin/bash", "-c", `$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)`)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Could not install homebrew. Exiting...")
		os.Exit(1)
	}
	fmt.Print("\nHomebrew downloaded successfully ✅\n\n")

	message, err := addHomebrewToPath()
	fmt.Print(message)
	if err != nil {
		os.Exit(1)
	}
}

func addHomebrewToPath() (string, error) {
	file, err := os.OpenFile("~/.zprofile", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "Could not open ~/.zprofile. Exiting...", err
	}
	defer file.Close()

	// Write the commands to the file
	_, err = file.WriteString("\n" + "eval \"$(/opt/homebrew/bin/brew shellenv)\"\n")
	if err != nil {
		return "Could not edit ~/.zprofile. Exiting...", err
	}

	// Execute the shell command to update the shell environment
	cmd := exec.Command("bash", "-c", "eval \"$(/opt/homebrew/bin/brew shellenv)\"")
	if err := cmd.Run(); err != nil {
		return "Could not source ~/.zprofile. Exiting...", err
	}

	return "Homebrew added to path ✅", nil
}

func InstallGit() {
	cmd := exec.Command("brew", "install", "git")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Could not install git. Exiting...")
		os.Exit(1)
	} else {
		fmt.Print("Git installed successfully ✅\n\n")
	}
}
