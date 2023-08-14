package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

type InstallationStep struct {
	Description string
	Action      func() error
	Progress    float64
}

func installZshAndOhMyZsh(m model, reportProgress func(float64)) (string, error) {
	steps := []InstallationStep{
		{"Installing Zsh", installZsh, 20},
		{"Installing Oh My Zsh", installOhMyZsh, 40},
		{"Cloning setupjunkie-dotfiles", cloneSetupJunkieDotfiles, 60},
		{"Overwriting .zshrc", overwriteZshrc, 80},
		{"Installing Zsh plugins", func() error {
			return installZshPlugins("zsh-users/zsh-autosuggestions", "zsh-users/zsh-syntax-highlighting")
		}, 100},
	}

	progressChan := make(chan float64)
	errorChan := make(chan error)

	go func() {
		for _, step := range steps {
			err := step.Action()
			if err != nil {
				errorChan <- err
				return
			}
			progressChan <- step.Progress
		}
		close(progressChan)
	}()

	for progress := range progressChan {
		reportProgress(progress)
	}

	select {
	case err := <-errorChan:
		return "", err
	default:
		return "Zsh, Oh My Zsh, and .zshrc setup completed.", nil
	}
}

func installZsh() error {
	cmd := exec.Command("sudo", "apt-get", "install", "-y", "zsh")
	return cmd.Run()
}

func installOhMyZsh() error {
	cmd := exec.Command("sh", "-c", `RUNZSH=no sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"`)
	return cmd.Run()
}

func cloneSetupJunkieDotfiles() error {
	tmpDir := "/tmp/setupjunkie-dotfiles"
	cmd := exec.Command("git", "clone", "https://github.com/andremore/setupjunkie-dotfiles.git", tmpDir)
	return cmd.Run()
}

func overwriteZshrc() error {
	tmpDir := "/tmp/setupjunkie-dotfiles"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user's home directory: %w", err)
	}
	cmd := exec.Command("sudo", "cp", tmpDir+"/.zshrc", homeDir+"/.zshrc")
	return cmd.Run()
}

func installZshPlugins(plugins ...string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user's home directory: %w", err)
	}

	for _, plugin := range plugins {
		url := fmt.Sprintf("https://github.com/%s.git", plugin)
		destination := fmt.Sprintf("%s/.oh-my-zsh/custom/plugins/%s", homeDir, strings.Split(plugin, "/")[1])

		cmd := exec.Command("git", "clone", url, destination)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install plugin %s: %w", plugin, err)
		}
	}

	return nil
}
