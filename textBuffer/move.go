package textBuffer

func (t *TextBuffer) MoveLeft() {
	if t.col == 0 {
		if t.line == 0 {
			return
		}

		t.line -= 1
		t.col = t.lineLengths[t.line]
		t.cursorOffset -= 1

		return
	}

	t.col -= 1
	t.cursorOffset -= 1
}

func (t *TextBuffer) MoveRight() {
	if t.col == t.lineLengths[t.line] {
		if t.line == len(t.lineLengths)-1 {
			return
		}

		t.line += 1
		t.col = 0
		t.cursorOffset += 1

		return
	}

	t.col += 1
	t.cursorOffset += 1
}

func (t *TextBuffer) MoveUp() {
	if t.line == 0 {
		return
	}

	t.line -= 1
	previousLineLength := t.lineLengths[t.line]

	if (previousLineLength) < t.col {
		t.cursorOffset -= t.col + 1
		t.col = previousLineLength

		return
	}

	t.cursorOffset -= t.col + 1 + (previousLineLength - t.col)
}

func (t *TextBuffer) MoveDown() {
	if t.line == len(t.lineLengths)-1 {
		return
	}

	t.line += 1

	nextLineLength := t.lineLengths[t.line]
	currentCol := t.col

	if nextLineLength < t.col {
		t.col = nextLineLength
	}

	t.cursorOffset += (t.lineLengths[t.line-1] - currentCol) + 1 + t.col
}

func (t *TextBuffer) MoveToEndOfLine() {
	lineLength := t.lineLengths[t.line]

	if t.col == lineLength {
		return
	}

	t.cursorOffset += lineLength - t.col
	t.col = lineLength
}

func (t *TextBuffer) MoveToStartOfLine() {
	if t.col == 0 {
		return
	}

	t.cursorOffset -= t.col
	t.col = 0
}
