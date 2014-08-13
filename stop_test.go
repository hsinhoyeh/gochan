package gochan

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestStopPNonBlocking(t *testing.T) {
	s := NewStopP(0, math.MaxInt64)
	go s.NonBlockingStop()
	time.Sleep(10 * time.Millisecond)
	s.Stop()
	fmt.Printf("count in nonblocking mode: %d\n", s.counter)
}

func TestStopPBlocking(t *testing.T) {
	s := NewStopP(0, math.MaxInt64)
	go s.BlockingStop()
	time.Sleep(10 * time.Millisecond)
	s.Stop()
	fmt.Printf("count in blocking mode: %d\n", s.counter)
}

func BenchmarkStopPNonBlocking(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewStopP(0, int64(math.MaxInt16)).NonBlockingStop()
	}
}

func BenchmarkStopPBlocking(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewStopP(0, int64(math.MaxInt16)).BlockingStop()
	}
}
