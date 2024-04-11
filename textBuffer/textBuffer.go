package textBuffer

import "github.com/adryanrosa/mili/gapBuffer"

type TextBuffer struct {
	gapBuffer gapBuffer.GapBuffer

	col  int
	line int

	shiftCursorBy int
	LineLengths   []int
}

func New() TextBuffer {
	return TextBuffer{
		gapBuffer: gapBuffer.New(),

		col:  0,
		line: 0,

		shiftCursorBy: 0,
		LineLengths:   []int{0},
	}
}

func (t TextBuffer) String() string {
	return t.gapBuffer.String(false)
}

func (t TextBuffer) CursorPosition() (int, int) {
	return t.col, t.line
}
