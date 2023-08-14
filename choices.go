package main

type Choice struct {
	Name   string
	Action func(m model) (string, error)
}

var choices = []Choice{
	{
		Name:   "Install Zsh (and Oh My Zsh)",
		Action: installZshAndOhMyZsh,
	},
}
