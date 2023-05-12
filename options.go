package sculptor

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// Options is used for extra Control during DataSculptor.Do()
type Options struct {
	Latency    time.Duration
	BufSize    int
	TagKey     string
	DebugLevel log.Level
}

// DefaultOptions is the default options used by DataSculptor
var DefaultOptions = Options{
	Latency:    1 * time.Millisecond,
	BufSize:    10,
	TagKey:     "select",
	DebugLevel: log.WarnLevel,
}

func (d *DataSculptor) debugInfo() {
	log.SetLevel(d.options.DebugLevel)
}

// SetOption set options to replace default options
func (d *DataSculptor) SetOption(o Options) *DataSculptor {
	d.options = o
	return d
}

// SetBufSize will reset the buffer size of the channel
func (d *DataSculptor) SetBufSize(size int) {
	d.options.BufSize = size
	d.ConstructedOutput = make(chan any, size)
}
