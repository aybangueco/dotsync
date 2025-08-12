package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/aybangueco/dotsync/internal/config"
	"github.com/aybangueco/dotsync/internal/helpers"
	"github.com/urfave/cli/v3"
)

var RemoveCommand = &cli.Command{
	Name:  "remove",
	Usage: "The opposite of sync command, it removes targeted files and directories",
	Action: func(ctx context.Context, c *cli.Command) error {
		conf, err := config.ReadConfig()
		if err != nil {
			return err
		}

		for _, c := range conf {
			if c.Target == "" {
				return fmt.Errorf("%s's target path must not be empty", c.Source)
			}

			if c.Source == "" {
				return errors.New("one of the sources are empty, either remove it or add some source")
			}

			var targetDoesExist bool
			_, err := os.Stat(c.Target)
			if err == nil {
				targetDoesExist = true
			}

			if os.IsNotExist(err) {
				targetDoesExist = false
			}

			target, err := helpers.ExpandPath(c.Target)
			if err != nil {
				return err
			}

			if targetDoesExist {
				if c.IsDirectory {
					rmDir := exec.Command("rm", "-r", target)

					if output, err := rmDir.CombinedOutput(); err != nil {
						return fmt.Errorf("rm -r failed: %v\nOutput: %s", err, string(output))
					}
				} else {
					rm := exec.Command("rm", target)

					if output, err := rm.CombinedOutput(); err != nil {
						return fmt.Errorf("rm failed: %v\nOutput: %s", err, string(output))
					}
				}
			}
		}

		return nil
	},
}
