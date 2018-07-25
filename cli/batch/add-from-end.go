package batch

import (
	"errors"

	"github.com/chzyer/readline"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"add-from-end",
			readline.PcItemDynamic(completionCalendarList),
		),
	)

	commands["add-from-end"] = func(args []string) error {
		var err error

		if len(args) != 4 {
			return errors.New("usage: add-from-end <calendar> <offset> <source> <target>")
		}

		calendar, offset, source, target := args[0], args[1], args[2], args[3]

		if stamp, ok := state[source]; ok {
			state[target], err = stamp.AddFromEnd(offset, calendar)
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
