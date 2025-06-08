package main

import (
	"math/rand"
)

type CircuitBreaker struct {
	// The maximum number of failures before the circuit opens
	threshold int
	// The current state of the circuit
	state    string
	failures int
}

func (c *CircuitBreaker) open() {
	println("circuit_breaker_opened")
	c.state = "open"
	c.threshold = 3
}

func (c *CircuitBreaker) close() {
	println("circuit_breaker_closed")
	c.state = "closed"
}

type simulator struct {
	pointer int
	flow    []int
}

func (s *simulator) new(flow []int) {
	s.pointer = 0
	s.flow = flow
}

func (s *simulator) run() int {
	isEnd := s.pointer >= len(s.flow)
	if isEnd {
		println("-simulator_end")
		return -1
	}
	randv := rand.Intn(2)
	if s.flow[s.pointer] == randv {
		s.pointer++
		println("-simulator_success", s.pointer)
		return 0
	} else {
		println("-simulator_fail", s.pointer)
		return 1
	}
}

func main() {
	var circuitBreaker CircuitBreaker = CircuitBreaker{
		threshold: 3,
		state:     "closed",
		failures:  0,
	}
	// generate a random flow of 0s and 1s dynamically
	f := make([]int, 10)
	for i := 0; i < len(f); i++ {
		f[i] = rand.Intn(2)
	}
	var s simulator = simulator{
		pointer: 0,
		flow:    f,
	}
	s.new(s.flow)
	for {
		v := s.run()
		if v == -1 {
			break
		} else if v == 1 {
			circuitBreaker.failures++
			if circuitBreaker.failures >= circuitBreaker.threshold {
				circuitBreaker.open()
				break
			}
		} else {
			circuitBreaker.failures = 0
			circuitBreaker.close()
		}
	}
}
