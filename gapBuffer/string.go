package gapBuffer

import "strings"

func (g GapBuffer) String(withGap bool) string {
	var builder strings.Builder

	if withGap {
		builder.Grow(len(g.data))
		builder.Write(g.data)

		return builder.String()
	}

	builder.Grow(len(g.data) - (g.gapEnd - g.gapStart))
	builder.Write(g.data[:g.gapStart])
	builder.Write(g.data[g.gapEnd:])

	return builder.String()
}
