package executor

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/yohamta/dagu/internal/config"
)

type Executor interface {
	SetStdout(out io.Writer)
	SetStderr(out io.Writer)
	Kill(sig os.Signal) error
	Run() error
}

type Creator func(ctx context.Context, step *config.Step) (Executor, error)

var executors = make(map[string]Creator)

func Register(name string, register Creator) {
	executors[name] = register
}

func CreateExecutor(ctx context.Context, step *config.Step) (Executor, error) {
	f, ok := executors[step.Executor]
	if ok {
		return f(ctx, step)
	}
	return nil, fmt.Errorf("invalid executor: %s", step.Executor)
}