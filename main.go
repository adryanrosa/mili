package main

import (
	"fmt"
	"os"

	"github.com/adryanrosa/mili/textBuffer"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textBuffer textBuffer.TextBuffer
	ready      bool

	cursor   cursor.Model
	viewport viewport.Model
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
	return nil
}

func (m model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	return m.viewport.View()
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {

	case tea.WindowSizeMsg:
		if !m.ready {
			m.viewport = viewport.New(message.Width, message.Height)
			m.ready = true
		} else {
			m.viewport.Width = message.Width
			m.viewport.Height = message.Height
		}

	case tea.KeyMsg:
		keyCommand := m.handleKeyMessage(message)

		if keyCommand != nil {
			return m, keyCommand
		}
	}

	preCursor, cursor, postCursor := m.textBuffer.Text()
	m.cursor.SetChar(cursor)
	m.viewport.SetContent(preCursor + m.cursor.View() + postCursor)

	var viewportCommand tea.Cmd
	m.viewport, viewportCommand = m.viewport.Update(message)

	return m, viewportCommand
}

func main() {
	program := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := program.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)

		os.Exit(1)
	}
}
