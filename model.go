package main

import "fmt"

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
	for index := range m.selected {
		if ch[index].Name == "Submit" {
			continue
		}

		output, err := ch[index].Action(m)
		if err != nil {
			return "", err
		}
		fmt.Println(output)
	}
	return "All actions executed successfully.", nil
}

func submitSelectedChoices(m model) (string, error) {
	m.submitted = true
	return processSelectedChoices(m, choices)
}
