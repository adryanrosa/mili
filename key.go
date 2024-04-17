package main

import tea "github.com/charmbracelet/bubbletea"

func (m *model) handleKeyMessage(message tea.KeyMsg) tea.Cmd {
	switch message.Type {

	case tea.KeyCtrlC:
		return tea.Quit

	case tea.KeyLeft:
		m.textBuffer.MoveLeft()

	case tea.KeyRight:
		m.textBuffer.MoveRight()

	case tea.KeyUp:
		m.textBuffer.MoveUp()

	case tea.KeyDown:
		m.textBuffer.MoveDown()

	case tea.KeyHome:
		m.textBuffer.MoveToStartOfLine()

	case tea.KeyEnd:
		m.textBuffer.MoveToEndOfLine()

	case tea.KeyEnter:
		m.textBuffer.InsertRune('\n')

	case tea.KeyRunes, tea.KeySpace:
		m.textBuffer.InsertRune(message.Runes[0])
	}

	return nil
}
