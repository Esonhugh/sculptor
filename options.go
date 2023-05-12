package sculptor

import "time"

// Options is used for extra Control during DataSculptor.Do()
type Options struct {
	Latency time.Duration
	BufSize int
	TagKey  string
}

// DefaultOptions is the default options used by DataSculptor
var DefaultOptions = Options{
	Latency: 1 * time.Millisecond,
	BufSize: 10,
	TagKey:  "select",
}

// SetOption set options to replace default options
func (d *DataSculptor) SetOption(o Options) {
	d.options = o
}

// SetBufSize will reset the buffer size of the channel
func (d *DataSculptor) SetBufSize(size int) {
	d.options.BufSize = size
	d.ConstructedOutput = make(chan any, size)
}
