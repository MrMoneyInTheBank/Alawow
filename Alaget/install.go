package Alaget

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func progressDots(done chan struct{}, pack string) {
	defer close(done)
	fmt.Printf("Downloading %s ", pack)
	for {
		select {
		case <-done:
			fmt.Println()
			return
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Print(".")
		}
	}
}

func InstallAlacritty() {
	cmd := exec.Command("brew", "install", "--cask", "alacritty")

	done := make(chan struct{})
	go progressDots(done, "Alacritty")

	err := cmd.Start()
	if err != nil {
		fmt.Println("Could not install Alacritty")
		os.Exit(1)
	}

	err = cmd.Wait()
	done <- struct{}{}

	if err != nil {
		fmt.Println("Could not install Alacritty")
		os.Exit(1)
	} else {
		fmt.Println("\nAlacritty installed successfully ✅")
	}
}
