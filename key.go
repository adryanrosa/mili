package main

import (
	"github.com/adryanrosa/mili/textBuffer"
	"github.com/gdamore/tcell/v2"
)

func keyHandler(event tcell.EventKey, textBuffer *textBuffer.TextBuffer) {
	switch event.Key() {

	case tcell.KeyLeft:
		textBuffer.MoveLeft()

	case tcell.KeyRight:
		textBuffer.MoveRight()

	case tcell.KeyUp:
		textBuffer.MoveUp()

	case tcell.KeyDown:
		textBuffer.MoveDown()

	case tcell.KeyHome:
		textBuffer.MoveToStartOfLine()

	case tcell.KeyEnd:
		textBuffer.MoveToEndOfLine()

	case tcell.KeyEnter:
		textBuffer.InsertRune('\n')

	case tcell.KeyRune:
		textBuffer.InsertRune(event.Rune())

	case tcell.KeyBackspace, tcell.KeyBackspace2:
		break

	case tcell.KeyDelete:
		break
	}
}
