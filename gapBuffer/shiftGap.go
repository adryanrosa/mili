package gapBuffer

func (g *GapBuffer) ShiftGapBy(difference int) {
	newPosition := g.gapStart + difference

	gapLength := g.gapEnd - g.gapStart
	newPosition = min(newPosition, len(g.data)-gapLength)

	if g.gapStart == newPosition {
		return
	}

	if newPosition < g.gapStart {
		delta := g.gapStart - newPosition

		copy(g.data[g.gapEnd-delta:], g.data[g.gapStart-delta:g.gapStart])
		g.gapStart -= delta
		g.gapEnd -= delta

		return
	}

	delta := newPosition - g.gapStart

	copy(g.data[g.gapStart:], g.data[g.gapEnd:g.gapEnd+delta])
	g.gapStart += delta
	g.gapEnd += delta
}
