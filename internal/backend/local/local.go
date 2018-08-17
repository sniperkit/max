/*
Sniperkit-Bot
- Status: analyzed
*/

package local

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/sniperkit/snk.fork.max/internal/backend"
	"github.com/sniperkit/snk.fork.max/internal/backend/config"
	"github.com/sniperkit/snk.fork.max/internal/exec"
	"github.com/sniperkit/snk.fork.max/internal/task"
)

type engine struct {
	config *config.Backend
}

// New creates a new local engine.
func New(config *config.Backend) backend.Engine {
	return &engine{config}
}

// Name returns engine name.
func (e *engine) Name() string {
	return "local"
}

// Setup setups local engine.
func (e *engine) Setup(ctx context.Context, t *task.Task) error {
	return nil
}

// Exec executes a task.
func (e *engine) Exec(ctx context.Context, t *task.Task) error {
	for _, c := range t.Commands.Values {
		if e.config.Verbose {
			log.Print(fmt.Sprintf("$ %s", c))
		}

		opts := &exec.Options{
			Context: ctx,
			Dir:     t.Dir,
			Env:     toEnv(t.Variables),
			Command: c,
			Stdin:   e.config.Stdin,
			Stdout:  e.config.Stdout,
			Stderr:  e.config.Stderr,
		}

		// Execute command.
		if err := exec.Exec(opts); err != nil {
			log.Print(c)
			return err
		}
	}

	return nil
}

// Logs returns logs from the local engine.
func (e *engine) Logs(ctx context.Context, t *task.Task) (io.ReadCloser, error) {
	return nil, nil
}

// Destroy destroys the local engine.
func (e *engine) Destroy(ctx context.Context, t *task.Task) error {
	return nil
}

// Wait check if the local engine is done or not.
func (e *engine) Wait(ctx context.Context, t *task.Task) (bool, error) {
	return true, nil
}
