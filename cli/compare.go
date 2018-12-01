package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/danhunsaker/calends"
	"github.com/urfave/cli"
)

func init() {
	commands = append(commands, []cli.Command{
		{
			Name:      "compare",
			Usage:     "compare two timestamps using a given comparison method",
			ArgsUsage: "<method> [<stamp1> [<stamp2>]]",

			Action: func(c *cli.Context) error {
				if c.NArg() < 1 || c.NArg() > 3 {
					return errArgMismatch
				}

				method := c.Args()[0]
				stamp1 := ""
				if c.NArg() > 1 {
					stamp1 = c.Args()[1]
				} else {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						stamp1 = fmt.Sprintf("%s%s", stamp1, scanner.Text())
					}
				}
				stamp2 := ""
				if c.NArg() == 3 {
					stamp2 = c.Args()[2]
				} else {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						stamp2 = fmt.Sprintf("%s%s", stamp2, scanner.Text())
					}
				}

				var moment1 calends.Calends
				err := json.Unmarshal([]byte(stamp1), &moment1)
				if err != nil {
					return cli.NewExitError(err, 2)
				}

				var moment2 calends.Calends
				err = json.Unmarshal([]byte(stamp2), &moment2)
				if err != nil {
					return cli.NewExitError(err, 2)
				}

				output, err := callComparisonMethod(method, moment1, moment2)
				if err != nil {
					return cli.NewExitError(err, 1)
				}

				fmt.Printf("%v\n", output)
				return nil
			},
		},
	}...)
}

func callComparisonMethod(method string, moment1, moment2 calends.Calends) (ret bool, err error) {
	switch method {
	case "contains":
		ret = moment1.Contains(moment2)
	case "overlaps":
		ret = moment1.Overlaps(moment2)
	case "abuts":
		ret = moment1.Abuts(moment2)
	case "same":
		ret = moment1.IsSame(moment2)
	case "shorter":
		ret = moment1.IsShorter(moment2)
	case "same-duration":
		ret = moment1.IsSameDuration(moment2)
	case "longer":
		ret = moment1.IsLonger(moment2)
	case "before":
		ret = moment1.IsBefore(moment2)
	case "start-before":
		ret = moment1.StartsBefore(moment2)
	case "end-before":
		ret = moment1.EndsBefore(moment2)
	case "during":
		ret = moment1.IsDuring(moment2)
	case "start-during":
		ret = moment1.StartsDuring(moment2)
	case "end-during":
		ret = moment1.EndsDuring(moment2)
	case "after":
		ret = moment1.IsAfter(moment2)
	case "start-after":
		ret = moment1.StartsAfter(moment2)
	case "end-after":
		ret = moment1.EndsAfter(moment2)
	default:
		err = errors.New("Unsupported comparison method")
	}

	return
}
