package main

import (
	"time"
	"context"
	"os/exec"

)

type timeOutStep struct{
	step
	timeOut time.Duration
}


func NewTimeoutStep(name, exe, message, projDir string, args []string, timeout time.Duration) *timeOutStep{
	if timeout == 0{
		timeout = time.Second *30
	}
	return &timeOutStep{
		step: *NewStep(
			name,
			exe,
			message,
			projDir,
			args,
		),
		timeOut: timeout,
	}
}


func (t *timeOutStep)execute()(string, error){
	ctx, cancel := context.WithTimeout(context.Background(), t.timeOut)
	defer cancel()
	cmd := exec.CommandContext(ctx, t.exe, t.args...)
	if err:= cmd.Run(); err!= nil{
		if ctx.Err() == context.DeadlineExceeded{
			return "", &stepErr{
				step: t.name,
				msg: "failed to execute",
				cause: err,
			}
		}
		return "", &stepErr{
			step: t.name,
			msg: "failed to execute",
			cause: err,
		}
	}

	return t.message, nil

}