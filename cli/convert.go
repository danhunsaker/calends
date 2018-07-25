package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/danhunsaker/calends"
	"github.com/urfave/cli"
)

func init() {
	commands = append(commands, []cli.Command{
		cli.Command{
			Name:      "convert",
			Usage:     "convert a date/time from one calendar and format to another",
			ArgsUsage: "<from-calendar> <from-format> <to-calendar> <to-format> [<date>]",

			Action: func(c *cli.Context) error {
				if c.NArg() < 4 || c.NArg() > 5 {
					return errArgMismatch
				}

				in_cal := c.Args()[0]
				in_form := c.Args()[1]
				out_cal := c.Args()[2]
				out_form := c.Args()[3]
				date := ""
				if c.NArg() == 5 {
					date = c.Args()[4]
				} else {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						date = fmt.Sprintf("%s%s", date, scanner.Text())
					}
				}

				moment, err := calends.Create(date, in_cal, in_form)
				if err != nil {
					return cli.NewExitError(err, 2)
				}

				output, err := moment.Date(out_cal, out_form)
				if err != nil {
					return cli.NewExitError(err, 2)
				}

				fmt.Println(output)
				return nil
			},
		},
	}...)
}
