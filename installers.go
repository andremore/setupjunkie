package main

import (
	"fmt"
	"os/exec"
)

type Installer interface {
	Install() (string, error)
	InstallSteps() []InstallationStep
}

type InstallationStep struct {
	Description string
	Action      func() error
}

func isCommandAvailable(name string) bool {
	cmd := exec.Command("command", "-v", name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func ensureGitAndCurlInstalled() error {
	return ensureCommandsInstalled([]string{"git", "curl"})
}

func ensureCommandsInstalled(commands []string) error {
	for _, cmd := range commands {
		if !isCommandAvailable(cmd) {
			installCmd := exec.Command("sudo", "apt-get", "install", "-y", cmd)
			if err := installCmd.Run(); err != nil {
				return fmt.Errorf("failed to install %s: %w", cmd, err)
			}
		}
	}
	return nil
}
