package batch

import (
	"errors"

	"github.com/chzyer/readline"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"merge",
			readline.PcItemDynamic(
				completionStatesList,
				readline.PcItemDynamic(completionStatesList),
			),
		),
	)

	commands["merge"] = func(args []string) error {
		var err error

		if len(args) != 3 {
			return errors.New("usage: merge <source> <combine> <target>")
		}

		source, combine, target := args[0], args[1], args[2]

		if stamp1, ok := state[source]; ok {
			if stamp2, ok := state[combine]; ok {
				state[target], err = stamp1.Merge(stamp2)
				if err != nil {
					return err
				}

				printf("%s = %s\n", target, state[target])
			} else {
				return errors.New("timestamp '" + combine + "' not set")
			}
		} else {
			return errors.New("timestamp '" + source + "' not set")
		}

		return nil
	}
}
