package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Count up or down"
	app.Commands = []*cli.Command{
		{
			Name: "up",
			Aliases: []string{"u"},
			Usage: "Count up",
			Flags: []cli.Flag {
				&cli.IntFlag{
					Name: "stop",
					Aliases: []string{"s"},
					Usage: "Value to count up to",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				stop := c.Int("stop")
				if stop <= 0 {
					return errors.New("stop cannot be negative or less zero")
				}
				for i:= 1; i <= stop; i++ {
					fmt.Println(i)
				}

				return nil
			},
		},
		{
			Name: "down",
			Aliases: []string{"d"},
			Usage: "Count down",
			Flags: []cli.Flag {
				&cli.IntFlag{
					Name: "start",
					Aliases: []string{"s"},
					Usage: "Value to count down from",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start <= 0 {
					return errors.New("start cannot be negative or less zero")
				}
				for i:= start; i >= 0; i-- {
					fmt.Println(i)
				}

				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
