package main

import (
	"fmt"
	"os"
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

	// Clone the setupjunkie-dotfiles repo
	tmpDir := "/tmp/setupjunkie-dotfiles"
	cmd = exec.Command("git", "clone", "https://github.com/andremore/setupjunkie-dotfiles.git", tmpDir)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to clone setupjunkie-dotfiles: %w", err)
	}

    // Overwrite the .zshrc file
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return "", fmt.Errorf("failed to get user's home directory: %w", err)
    }
    cmd = exec.Command("sudo", "cp", tmpDir+"/.zshrc", homeDir+"/.zshrc")
    if err := cmd.Run(); err != nil {
        return "", fmt.Errorf("failed to overwrite .zshrc: %w", err)
    }

	// Install zsh-autosuggestions
    cmd = exec.Command("git", "clone", "https://github.com/zsh-users/zsh-autosuggestions", fmt.Sprintf("%s/.oh-my-zsh/custom/plugins/zsh-autosuggestions", homeDir))
    if err := cmd.Run(); err != nil {
        return "", fmt.Errorf("failed to install zsh-autosuggestions: %w", err)
    }

    // Install zsh-syntax-highlighting
    cmd = exec.Command("git", "clone", "https://github.com/zsh-users/zsh-syntax-highlighting.git", fmt.Sprintf("%s/.oh-my-zsh/custom/plugins/zsh-syntax-highlighting", homeDir))
    if err := cmd.Run(); err != nil {
        return "", fmt.Errorf("failed to install zsh-syntax-highlighting: %w", err)
    }

	return "Zsh, Oh My Zsh, and .zshrc setup completed.", nil
}
