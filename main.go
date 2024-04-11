package main

import (
	"log"

	"github.com/adryanrosa/mili/textBuffer"
	"github.com/gdamore/tcell/v2"
)

func main() {
	textBuffer := textBuffer.New()
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defer cleanup(screen)

	for {
		screen.Show()
		screen.Clear()

		switch event := screen.PollEvent().(type) {

		case *tcell.EventResize:
			screen.Sync()

		case *tcell.EventKey:
			if event.Key() == tcell.KeyCtrlC {
				return
			}

			keyHandler(*event, &textBuffer)
		}

		screen.ShowCursor(textBuffer.CursorPosition())
		renderText(screen, textBuffer.String())
	}
}
