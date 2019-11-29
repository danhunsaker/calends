// +build wasm

package main

import (
	"errors"

	"github.com/urfave/cli/v2"
)

var defaultAction = func(c *cli.Context) error {
	return errors.New("Batch mode not supported in WASM")
}
