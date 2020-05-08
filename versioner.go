package main

import (
	"github.com/urfave/cli"
	"gitlab.com/hectorjsmith/git-versioner/git"
	"gitlab.com/hectorjsmith/git-versioner/versioner/fix"
	"gitlab.com/hectorjsmith/git-versioner/versioner/latest"
	"gitlab.com/hectorjsmith/git-versioner/versioner/rel"
	"log"
	"os"
)

// Application version - Global variable set during build process
var appVersion string

func main() {
	if appVersion == "" {
		appVersion = "development"
	}

	app := cli.NewApp()
	app.Name = "Versioner"
	app.Usage = "Manage releases with git tags and branches"
	app.Description = "Small CLI application to manage releases using git tags and branches."
	app.Version = appVersion

	var minor bool
	var major bool
	var testRelease bool
	var version string
	var verbose bool
	var message string
	app.Commands = []cli.Command{
		{
			Name:  "release",
			Usage: "release a new version - by default the branch name is used to parse the new version",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "message",
					Usage:       "Message to put in git tag. Using this will create an annotated tag.",
					Destination: &message,
				},
				cli.BoolFlag{
					Name:        "major",
					Usage:       "major release (v1.5.10 -> v2.0.0)",
					Destination: &major,
				},
				cli.BoolFlag{
					Name:        "minor",
					Usage:       "minor release (v1.5.10 -> v1.6.0)",
					Destination: &minor,
				},
				cli.BoolFlag{
					Name:        "test",
					Usage:       "test release (v1.5.10-5-g600d3f2)",
					Destination: &testRelease,
				},
			},
			Before: func(c *cli.Context) error { runVersionerStartupValidations(true); return nil },
			Action: func(c *cli.Context) error { return rel.Run(major, minor, testRelease, message) },
			After: func(c *cli.Context) error { return logOperationComplete() },
		},
		{
			Name:  "fix",
			Usage: "create a fix branch for an existing version",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "version, v",
					Usage:       "version to fix (e.g. '1.2.0')",
					Destination: &version,
				},
			},
			Before: func(c *cli.Context) error { runVersionerStartupValidations(true); return nil },
			Action: func(c *cli.Context) error { return fix.Run(version) },
			After: func(c *cli.Context) error { return logOperationComplete() },
		},
		{
			Name:  "latest",
			Usage: "return the latest version info (highest value) - this is the version that will be incremented for release",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "verbose, v",
					Usage:       "print more useful information about the latest version",
					Destination: &verbose,
				},
			},
			Before: func(c *cli.Context) error { runVersionerStartupValidations(false); return nil },
			Action: func(c *cli.Context) error { return latest.Run(verbose) },
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runVersionerStartupValidations(ensureCleanRepo bool) {
	repo, err := git.GetRepositoryForPath(".")
	if err != nil {
		log.Fatal("Must run this tool in a git repository", err)
	}
	if ensureCleanRepo && !repo.IsClean() {
		log.Fatal("Must run this tool on a clean repository")
	}
}

func logOperationComplete() error {
	log.Print("Done")
	return nil
}
