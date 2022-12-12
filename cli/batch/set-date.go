package batch

import (
	"github.com/chzyer/readline"
	"github.com/go-errors/errors"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"set-date",
			readline.PcItemDynamic(completionCalendarList),
		),
	)

	commands["set-date"] = func(args []string) error {
		var err error

		if len(args) != 5 {
			return errors.New("usage: set-date <calendar> <format> <date> <source> <target>")
		}

		calendar, format, date, source, target := args[0], args[1], args[2], args[3], args[4]

		if stamp, ok := state[source]; ok {
			state[target], err = stamp.SetDate(date, calendar, format)
			if err != nil {
				return err
			}

			printf("%s = %s\n", target, state[target])
		} else {
			return errors.New("timestamp '" + source + "' not set")
		}

		return nil
	}
}
