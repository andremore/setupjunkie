package main


func moveCursorUp(m model) model {
	if m.Cursor > 0 {
		m.Cursor--
	}
	return m
}

func moveCursorDown(m model) model {
	if m.Cursor < len(choices)-1 {
		m.Cursor++
	}
	return m
}
