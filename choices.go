package main

type Choice struct {
	Name            string
	Action          func(m model, progressReporter func(float64)) (string, error)
	PostMessage     string
	CurrentProgress float64
}

var choices = []Choice{
	{
		Name:        "Install Zsh (and Oh My Zsh)",
		Action:      installZshAndOhMyZsh,
		PostMessage: "To make zsh your default shell, run : chsh -s $(which zsh)",
	},
}
