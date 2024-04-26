package gapBuffer

func (g *GapBuffer) RemoveLeftChar() {
	g.gapStart -= 1
}
