package sculptor

import "sync"

var CommonChannel chan any
var CommonWg *sync.WaitGroup

// Merge merges multiple DataSculptor output in one channel
// and return the common channel.
func Merge(sculptors ...*DataSculptor) chan any {
	if len(sculptors) == 0 {
		return nil
	}
	commChanSize := 0
	for _, v := range sculptors {
		commChanSize += v.options.BufSize
	}
	CommonChan := make(chan any, commChanSize)
	for i, _ := range sculptors {
		sculptors[i].ConstructedOutput = CommonChan
	}
	return CommonChan
}

// MergeAndDo merges multiple DataSculptor output in one channel and automatically call Do() for each DataSculptor
// and close the channel when all DataSculptor are done
func MergeAndDo(sculptors ...*DataSculptor) chan any {
	CommonChan := Merge(sculptors...)
	for _, v := range sculptors {
		v.Do()
	}
	go func() {
		for _, v := range sculptors {
			v.Wait() // All Closed
		}
		close(CommonChan)
	}()
	return CommonChan
}

// MergeAndDoWithWg merges multiple DataSculptor output in one channel and automatically call Do() for each DataSculptor
// and close the channel when all DataSculptor are done
func MergeAndDoWithWg(wg *sync.WaitGroup, sculptors ...*DataSculptor) chan any {
	commChan := Merge(sculptors...)
	for _, v := range sculptors {
		v.Wg = wg
		v.ConstructedOutput = commChan
		v.Do()
	}
	go func() {
		wg.Wait()
		close(commChan)
	}()
	return commChan
}

// MergeV2 merges multiple DataSculptor.ConstructedOutput in one channel and DataSculptor.Wg in one *sync.WaitGroup
// and return the common wg and chan
func MergeV2(sculptors ...*DataSculptor) (*sync.WaitGroup, chan any) {
	if len(sculptors) == 0 {
		return nil, nil
	}
	commChanSize := 0
	commWg := &sync.WaitGroup{} // Get first Wg
	for _, v := range sculptors {
		commChanSize += v.options.BufSize
	}
	commChan := make(chan any, commChanSize)
	for i, _ := range sculptors {
		sculptors[i].Wg = commWg
		sculptors[i].ConstructedOutput = commChan
	}
	return commWg, commChan
}

// MergeV3 merges multiple DataSculptor outputChannel in one channel and sync.WaitGroup in one wg
// And registry the CommonChannel and CommonWg to package scope global variable
// you can run AutoCloseV3() func to close CommonChannel when all DataSculptor are done
func MergeV3(sculptor ...*DataSculptor) chan any {
	CommonWg, CommonChannel = MergeV2(sculptor...)
	return CommonChannel
}

// AutoCloseV3 is a helper func to close (package global variable) CommonChannel when all DataSculptor are done
func AutoCloseV3() {
	go func() {
		CommonWg.Wait()
		close(CommonChannel)
	}()
}
