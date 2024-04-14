package textBuffer

func (t *TextBuffer) MoveLeft() {
	if t.col == 0 {
		if t.line == 0 {
			return
		}

		t.line -= 1
		t.col = t.lineLengths[t.line]
		t.shiftCursorBy -= 1

		return
	}

	t.col -= 1
	t.shiftCursorBy -= 1
}

func (t *TextBuffer) MoveRight() {
	if t.col == t.lineLengths[t.line] {
		if t.line == len(t.lineLengths)-1 {
			return
		}

		t.line += 1
		t.col = 0
		t.shiftCursorBy += 1

		return
	}

	t.col += 1
	t.shiftCursorBy += 1
}

func (t *TextBuffer) MoveUp() {
	if t.line == 0 {
		return
	}

	t.line -= 1
	previousLineLength := t.lineLengths[t.line]

	if (previousLineLength) < t.col {
		t.shiftCursorBy -= t.col + 1
		t.col = previousLineLength

		return
	}

	t.shiftCursorBy -= t.col + 1 + (previousLineLength - t.col)
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

	t.shiftCursorBy += (t.lineLengths[t.line-1] - currentCol) + 1 + t.col
}

func (t *TextBuffer) MoveToEndOfLine() {
	lineLength := t.lineLengths[t.line]

	if t.col == lineLength {
		return
	}

	t.shiftCursorBy += lineLength - t.col
	t.col = lineLength
}

func (t *TextBuffer) MoveToStartOfLine() {
	if t.col == 0 {
		return
	}

	t.shiftCursorBy -= t.col
	t.col = 0
}
