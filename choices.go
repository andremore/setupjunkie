package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Choice struct {
	Name        string
	Action      func(m model) (string, error)
	PostMessage string
}

var zsh = ZshInstaller{}

var choices = []Choice{
	{
		Name:        "Install Zsh (and Oh My Zsh)",
		Action:      func(m model) (string, error) { return zsh.Install() },
		PostMessage: "To make zsh your default shell, run : chsh -s $(which zsh)",
	},
}

func toggleChoice(m model) model {
	if _, ok := m.Selected[m.Cursor]; !ok {
		m.Selected[m.Cursor] = struct{}{}
	} else {
		delete(m.Selected, m.Cursor)
	}
	return m
}

func handleQuit(m model) (tea.Model, tea.Cmd) {
	if !m.Submitted {
		for index, choice := range choices {
			if _, selected := m.Selected[index]; selected {
				output, err := choice.Action(m)
				if err != nil {
					fmt.Println("Error:", err)
					return m, tea.Quit
				}
				fmt.Println(output)
				if choice.PostMessage != "" {
					fmt.Println(choice.PostMessage)
				}
			}
		}
	}

	return m, tea.Quit
}

func handleSubmit(m model) (tea.Model, tea.Cmd) {
	m.Submitted = true
	for index := range m.Selected {
		m.InstallQueue = append(m.InstallQueue, index)
	}
	return m, m.startInstalling(choices)
}