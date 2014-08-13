/*Package gochan implements most cases of channel patterns in golang
We not only provides the usage of that pattern, as well as the benchmark result, which could be helpful in implementation.
*/
package gochan

// StopP defines the stop pattern with channel
// for each pattern stop pattern implementation, we use a goroutine to keep increment the counter unitl it has reached maxCounter
type StopP struct {
	counter    int64
	maxCounter int64
	stopChan   chan bool
}

// NewStopP allocates and returns an instance of StopP
func NewStopP(c, max int64) *StopP {
	return &StopP{
		counter:    c,
		maxCounter: max,
		stopChan:   make(chan bool),
	}
}

// Stop stop the StopP
func (s *StopP) Stop() {
	s.stopChan <- true
}

// NonBlockingStop implements the stop pattern with nonblocking channel mode
func (s *StopP) NonBlockingStop() {
	for {
		select {
		case <-s.stopChan:
			// watch the stop signal, if it is, return immediately
			return
		default:
			// otherwise, keep incr 1
			s.counter = s.counter + 1
			if s.counter > s.maxCounter {
				return
			}
		}
	}
}

// BlockingStop implements the stop pattern with blocking channel mode
func (s *StopP) BlockingStop() {
	stop := false
	go func() {
		// watch the stop signal, if it is, turn on stop flag
		<-s.stopChan
		stop = true
	}()

	for {
		if stop {
			return
		}
		// otherwise, keep incr 1
		s.counter = s.counter + 1
		if s.counter > s.maxCounter {
			return
		}
	}
}
