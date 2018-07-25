package batch

import (
	"errors"

	"github.com/chzyer/readline"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"ends-during",
			readline.PcItemDynamic(
				completionStatesList,
				readline.PcItemDynamic(completionStatesList),
			),
		),
	)

	commands["ends-during"] = func(args []string) error {
		if len(args) != 2 {
			return errors.New("usage: ends-during <source> <compare>")
		}

		source, compare := args[0], args[1]

		if stamp1, ok := state[source]; ok {
			if stamp2, ok := state[compare]; ok {
				printf("%t\n", stamp1.EndsDuring(stamp2))
			} else {
				return errors.New("timestamp '" + compare + "' not set")
			}
		} else {
			return errors.New("timestamp '" + source + "' not set")
		}

		return nil
	}
}
