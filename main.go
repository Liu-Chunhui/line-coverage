package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"

	"github.com/Liu-Chunhui/line-coverage/cmd/report"
)

var (
	Version = "dev" // v1.0.0
	Commit  = ""
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Version:  Version,
		Compiled: time.Now().UTC(),
		Name:     "line-coverage",
		Usage:    "scans the files which are described in the coverage profile(e.g. coverage.out) to calculates the line coverage",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "coverprofile",
				Aliases:  []string{"c"},
				Usage:    "coverage profile filename",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "gomod",
				Aliases:  []string{"m"},
				Usage:    "go.mod with path",
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
				c.String("gomod"),
			)

			if err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

func initLogging(debugMode bool) {
	logrus.SetOutput(os.Stdout)
	logrus.WithField("Version", Version)
	logrus.WithField("CommitSHA", Commit)

	if debugMode {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("debug model is enabled.")
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
