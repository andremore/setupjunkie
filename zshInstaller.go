package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type ZshInstaller struct{}

func (z ZshInstaller) Install() (string, error) {
	steps := z.InstallSteps()
	for _, step := range steps {
		err := step.Action()
		if err != nil {
			return "", err
		}
	}
	return "Zsh, Oh My Zsh, and .zshrc setup completed.", nil
}

func (z ZshInstaller) InstallSteps() []InstallationStep {
	return []InstallationStep{
		{"Installing Zsh", installZsh},
		{"Installing Oh My Zsh", installOhMyZsh},
		{"Cloning setupjunkie-dotfiles", cloneSetupJunkieDotfiles},
		{"Overwriting .zshrc", overwriteZshrc},
		{"Installing Zsh plugins", func() error {
			return installZshPlugins("zsh-users/zsh-autosuggestions", "zsh-users/zsh-syntax-highlighting")
		}},
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
