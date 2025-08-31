package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/aybangueco/dotsync/internal/config"
)

var InitCommand = &cli.Command{
	Name:  "init",
	Usage: "Initialize a config file",
	Action: func(ctx context.Context, c *cli.Command) error {
		dir, err := os.ReadDir(".")
		if err != nil {
			return err
		}

		err = config.WriteConfig(dir)
		if err != nil {
			return err
		}

		fmt.Println("Dotsync configuration has been added")

		return nil
	},
}
