package batch

import (
	"github.com/chzyer/readline"
	"github.com/go-errors/errors"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"difference",
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

	commands["difference"] = func(args []string) error {
		if len(args) != 3 {
			return errors.New("usage: difference <source> <compare> <mode>")
		}

		source, compare, mode := args[0], args[1], args[2]

		if stamp1, ok := state[source]; ok {
			if stamp2, ok := state[compare]; ok {
				diff := stamp1.Difference(stamp2, mode)
				printf("%s\n", diff.Text('f', -6))
			} else {
				return errors.New("timestamp '" + compare + "' not set")
			}
		} else {
			return errors.New("timestamp '" + source + "' not set")
		}

		return nil
	}
}
