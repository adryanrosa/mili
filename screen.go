package main

import (
	"github.com/gdamore/tcell/v2"
)

func cleanup(screen tcell.Screen) {
	maybePanic := recover()

	screen.Fini()

	if maybePanic != nil {
		panic(maybePanic)
	}
}

func renderText(screen tcell.Screen, text string) {
	line := 0
	column := 0

	for _, char := range text {
		// FOR DEBUGGING BUFFER GAP
		if char == 0 {
			char = '_'
		}

		if char == '\n' {
			line += 1
			column = 0

			continue
		}

		if char == '\t' {
			column += 2

			continue
		}

		screen.SetContent(column, line, char, nil, tcell.StyleDefault)
		column += 1
	}
}
