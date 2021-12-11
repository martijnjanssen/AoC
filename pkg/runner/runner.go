package runner

type Runner interface {
	Run() (int, int)
}

func NoOpRunner() Runner {
	return &noOpRunner{}
}

type noOpRunner struct{}

func (r *noOpRunner) Run() (int, int) {
	return 0, 0
}
