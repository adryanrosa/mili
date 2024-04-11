package gapBuffer

import "unicode/utf8"

func (g *GapBuffer) growByFactor() {
	grownBuffer := make([]byte, len(g.data)*growFactor)
	copy(grownBuffer, g.data[:g.gapStart])

	end := len(grownBuffer) - (len(g.data) - g.gapEnd)
	copy(grownBuffer[end:], g.data[g.gapEnd:])

	g.data = grownBuffer
	g.gapEnd = end
}

func (g *GapBuffer) ensureGapCapacity(bytesLength int) {
	if g.gapEnd-g.gapStart <= bytesLength {
		g.growByFactor()
	}
}

func (g *GapBuffer) InsertRune(char rune) {
	bytesLength := utf8.RuneLen(char)
	g.ensureGapCapacity(bytesLength)

	utf8.EncodeRune(g.data[g.gapStart:], char)
	g.gapStart += bytesLength
}

func (g *GapBuffer) InsertText(text string) {
	bytesLength := len(text)
	g.ensureGapCapacity(bytesLength)

	copy(g.data[g.gapStart:], text)
	g.gapStart += bytesLength
}
