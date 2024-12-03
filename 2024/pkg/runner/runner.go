package runner

import "bufio"

type Runner interface {
	Run(*bufio.Reader) (int, int)
}

func NoOpRunner() Runner {
	return &noOpRunner{}
}

type noOpRunner struct{}

func (r *noOpRunner) Run(readFn *bufio.Reader) (int, int) {
	return 0, 0
}
