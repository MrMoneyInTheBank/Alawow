package main

import (
	"fmt"
	"os"

	ala "Alawow/Alaget"
	dep "Alawow/Deps"
)

func main() {
	dep.CheckOS()
	fmt.Fprintln(os.Stdout, []any{"OS Check passed ✅\n"}...)

	brewInstalled := dep.CheckDep("brew")
	if brewInstalled {
		fmt.Print("Homebrew executable in PATH ✅\n")
	} else {
		dep.InstallHomebrew()
	}

	gitIntstalled := dep.CheckDep("brew")
	if gitIntstalled {
		fmt.Print("Git executable in PATH ✅\n")
	} else {
		dep.InstallGit()
	}

	alacrittyInstalled := dep.CheckDep("alacritty")
	if alacrittyInstalled {
		fmt.Print("Alacritty executable in PATH ✅\n")
	} else {
		ala.InstallAlacritty()
	}
}
