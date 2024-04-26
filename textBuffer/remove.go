package textBuffer

func (t *TextBuffer) RemoveLeftChar() {
	if t.col == 0 && t.line == 0 {
		return
	}

	t.adjustGapForCursorPosition()
	t.gapBuffer.RemoveLeftChar()

	if t.col == 0 {
		previousLineLength := t.lineLengths[t.line-1]

		t.lineLengths[t.line-1] += t.lineLengths[t.line]
		t.lineLengths = append(t.lineLengths[:t.line], t.lineLengths[t.line+1:]...)

		t.line -= 1
		t.col = previousLineLength

		return
	}

	t.col -= 1
	t.lineLengths[t.line] -= 1
}
