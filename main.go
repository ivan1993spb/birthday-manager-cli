package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

// Application attributes
const (
	APP_NAME    = "birthday-manager-cli"
	APP_USAGE   = "birthday list on the command line"
	APP_VERSION = "1.0.0"
)

// Author info
const (
	AUTHOR_NAME  = "Ivan Pushkin"
	AUTHOR_EMAIL = ""
)

func main() {
	app := cli.NewApp()

	// Attributes
	app.Name = APP_NAME
	app.Usage = APP_USAGE
	app.Version = APP_VERSION
	app.Authors = []cli.Author{{AUTHOR_NAME, AUTHOR_EMAIL}}

	// Birthday storage
	var storage *BirthdayStorage

	// Checking storage file and loading data
	app.Before = func(c *cli.Context) error {
		stf := c.String("file")

		if stf == "" {
			return &Error{"passed empty birthday storage file name"}
		}

		if fi, err := os.Stat(stf); err != nil {
			if os.IsNotExist(err) {
				return &Error{"no such file " + stf}
			}
			return err
		} else if fi.Mode().IsDir() {
			return &Error{"cannot use folder as birthday storege"}
		}

		storage = NewBirthdayStorage(stf)

		if err := storage.Load(); err != nil {
			return &Error{"cannot load birthdays: " + err.Error()}
		}

		return nil
	}

	// Setup app commands
	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create birthday",
			Action: func(c *cli.Context) {

			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "remove birthdays",
			Action: func(c *cli.Context) {

			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "specify name",
				},
			},
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "print name, age and next birthday date",
			Action: func(c *cli.Context) {
				bset := storage.GetBirthdaySet()

				if c.IsSet("duration") {
					bset = bset.FilterByDuration(c.Duration("duration"))
				}
				if c.IsSet("name") {
					bset = bset.FilterByName(c.String("name"))
				}

				for _, bday := range bset {
					fmt.Printf("Birthday: %s, %d %s\n",
						bday.Name,
						bday.GetTime().Day(),
						bday.GetTime().Month(),
					)
				}
			},
			Flags: []cli.Flag{
				&cli.DurationFlag{
					Name:  "duration",
					Usage: "specify time range",
				},
				&cli.StringFlag{
					Name:  "name",
					Usage: "specify name",
				},
			},
		},
	}

	// Application flags
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "file",
			Usage: "specify file of birthdays storage",
		},
	}

	// Run
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
