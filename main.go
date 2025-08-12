package main

import (
	"context"
	"log"
	"os"

	"github.com/aybangueco/dotsync/cmd"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "dotsync",
		Usage: "Easily sync your dotfiles changes from your local repository to a specified path",
		Commands: []*cli.Command{
			cmd.InitCommand,
			cmd.SyncCommand,
			cmd.RemoveCommand,
			cmd.VersionCommand,
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Println(err)
	}
}
