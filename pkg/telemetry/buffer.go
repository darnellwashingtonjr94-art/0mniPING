package telemetry

type ResultBuffer struct {
	size   int
	buffer []Result
	head   int
	count  int
}

func NewResultBuffer(size int) *ResultBuffer {
	return &ResultBuffer{
		size:   size,
		buffer: make([]Result, size),
	}
}

func (rb *ResultBuffer) Push(res Result) {
	rb.buffer[rb.head] = res
	rb.head = (rb.head + 1) % rb.size
	if rb.count < rb.size {
		rb.count++
	}
}

func (rb *ResultBuffer) GetAll() []Result {
	results := make([]Result, rb.count)
	for i := 0; i < rb.count; i++ {
		idx := (rb.head - rb.count + i + rb.size) % rb.size
		results[i] = rb.buffer[idx]
	}
	return results
}
