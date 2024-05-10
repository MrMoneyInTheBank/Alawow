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
	} else {
		fmt.Print("Homebrew installed successfully ✅\n\n")
	}
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
