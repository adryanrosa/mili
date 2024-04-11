package textBuffer

func (t *TextBuffer) adjustGapForCursorPosition() {
	if t.shiftCursorBy == 0 {
		return
	}

	t.gapBuffer.ShiftGapBy(t.shiftCursorBy)
	t.shiftCursorBy = 0
}

func (t *TextBuffer) insertNewLine() {
	previousCol := t.col
	previousLineLength := t.lineLengths[t.line]

	t.col = 0
	t.line += 1

	split := previousLineLength - previousCol

	t.lineLengths = append(t.lineLengths, 0)
	copy(t.lineLengths[t.line+1:], t.lineLengths[t.line:])
	t.lineLengths[t.line] = split

	if split > 0 {
		t.lineLengths[t.line-1] = previousLineLength - split
	}
}

func (t *TextBuffer) InsertRune(char rune) {
	t.adjustGapForCursorPosition()
	t.gapBuffer.InsertRune(char)

	if char == '\n' {
		t.insertNewLine()

		return
	}

	t.col += 1
	t.lineLengths[t.line] += 1
}
