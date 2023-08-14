package main

import (
	"fmt"
	"os/exec"
)

func isCommandAvailable(name string) bool {
	cmd := exec.Command("command", "-v", name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func ensureGitAndCurlInstalled() error {
	if !isCommandAvailable("git") {
		cmd := exec.Command("sudo", "apt-get", "install", "-y", "git")
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install git: %w", err)
		}
	}

	if !isCommandAvailable("curl") {
		cmd := exec.Command("sudo", "apt-get", "install", "-y", "curl")
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install curl: %w", err)
		}
	}

	return nil
}

func installZshAndOhMyZsh(m model) (string, error) {
	// Install Zsh
	cmd := exec.Command("sudo", "apt-get", "install", "-y", "zsh")
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to install zsh: %w", err)
	}

	// Install Oh My Zsh without changing the shell immediately
	cmd = exec.Command("sh", "-c", `RUNZSH=no sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"`)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to install Oh My Zsh: %w", err)
	}

	return "Zsh and Oh My Zsh installed successfully.", nil
}
