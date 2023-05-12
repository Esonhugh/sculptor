package sculptor

// Merge merges multiple DataSculptor output in one channel
func Merge(sculptors ...*DataSculptor) chan any {
	if len(sculptors) == 0 {
		return nil
	}
	commChanSize := 0
	for _, v := range sculptors {
		commChanSize += v.options.BufSize
	}
	commChan := make(chan any, commChanSize)
	for i, _ := range sculptors {
		sculptors[i].ConstructedOutput = commChan
	}
	return commChan
}
