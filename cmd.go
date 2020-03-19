package main

import (
	"github.com/urfave/cli"
	"os"
	"sort"
	"strings"
)

type config struct {
	Port int
	Wait int
	Args []string
	Arg  string
}

var cfg config

const DEF_WSEC = 0
const DEF_PORT = 80
const DEF_STR = "{\"Message\":\"Hello ppp!\"}"

func cmd() error {
	app := &cli.App{
		Name:    "ppp",
		Usage:   "This is cli ready-to-run web server.",
		Version: "0.0.1",
		Action:  action,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "wait",
				Aliases: []string{"w"},
				Value:   DEF_WSEC,
				Usage:   "If this flag is set, web server pause for the specified seconds before returning a response",
			},
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   DEF_PORT,
				Usage:   "You can specify the port for this web server.",
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	err := app.Run(os.Args)
	return err
}

func action(c *cli.Context) error {
	cfg.Wait = c.Int("wait")
	cfg.Port = c.Int("port")
	cfg.Args = c.Args().Slice()
	cfg.Arg = strings.Join(cfg.Args, " ")
	if cfg.Arg == "" {
		cfg.Arg = DEF_STR
	}
	return nil
}
