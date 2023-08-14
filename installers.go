package main

import (
	"bytes"
	"os/exec"
)

// func installZshAndOhMyZsh() error {
// // Install Zsh
// // cmd := exec.Command("sh", "-c", "sudo apt-get install zsh")
// cmd := exec.Command("sh", "-c", "zsh --version")
// if err := cmd.Run(); err != nil {
// return err
// }

// // Install Oh My Zsh
// // cmd = exec.Command("sh", "-c", "sh -c \"$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)\"")
// return cmd.Run()
// }

// ...

func installZshAndOhMyZsh(m model) (string, error) {
	cmd := exec.Command("sh", "-c", "zsh --version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}
