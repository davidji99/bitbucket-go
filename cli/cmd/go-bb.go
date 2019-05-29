package main

import (
	"fmt"
	"github.com/davidji99/go-bitbucket/cli/command"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"sort"
	"strings"
)

var (
	CliName    = "go-bb"
	Version    = "0.0.1"
	BuildSHA   = "<dev>"
	GibVersion = strings.Join([]string{Version, BuildSHA}, "-")
)

func main() {
	app := cli.NewApp()
	app.Name = CliName
	app.Usage = "Bitbucket APIv2 CLI tool"
	app.HelpName = CliName
	app.ArgsUsage = ""
	app.UsageText = fmt.Sprintf("%s <COMMAND> [options]", CliName)
	app.Version = GibVersion
	app.Description = "https://github.com/davidji99/go-bitbucket.git"
	app.Authors = []cli.Author{
		{Name: "David Ji"},
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version",
		Usage: "show current version",
	}

	app.Before = func(c *cli.Context) error {
		return nil
	}

	// Stores all commands
	commands := []cli.Command{
		command.IssueList(),
	}
	app.Commands = commands

	app.CommandNotFound = func(ctx *cli.Context, command string) {
		fmt.Printf("Command not found: %v\n", command)
		os.Exit(1)
	}

	// Sort Flags/Commands alphabetically
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	} else {
		os.Exit(0)
	}
}
