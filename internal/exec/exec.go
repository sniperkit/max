/*
Sniperkit-Bot
- Status: analyzed
*/

package exec

import (
	"context"
	"io"
	"os"
	"strings"

	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
)

// Options represents execute options.
type Options struct {
	Context context.Context
	Dir     string
	Env     []string
	Command string
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
}

// Exec will execute a input cmd string.
func Exec(opts *Options) error {
	var path string

	if len(opts.Dir) > 0 {
		path = opts.Dir
	}

	if len(path) == 0 {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		path = wd
	}

	p, err := syntax.NewParser().Parse(strings.NewReader(opts.Command), "")
	if err != nil {
		return err
	}

	env := os.Environ()
	env = append(env, opts.Env...)
	envi, err := interp.EnvFromList(env)
	if err != nil {
		return err
	}

	r := interp.Runner{
		Context: opts.Context,
		Env:     envi,
		Dir:     path,
		Exec:    interp.DefaultExec,
		Open:    interp.OpenDevImpls(interp.DefaultOpen),
		Stdin:   opts.Stdin,
		Stdout:  opts.Stdout,
		Stderr:  opts.Stderr,
	}

	if err = r.Reset(); err != nil {
		return err
	}

	return r.Run(p)
}
