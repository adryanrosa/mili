package textBuffer

import (
	"github.com/adryanrosa/mili/gapBuffer"
)

type TextBuffer struct {
	gapBuffer gapBuffer.GapBuffer

	col  int
	line int

	cursorOffset int
	lineLengths  []int
}

func New() TextBuffer {
	return TextBuffer{
		gapBuffer: gapBuffer.New(),

		col:  0,
		line: 0,

		cursorOffset: 0,
		lineLengths:  []int{0},
	}
}
