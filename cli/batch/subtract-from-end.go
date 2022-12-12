package batch

import (
	"github.com/chzyer/readline"
	"github.com/go-errors/errors"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"subtract-from-end",
			readline.PcItemDynamic(completionCalendarList),
		),
	)

	commands["subtract-from-end"] = func(args []string) error {
		var err error

		if len(args) != 4 {
			return errors.New("usage: subtract-from-end <calendar> <offset> <source> <target>")
		}

		calendar, offset, source, target := args[0], args[1], args[2], args[3]

		if stamp, ok := state[source]; ok {
			state[target], err = stamp.SubtractFromEnd(offset, calendar)
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
