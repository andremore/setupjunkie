package main

import "strings"

type model struct {
	cursor    int
	selected  map[int]struct{}
	submitted bool
}

func initialModel() model {
	return model{
		selected: make(map[int]struct{}),
	}
}

func processSelectedChoices(m model, ch []Choice) (string, error) {
	messages := []string{}

	for index := range m.selected {
		if ch[index].Name == "Submit" {
			continue
		}

		output, err := ch[index].Action(m)
		if err != nil {
			return "", err
		}

		messages = append(messages, output)
		if ch[index].PostMessage != "" {
			messages = append(messages, ch[index].PostMessage)
		}
	}

	return strings.Join(messages, "\n"), nil
}

func submitSelectedChoices(m model) (string, error) {
	m.submitted = true
	return processSelectedChoices(m, choices)
}
