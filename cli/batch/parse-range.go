package batch

import (
	"errors"

	"github.com/chzyer/readline"
	"github.com/danhunsaker/calends"
)

func init() {
	completions = append(
		completions,
		readline.PcItem(
			"parse-range",
			readline.PcItemDynamic(completionCalendarList),
		),
	)

	commands["parse-range"] = func(args []string) error {
		var err error

		if len(args) != 5 {
			return errors.New("usage: parse-range <calendar> <format> <date> <end-date> <target>")
		}

		calendar, format, date, end, target := args[0], args[1], args[2], args[3], args[4]

		state[target], err = calends.Create(map[string]interface{}{"start": date, "end": end}, calendar, format)
		if err != nil {
			return err
		}

		printf("%s = %s\n", target, state[target])

		return nil
	}
}
