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
			Name:      "offset",
			Usage:     "adjusts a timestamp by an offset in a given calendar",
			ArgsUsage: "<offset-calendar> [<offset> [<stamp>]]",

			Action: func(c *cli.Context) error {
				if c.NArg() < 1 || c.NArg() > 3 {
					return errArgMismatch
				}

				offCal := c.Args().Get(0)
				offset := ""
				if c.NArg() > 1 {
					offset = c.Args().Get(1)
				} else {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						offset = fmt.Sprintf("%s%s", offset, scanner.Text())
					}
				}
				stamp := ""
				if c.NArg() == 3 {
					stamp = c.Args().Get(2)
				} else {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						stamp = fmt.Sprintf("%s%s", stamp, scanner.Text())
					}
				}

				var moment calends.Calends
				err := json.Unmarshal([]byte(stamp), &moment)
				if err != nil {
					return cli.Exit(err, 2)
				}

				moment, err = moment.Add(offset, offCal)
				if err != nil {
					return cli.Exit(err, 2)
				}

				moment, err = moment.SetDuration("0", "tai64")
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
