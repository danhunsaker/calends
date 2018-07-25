package batch

import (
	"errors"

	"github.com/chzyer/readline"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"set-end-date",
			readline.PcItemDynamic(completionCalendarList),
		),
	)

	commands["set-end-date"] = func(args []string) error {
		var err error

		if len(args) != 5 {
			return errors.New("usage: set-end-date <calendar> <format> <date> <source> <target>")
		}

		calendar, format, date, source, target := args[0], args[1], args[2], args[3], args[4]

		if stamp, ok := state[source]; ok {
			state[target], err = stamp.SetEndDate(date, calendar, format)
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
