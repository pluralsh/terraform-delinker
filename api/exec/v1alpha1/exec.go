package v1alpha1

import (
	"os"
	"os/exec"
)

func WithDir(workingDirectory string) ExecutableOption {
	return func(t *Executable) {
		t.workingDirectory = workingDirectory
	}
}

func (this *Executable) Run(args ...string) (output []byte, err error) {
	cmd := exec.Command(string(this.command), args...)
	cmd.Stderr = os.Stderr

	if len(this.workingDirectory) > 0 {
		cmd.Dir = this.workingDirectory
	}

	output, err = cmd.Output()
	return
}

func NewExecutable(command Command, options ...ExecutableOption) *Executable {
	exec := &Executable{command: command}

	for _, o := range options {
		o(exec)
	}

	return exec
}
