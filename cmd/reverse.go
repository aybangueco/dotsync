package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/aybangueco/dotsync/internal/config"
	"github.com/aybangueco/dotsync/internal/helpers"
	"github.com/urfave/cli/v3"
)

var ReverseCommand = &cli.Command{
	Name:  "reverse",
	Usage: "Equivalent of sync command, ",
	Action: func(ctx context.Context, c *cli.Command) error {
		conf, err := config.ReadConfig()
		if err != nil {
			return err
		}

		if err = helpers.ValidateConfig(conf); err != nil {
			return err
		}

		// File operations
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

			fmt.Printf("target: %s exist: %t \n", c.Source, targetDoesExist)

			// Remove existing file and directory
			if targetDoesExist {
				fmt.Printf("Removing %s from %s \n", c.Source, target)

				err = helpers.RemoveFromSource(c)
				if err != nil {
					return err
				}
			}

			// Copy the file or directory to the specified target, this assumes the directory or file is deleted or not existing
			if c.IsDirectory {
				cp := exec.Command("cp", "-r", fmt.Sprintf("%s/%s", target, c.Source), c.Source)
				if output, err := cp.CombinedOutput(); err != nil {
					return fmt.Errorf("Error copying directory: %v\nOutput: %s", err, string(output))
				}
			} else {
				cp := exec.Command("cp", fmt.Sprintf("%s/%s", target, c.Source), c.Source)
				if output, err := cp.CombinedOutput(); err != nil {
					return fmt.Errorf("Error copying file: %v\nOutput: %s", err, string(output))
				}
			}
		}

		return nil
	},
}
