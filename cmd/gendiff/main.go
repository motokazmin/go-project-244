package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

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

			if len(args) < 2 {
				fmt.Println("Error: two file paths are required")
				return cli.ShowAppHelp(cmd)
			}

			format := cmd.String("format")

			path1, err := filepath.Abs(args[0])
			if err != nil {
				return fmt.Errorf("invalid path1: %w", err)
			}

			path2, err := filepath.Abs(args[1])
			if err != nil {
				return fmt.Errorf("invalid path2: %w", err)
			}

			data1, err := code.ParseFile(path1)
			if err != nil {
				return fmt.Errorf("failed to parse %s: %w", path1, err)
			}

			data2, err := code.ParseFile(path2)
			if err != nil {
				return fmt.Errorf("failed to parse %s: %w", path2, err)
			}

			diff := code.GenDiff(data1, data2, format)
			fmt.Println(diff)

			fmt.Printf("Comparing files:\n  Path1: %s\n  Path2: %s\n", path1, path2)

			return nil
		},
	}

	if err := command.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
