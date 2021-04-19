package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/Liu-Chunhui/line-coverage/cmd/report"
)

var (
	Version = "dev" // v0.1.0
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Version: Version,
		Name:    "line-coverage",
		Usage:   "scans the files which are described in the coverage profile(e.g. coverage.out) to calculates the line coverage",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "coverprofile",
				Aliases:  []string{"c"},
				Usage:    "coverage profile filename",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "module",
				Aliases:  []string{"m"},
				Usage:    "module name.(e.g. github.com/Liu-Chunhui/line-coverage)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "location",
				Aliases:  []string{"l"},
				Usage:    "the root level location of the files are described in the coverage profile.",
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "enable debug model. More details is provided.",
			},
		},
		Action: func(c *cli.Context) error {
			initLogging(c.Bool("debug"))

			err := report.Report(
				c.String("coverprofile"),
				c.String("module"),
				c.String("location"),
			)

			if err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func initLogging(debugMode bool) {
	log.SetOutput(os.Stdout)

	if debugMode {
		log.SetLevel(log.DebugLevel)
		log.Debug("debug model is enabled.")
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
