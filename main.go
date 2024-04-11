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
			{
				switch event.Key() {

				case tcell.KeyEscape, tcell.KeyCtrlC:
					return

				case tcell.KeyLeft:
					textBuffer.MoveLeft()

				case tcell.KeyRight:
					textBuffer.MoveRight()

				case tcell.KeyUp:
					textBuffer.MoveUp()

				case tcell.KeyDown:
					textBuffer.MoveDown()

				case tcell.KeyEnter:
					textBuffer.InsertRune('\n')

				case tcell.KeyRune:
					textBuffer.InsertRune(event.Rune())

				case tcell.KeyBackspace, tcell.KeyBackspace2:
					//

				case tcell.KeyDelete:
					//
				}
			}
		}

		screen.ShowCursor(textBuffer.CursorPosition())
		renderText(screen, textBuffer.String())
	}
}
