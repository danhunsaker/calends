package batch

import (
	"github.com/chzyer/readline"
	"github.com/go-errors/errors"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"compare",
			readline.PcItemDynamic(
				completionStatesList,
				readline.PcItemDynamic(
					completionStatesList,
					readline.PcItemDynamic(func(arg string) []string {
						return []string{"start", "end", "start-end", "end-start", "duration"}
					}),
				),
			),
		),
	)

	commands["compare"] = func(args []string) error {
		if len(args) != 3 {
			return errors.New("usage: compare <source> <compare> <mode>")
		}

		source, compare, mode := args[0], args[1], args[2]

		if stamp1, ok := state[source]; ok {
			if stamp2, ok := state[compare]; ok {
				printf("%d\n", stamp1.Compare(stamp2, mode))
			} else {
				return errors.New("timestamp '" + compare + "' not set")
			}
		} else {
			return errors.New("timestamp '" + source + "' not set")
		}

		return nil
	}
}
