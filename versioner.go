package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gitlab.com/hectorjsmith/git-versioner/cmd/checkout"
	"gitlab.com/hectorjsmith/git-versioner/cmd/fix"
	"gitlab.com/hectorjsmith/git-versioner/cmd/latest"
	"gitlab.com/hectorjsmith/git-versioner/cmd/list"
	"gitlab.com/hectorjsmith/git-versioner/cmd/rel"
	"gitlab.com/hectorjsmith/git-versioner/git"
	"log"
	"os"
)

// Global variables set during build process
var (
	version = "dev"
)

func main() {
	app := cli.NewApp()
	app.Name = "Versioner"
	app.Usage = "Manage releases with git tags and branches"
	app.Description = "Small CLI application to manage releases using git tags and branches."
	app.Version = version

	app.Commands = []*cli.Command{
		listCommand(),
		releaseCommand(),
		fixCommand(),
		latestCommand(),
		checkoutCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func listCommand() *cli.Command {
	options := list.CommandOptions{}

	return &cli.Command{
		Name: "list",
		Aliases: []string{"ls"},
		Usage: "List all git versions",
		Description: "",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "tag",
				Aliases: []string{"t"},
				Usage: "Print git tags instead of version numbers",
				Destination: &options.Tag,
			},
			&cli.BoolFlag{
				Name: "test",
				Usage: "Include test versions (only applies when --tag is used)",
				Destination: &options.Test,
			},
		},
		Before: func(c *cli.Context) error {
			return runStartupValidations(false)
		},
		Action: func(c *cli.Context) error {
			return list.Run(options)
		},
	}
}

func releaseCommand() *cli.Command {
	options := rel.CommandOptions{}

	return &cli.Command{
		Name:  "release",
		Usage: "Create new version tag",
		Description: "Create a new version git tag named by taking the latest version and incrementing it.\n" +
			"This command assumes the use of semantic versioning. The version string is parsed as: " +
			"<major>.<minor>.<bugfix>\n" +
			"Options are available to increment the major or minor versions.\n" +
			"If no options are provided, the version number to use will be parsed from the current branch.\n" +
			"For example, if run on a branch named 'release/v1.2.3', the new tag would be 'v1.2.3'.\n\n" +
			"The repository must not have un-staged changes - i.e. the repo cannot be dirty.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "message",
				Aliases:     []string{"m"},
				Usage:       "Message to put in git tag. Using this will create an annotated tag.",
				Destination: &options.Message,
			},
			&cli.BoolFlag{
				Name:        "major",
				Usage:       "major release (v1.5.10 -> v2.0.0)",
				Destination: &options.Major,
			},
			&cli.BoolFlag{
				Name:        "minor",
				Usage:       "minor release (v1.5.10 -> v1.6.0)",
				Destination: &options.Minor,
			},
			&cli.BoolFlag{
				Name:        "test",
				Usage:       "test release (v1.5.10-5-g600d3f2)",
				Destination: &options.Test,
			},
		},
		Before: func(c *cli.Context) error {
			return runStartupValidations(true)
		},
		Action: func(c *cli.Context) error {
			return rel.Run(options)
		},
		After: func(c *cli.Context) error {
			log.Print(c.Err())
			if c.Err() == nil {
				log.Print("Done")
			}
			return nil
		},
	}
}

func fixCommand() *cli.Command {
	options := fix.CommandOptions{}

	return &cli.Command{
		Name:  "fix",
		Usage: "Create fix branch",
		Description: "Create a fix branch for the specified version (or latest version).\n" +
			"This command will checkout the selected version (based on the corresponding git tag) and create a new " +
			"fix branch.\n\n" +
			"The repository must not have un-staged changes - i.e. the repo cannot be dirty.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "version",
				Aliases:     []string{"v"},
				Usage:       "version to fix (e.g. '1.2.0')",
				Destination: &options.Version,
			},
		},
		Before: func(c *cli.Context) error {
			return runStartupValidations(true)
		},
		Action: func(c *cli.Context) error {
			return fix.Run(options)
		},
		After: func(c *cli.Context) error {
			log.Print("Done")
			return nil
		},
	}
}

func latestCommand() *cli.Command {
	options := latest.CommandOptions{}

	return &cli.Command{
		Name:  "latest",
		Usage: "Show latest version info",
		Description: "Show the latest version for this repository.\nThe version data is parsed from git tags found in " +
			"the repository.\n\n" +
			"By default prints the version string (e.g. 3.12.1).",
		Before: func(c *cli.Context) error {
			return runStartupValidations(false)
		},
		Action: func(c *cli.Context) error {
			return latest.Run(options)
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Usage:       "print more useful information about the latest version",
				Destination: &options.Verbose,
			},
			&cli.BoolFlag{
				Name:        "tag",
				Aliases:     []string{"t"},
				Usage:       "show the latest version tag instead of the parsed version info",
				Destination: &options.Tag,
			},
		},
	}
}

func checkoutCommand() *cli.Command {
	options := checkout.CommandOptions{}
	return &cli.Command{
		Name:  "checkout",
		Usage: "Check out specific version",
		Description: "Check out the git tag for the specified version. If no version is provided, the latest version " +
			"is used.\n" +
			"The version should be provided in the <major>.<minor>.<bugfix> syntax (e.g. '1.3.4').\n\n" +
			"The repository must not have un-staged changes - i.e. the repo cannot be dirty.",
		Before: func(c *cli.Context) error {
			return runStartupValidations(true)
		},
		Action: func(c *cli.Context) error {
			return checkout.Run(options)
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "latest",
				Usage:       "check out the latest version",
				Destination: &options.Latest,
			},
			&cli.StringFlag{
				Name:        "version",
				Usage:       "check out specific version",
				Destination: &options.Version,
			},
		},
	}
}

func runStartupValidations(ensureCleanRepo bool) error {
	repo, err := git.NewRepository(".")
	if err != nil {
		return fmt.Errorf("not a valid git repository: %v", err)
	}
	if ensureCleanRepo && !repo.IsClean() {
		return fmt.Errorf("git repository is dirty")
	}
	return nil
}
