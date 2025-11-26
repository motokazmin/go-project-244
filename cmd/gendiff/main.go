package main

import (
	"context"
	"log"
	"os"

	cli "github.com/urfave/cli/v3"
)

func main() {
	command := &cli.Command{
		Name:      "gendiff",
		Usage:     "compares two configuration files and shows a difference",
		UsageText: "gendiff [global options]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Usage:   "output format (default: stylish)",
				Value:   "stylish",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args().Slice()

			if len(args) == 0 {
				return cli.ShowAppHelp(cmd)
			}

			return nil
		},
	}

	if err := command.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
