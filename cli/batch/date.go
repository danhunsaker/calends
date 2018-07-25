package batch

import (
	"errors"

	"github.com/chzyer/readline"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"date",
			readline.PcItemDynamic(completionCalendarList),
		),
	)

	commands["date"] = func(args []string) error {
		if len(args) != 3 {
			return errors.New("usage: date <calendar> <format> <source>")
		}

		calendar, format, source := args[0], args[1], args[2]

		if stamp, ok := state[source]; ok {
			date, err := stamp.Date(calendar, format)
			if err != nil {
				return err
			}

			printf("%s\n", date)
		} else {
			return errors.New("timestamp '" + source + "' not set")
		}

		return nil
	}
}
