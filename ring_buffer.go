package sugar

type RingBuffer struct {
	inputChannel  <-chan any
	outputChannel chan any
}

func NewRingBuffer(inputChannel <-chan any, outputChannel chan any) *RingBuffer {
	return &RingBuffer{inputChannel, outputChannel}
}

func (r *RingBuffer) Run() {
	for v := range r.inputChannel {
		select {
		case r.outputChannel <- v:
		default:
			<-r.outputChannel
			r.outputChannel <- v
		}
	}
	close(r.outputChannel)
}
