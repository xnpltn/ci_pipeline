package main

import (
	"os/exec"
)

type step struct {
	name    string
	exe     string // exe, string representing the executable name of the external tool we want to execute,
	args    []string
	message string
	projDir string
}

func NewStep(name, exe, message, projDir string, args []string) *step {
	return &step{
		name:    name,
		exe:     exe,
		args:    args,
		message: message,
		projDir: projDir,
	}
}

func (s *step) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)
	cmd.Dir = s.projDir
	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: err,
		}
	}

	return s.message, nil
}
