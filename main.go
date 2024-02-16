package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type executer interface {
	execute() (string, error)
}

func run(projDir string, out io.Writer) error {
	if projDir == "" {
		return fmt.Errorf("Project directory is required: %w", ErrorValidation)
	}
	pipeline := make([]executer, 4)

	pipeline[0] = NewStep(
		"go build",
		"go",
		"Go Build: SUCCESS",
		projDir,
		[]string{"build", ".", "errors"},
	)

	pipeline[1] = NewStep(
		"go test",
		"go",
		"Go Test: SUCCESS",
		projDir,
		[]string{"test", "-v"},
	)

	pipeline[2] = NewExceptionStep(
		"go fmt",
		"gofmt",
		"Gofmt: SUCCESS",
		projDir,
		[]string{"-l", "."},
	)

	pipeline[3] = NewTimeoutStep(
		"Git Push",
		"git",
		"Git Push: SUCCESS",
		projDir,
		[]string{"push", "origin", "master"},
		time.Second*30,
	)

	for _, s := range pipeline {
		message, err := s.execute()
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(out, message)
		if err != nil {
			return err
		}
	}
	return nil

}

func main() {
	projDir := flag.String("p", "", "FIle Directory")
	flag.Parse()

	if err := run(*projDir, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
