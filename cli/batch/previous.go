package batch

import (
	"errors"

	"github.com/chzyer/readline"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"previous",
			readline.PcItemDynamic(completionCalendarList),
		),
	)

	commands["previous"] = func(args []string) error {
		var err error

		if len(args) != 4 {
			return errors.New("usage: previous <calendar> <offset> <source> <target>")
		}

		calendar, offset, source, target := args[0], args[1], args[2], args[3]

		if stamp, ok := state[source]; ok {
			state[target], err = stamp.Previous(offset, calendar)
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
