package sculptor

import "time"

// Options is used for extra Control during DataSculptor.Do()
type Options struct {
	Latency time.Duration
	TagKey  string
}

// SetOption set options to replace default options
func (d *DataSculptor) SetOption(o Options) {
	d.options = o
}
