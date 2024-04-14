package gapBuffer

import "strings"

func (g GapBuffer) string() string {
	var builder strings.Builder

	builder.Grow(len(g.data) - (g.gapEnd - g.gapStart))
	builder.Write(g.data[:g.gapStart])
	builder.Write(g.data[g.gapEnd:])

	return builder.String()
}

func (g GapBuffer) Text() (beforeGap string, afterGap string) {
	text := g.string()

	return text[:g.gapStart], text[g.gapStart:]
}
