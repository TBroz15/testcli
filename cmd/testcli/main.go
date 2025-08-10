// my main.go was to blow up

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main()  {
	subCommands := []*cli.Command{
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "test stuff.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "test",
					Aliases:  []string{"t"},
					Usage:    "test stuff.", 
					Required: false,
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				test := cmd.String("test")

				if (test == "") {
					return fmt.Errorf("where text??????")
				}

				fmt.Println("riddle me this batman if quizzes are quizzical then what are tests: %v", test)

				return nil
			},
		},
	}

	cmd := &cli.Command{
		EnableShellCompletion: true,
        Name:  "testcli",
        Usage: "riddle me this batman if quizzes are quizzical then what are tests",
		Commands: subCommands,
    }

	if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}