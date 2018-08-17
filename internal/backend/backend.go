/*
Sniperkit-Bot
- Status: analyzed
*/

package backend

import (
	"context"
	"io"

	"github.com/sniperkit/snk.fork.max/internal/task"
)

// Engine defines a engine that can run tasks.
type Engine interface {
	Destroy(context.Context, *task.Task) error
	Exec(context.Context, *task.Task) error
	Logs(context.Context, *task.Task) (io.ReadCloser, error)
	Name() string
	Setup(context.Context, *task.Task) error
	Wait(context.Context, *task.Task) (bool, error)
}
