package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/aybangueco/dotsync/internal/config"
	"github.com/aybangueco/dotsync/internal/helpers"
	"github.com/urfave/cli/v3"
)

var RemoveCommand = &cli.Command{
	Name:  "remove",
	Usage: "The opposite of sync command, it removes files and directories from target",
	Action: func(ctx context.Context, c *cli.Command) error {
		conf, err := config.ReadConfig()
		if err != nil {
			return err
		}

		if err = helpers.ValidateConfig(conf); err != nil {
			return err
		}

		for _, c := range conf {
			target, err := helpers.ExpandPath(c.Target)
			if err != nil {
				return err
			}

			var targetDoesExist bool
			_, err = os.Stat(target)
			if err == nil {
				targetDoesExist = true
			}

			if os.IsNotExist(err) {
				targetDoesExist = false
			}

			fmt.Printf("target: %s exist: %t \n", target, targetDoesExist)

			// Remove existing file and directory
			if targetDoesExist {
				fmt.Printf("Removing %s from %s \n", c.Source, target)

				err = helpers.RemoveFromTarget(c, target)
				if err != nil {
					return err
				}
			} else {
				fmt.Printf("%s from %s does not exist", c.Source, fmt.Sprintf("%s/%s", c.Target, c.Source))
			}
		}

		return nil
	},
}
