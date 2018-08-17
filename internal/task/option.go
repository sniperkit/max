/*
Sniperkit-Bot
- Status: analyzed
*/

package task

import (
	"log"
)

// Option configures a runtime option.
type Option func(*Task)

// Args returns an option configured with a args value.
func Args(args map[string]interface{}) Option {
	return func(t *Task) {
		if t.Args == nil {
			t.Args = make(map[string]interface{})
		}

		for k, v := range args {
			t.Args[k] = v
		}
	}
}

// Log returns an option configured with a log value.
func Log(log *log.Logger) Option {
	return func(t *Task) {
		t.log = log
	}
}

// Variables returns an option configured with a variables value.
func Variables(vars map[string]string) Option {
	return func(t *Task) {
		if t.Variables == nil {
			t.Variables = make(map[string]string)
		}

		for k, v := range vars {
			t.Variables[k] = v
		}
	}
}
