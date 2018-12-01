package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/danhunsaker/calends"
	"github.com/urfave/cli"
)

func init() {
	commands = append(commands, []cli.Command{
		{
			Name:      "format",
			Usage:     "format a timestamp in a given calendar system and format",
			ArgsUsage: "<to-calendar> <to-format> [<stamp>]",

			Action: func(c *cli.Context) error {
				if c.NArg() < 2 || c.NArg() > 3 {
					return errArgMismatch
				}

				outCal := c.Args()[0]
				outForm := c.Args()[1]
				stamp := ""
				if c.NArg() == 3 {
					stamp = c.Args()[2]
				} else {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						stamp = fmt.Sprintf("%s%s", stamp, scanner.Text())
					}
				}

				var moment calends.Calends
				err := json.Unmarshal([]byte(stamp), &moment)
				if err != nil {
					return cli.NewExitError(err, 2)
				}

				output, err := moment.Date(outCal, outForm)
				if err != nil {
					return cli.NewExitError(err, 2)
				}

				fmt.Println(output)
				return nil
			},
		},
	}...)
}
