package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if !m.submitted {
				for index := range m.selected {
					output, err := choices[index].Action(m, func(progress float64) {
						choices[index].CurrentProgress = progress
					})
					if err != nil {
						fmt.Println("Error:", err)
						return m, tea.Quit
					}
					fmt.Println(output)
				}
			}
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; !ok {
				m.selected[m.cursor] = struct{}{}
			} else {
				delete(m.selected, m.cursor)
			}
		case "s":
			output, err := submitSelectedChoices(m)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println(output)
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What do you want to set up?\n\n"
	for i, choice := range choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s", cursor, checked, choice.Name)
		if choice.CurrentProgress > 0 {
			progress := progressBar(choice.CurrentProgress, 20)
			s += fmt.Sprintf(" %s", progress)
		}

		s += "\n"
	}
	s += "\nPress q to quit. Press s to submit.\n"

	return s
}

func main() {
	if err := ensureGitAndCurlInstalled(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
