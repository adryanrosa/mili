package gapBuffer

func newFromCapacity(capacity int) GapBuffer {
	return GapBuffer{
		data: make([]byte, capacity),

		gapStart: 0,
		gapEnd:   capacity,
	}
}

func New() GapBuffer {
	return newFromCapacity(defaultCapacity)
}

func NewFromText(text string) GapBuffer {
	if text == "" {
		return newFromCapacity(defaultCapacity)
	}

	size := max(defaultCapacity, len(text)*growFactor)
	buffer := make([]byte, size)
	occupiedBytes := copy(buffer, text)

	return GapBuffer{
		data: buffer,

		gapStart: occupiedBytes,
		gapEnd:   size,
	}
}
