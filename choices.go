package main

type Choice struct {
	Name   string
	Action func(m model) (string, error)
	PostMessage string
}

var choices = []Choice{
	{
		Name:   "Install Zsh (and Oh My Zsh)",
		Action: installZshAndOhMyZsh,
		PostMessage: "To make zsh your default shell, run : chsh -s $(which zsh)",
	},
}
