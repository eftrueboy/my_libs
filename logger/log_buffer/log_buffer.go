package log_buffer

type LogBuffer struct {
	buffer chan string
}

func NewLogBuffer(length int) *LogBuffer {
	return &LogBuffer{buffer: make(chan string, length)}
}

func (this *LogBuffer) Write(text string) {
	this.buffer <- text
}

func (this *LogBuffer) Read() (text string) {
	text = <-this.buffer
	return
}
