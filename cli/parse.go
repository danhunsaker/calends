package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/danhunsaker/calends"
	"github.com/urfave/cli/v2"
)

func init() {
	commands = append(commands, []*cli.Command{
		{
			Name:      "parse",
			Usage:     "parse a date/time given a calendar system and format",
			ArgsUsage: "<from-calendar> <from-format> [<date>]",

			Action: func(c *cli.Context) error {
				if c.NArg() < 2 || c.NArg() > 3 {
					return errArgMismatch
				}

				inCal := c.Args().Get(0)
				inForm := c.Args().Get(1)
				date := ""
				if c.NArg() == 3 {
					date = c.Args().Get(2)
				} else {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						date = fmt.Sprintf("%s%s", date, scanner.Text())
					}
				}

				moment, err := calends.Create(date, inCal, inForm)
				if err != nil {
					return cli.Exit(err, 2)
				}

				output, err := json.Marshal(moment)
				if err != nil {
					return cli.Exit(err, 2)
				}

				fmt.Println(string(output))
				return nil
			},
		},
	}...)
}
