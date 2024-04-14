package main

import (
	"fmt"
	"os"

	"github.com/adryanrosa/mili/textBuffer"
	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textBuffer textBuffer.TextBuffer
	cursor     cursor.Model
}

func initialModel() model {
	cursor := cursor.New()
	cursor.Focus()

	return model{
		textBuffer: textBuffer.New(),
		cursor:     cursor,
	}
}

func (m model) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m model) View() string {
	preCursor, cursor, postCursor := m.textBuffer.Text()
	m.cursor.SetChar(cursor)

	return preCursor + m.cursor.View() + postCursor
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch messageType := message.(type) {

	case tea.KeyMsg:
		switch messageType.Type {

		case tea.KeyCtrlC:
			return m, tea.Quit

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
			m.textBuffer.InsertRune(messageType.Runes[0])
		}
	}

	return m, nil
}

func main() {
	program := tea.NewProgram(initialModel())

	if _, err := program.Run(); err != nil {
		fmt.Printf("Error: %v", err)

		os.Exit(1)
	}
}
