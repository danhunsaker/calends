// Package cli/batch creates the batch operations interface to the calends
// library.
package batch

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/chzyer/readline"
	"github.com/danhunsaker/calends"
	"github.com/danhunsaker/calends/calendars"
	"github.com/mattn/go-shellwords"
	"github.com/urfave/cli/v2"
)

// Collections
var completions []readline.PrefixCompleterInterface
var commands = make(map[string]func([]string) error)
var state = make(map[string]calends.Calends)

// Helper functions
var printf func(string, ...interface{}) // Need the readline session to properly define this one
var completionCalendarList = func(arg string) []string {
	return calendars.ListRegistered()
}
var completionStatesList = func(arg string) (list []string) {
	for name := range state {
		list = append(list, name)
	}
	return
}

// Console contains the main logic for batch mode
var Console = func(c *cli.Context) error {
	completions = append(
		completions,
		readline.PcItem("help"),
		readline.PcItem("exit"),
		readline.PcItem("quit"),
	)
	completer := readline.NewPrefixCompleter(completions...)
	shellwords.ParseEnv = true
	shellwords.ParseBacktick = true

	l, err := readline.NewEx(&readline.Config{
		Prompt:            "\033[35mcalends \033[36m» \033[0m",
		HistoryFile:       "",
		AutoComplete:      completer,
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()

	// Normally this helper would be defined entirely outside the main logic
	// function, but we need the readline session, first, to make it work...
	printf = func(message string, args ...interface{}) {
		io.WriteString(l.Stdout(), fmt.Sprintf(message, args...))
	}

	log.SetOutput(l.Stderr())
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)

		if line == "" || line == "help" {
			printf("commands:\n")
			printf(completer.Tree("    "))
			continue
		} else if line == "quit" || line == "exit" {
			break
		}

		args, err := shellwords.Parse(line)
		if err != nil {
			log.Printf("%v\n", err)
		} else if cmd, ok := commands[args[0]]; ok {
			err = cmd(args[1:])
			if err != nil {
				log.Printf("%v\n", err)
			}
		} else {
			log.Printf("unknown command %q\n", args[0])
		}
	}
	return nil
}
