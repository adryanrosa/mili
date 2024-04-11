package gapBuffer

const (
	defaultCapacity = 1024
	growFactor      = 2
)

type GapBuffer struct {
	data []byte

	gapStart int
	gapEnd   int
}
