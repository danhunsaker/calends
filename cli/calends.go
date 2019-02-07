// Package cli creates the command-line interface to the calends library.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/danhunsaker/calends"
	"github.com/urfave/cli"
)

var commands []cli.Command

var (
	errArgMismatch = cli.NewExitError("Incorrect argument count\n", 1)
)

func main() {
	app := cli.NewApp()

	// metadata
	app.Name = "calends"
	app.Usage = "manipulate dates and times in multiple calendar systems at once"
	app.Version = calends.Version
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		{
			Name:  "Dan Hunsaker",
			Email: "danhunsaker@gmail.com",
		},
	}
	app.Copyright = "(c) 2018 Dan Hunsaker"

	// configuration
	app.EnableBashCompletion = true
	app.Commands = commands
	app.ExitErrHandler = handleExitError
	app.Action = defaultAction

	// Make it go!!!!
	app.Run(os.Args)
}

// replicated and adapted from cli package...
func handleExitError(context *cli.Context, err error) {
	if err == nil {
		return
	}

	if exitErr, ok := err.(cli.ExitCoder); ok {
		if err.Error() != "" {
			if _, ok := exitErr.(cli.ErrorFormatter); ok {
				fmt.Fprintf(cli.ErrWriter, "%+v\n", err)
			} else {
				fmt.Fprintln(cli.ErrWriter, err)
			}
		}
		if exitErr.ExitCode() == 1 {
			cli.ShowCommandHelp(context, context.Command.Name)
		}
		cli.OsExiter(exitErr.ExitCode())
		return
	}

	if multiErr, ok := err.(cli.MultiError); ok {
		code, isUsageError := handleMultiError(multiErr)
		if isUsageError {
			cli.ShowCommandHelp(context, context.Command.Name)
		}
		cli.OsExiter(code)
		return
	}
}

func handleMultiError(multiErr cli.MultiError) (int, bool) {
	code := 127
	isUsageError := false
	for _, merr := range multiErr.Errors {
		if multiErr2, ok := merr.(cli.MultiError); ok {
			code, isUsageError = handleMultiError(multiErr2)
		} else {
			fmt.Fprintln(cli.ErrWriter, merr)
			if exitErr, ok := merr.(cli.ExitCoder); ok {
				if exitErr.ExitCode() == 1 {
					isUsageError = true
				}
				code = exitErr.ExitCode()
			}
		}
	}
	return code, isUsageError
}
