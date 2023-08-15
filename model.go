package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Cursor       int
	Selected     map[int]struct{}
	Submitted    bool
	InstallQueue []int
	IsInstalling bool
}

func initialModel() model {
	return model{
		Selected: make(map[int]struct{}),
	}
}

func (m *model) startInstalling(ch []Choice) tea.Cmd {
	if len(m.InstallQueue) == 0 {
		m.IsInstalling = false
		return tea.Quit
	}

	choiceIndex := m.InstallQueue[0]
	m.IsInstalling = true
	choice := ch[choiceIndex]

	return func() tea.Msg {
		output, err := choice.Action(*m)
		if err != nil {
			return err
		}

		if choice.PostMessage != "" {
			fmt.Println(choice.PostMessage)
		}

		m.InstallQueue = m.InstallQueue[1:]
		if len(m.InstallQueue) > 0 {
			return m.startInstalling(ch)
		}

		return output
	}
}
