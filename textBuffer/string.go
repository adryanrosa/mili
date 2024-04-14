package textBuffer

func (t TextBuffer) textBeforeAndAfterCursor() (beforeCursor string, afterCursor string) {
	beforeGap, afterGap := t.gapBuffer.Text()

	if t.cursorOffset < 0 {
		split := len(beforeGap) + t.cursorOffset

		return beforeGap[:split], beforeGap[split:] + afterGap
	}

	if t.cursorOffset > 0 {
		return beforeGap + afterGap[:t.cursorOffset], afterGap[t.cursorOffset:]
	}

	return beforeGap, afterGap
}

func (t TextBuffer) Text() (beforeCursor string, cursor string, afterCursor string) {
	beforeCursor, afterCursor = t.textBeforeAndAfterCursor()
	emptyCursor := " "

	if len(afterCursor) == 0 {
		return beforeCursor, emptyCursor, afterCursor
	}

	cursor = afterCursor[:1]

	if cursor == "\n" {
		return beforeCursor, emptyCursor, afterCursor
	}

	return beforeCursor, cursor, afterCursor[1:]
}
