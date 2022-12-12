package batch

import (
	"github.com/chzyer/readline"
	"github.com/danhunsaker/calends"
	"github.com/go-errors/errors"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"parse",
			readline.PcItemDynamic(completionCalendarList),
		),
	)

	commands["parse"] = func(args []string) error {
		var err error

		if len(args) != 4 {
			return errors.New("usage: parse <calendar> <format> <date> <target>")
		}

		calendar, format, date, target := args[0], args[1], args[2], args[3]

		state[target], err = calends.Create(date, calendar, format)
		if err != nil {
			return err
		}

		printf("%s = %s\n", target, state[target])

		return nil
	}
}
