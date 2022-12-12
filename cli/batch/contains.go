package batch

import (
	"github.com/chzyer/readline"
	"github.com/go-errors/errors"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"contains",
			readline.PcItemDynamic(
				completionStatesList,
				readline.PcItemDynamic(completionStatesList),
			),
		),
	)

	commands["contains"] = func(args []string) error {
		if len(args) != 2 {
			return errors.New("usage: contains <source> <compare>")
		}

		source, compare := args[0], args[1]

		if stamp1, ok := state[source]; ok {
			if stamp2, ok := state[compare]; ok {
				printf("%t\n", stamp1.Contains(stamp2))
			} else {
				return errors.New("timestamp '" + compare + "' not set")
			}
		} else {
			return errors.New("timestamp '" + source + "' not set")
		}

		return nil
	}
}
