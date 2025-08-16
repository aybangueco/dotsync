package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var VersionCommand = &cli.Command{
	Name:  "version",
	Usage: "Display's current version number",
	Action: func(ctx context.Context, c *cli.Command) error {
		fmt.Println("DotSync v1.0.5")
		return nil
	},
}
