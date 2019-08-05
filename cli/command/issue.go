package command

import (
	"fmt"
	"github.com/davidji99/bitbucket-go/bitbucket"
	cli2 "github.com/davidji99/bitbucket-go/cli/cli"
	"gopkg.in/urfave/cli.v1"
)

func Issues(bbcli *cli2.BBCli) cli.Command {
	return cli.Command{
		Name:  "issues",
		Usage: "Interact with Bitbucket issues",
		Subcommands: []cli.Command{
			{
				Name:  "list",
				Usage: "Returns the issues in the issue tracker",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name: GetFlagCliString("owner"),
					},
					cli.StringFlag{
						Name: GetFlagCliString("repo"),
					},
					cli.StringFlag{
						Name: GetFlagCliString("query"),
					},
				},
				Action: func(c *cli.Context) error {
					api := bitbucket.NewBasicAuth(bbcli.GetUser(), bbcli.GetPass())

					result, _, err := api.Issues.List(c.String("owner"), c.String("repo"), nil)
					if err != nil {
						return err
					}

					for _, i := range result.Values {
						fmt.Println(i)
					}

					return nil
				},
			},
			{
				Name:  "create",
				Usage: "Create a new issue",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name: GetFlagCliString("title"),
					},
					cli.StringFlag{
						Name: GetFlagCliString("kind"),
					},
					cli.StringFlag{
						Name: GetFlagCliString("priority"),
					},
					cli.StringFlag{
						Name: GetFlagCliString("owner"),
					},
					cli.StringFlag{
						Name: GetFlagCliString("repo"),
					},
				},
				Action: func(c *cli.Context) error {
					api := bitbucket.NewBasicAuth(bbcli.GetUser(), bbcli.GetPass())
					newIssueOpts := &bitbucket.IssueRequest{
						Title:    NewStringPointer(c.String("title")),
						Kind:     NewStringPointer(c.String("kind")),
						Priority: NewStringPointer(c.String("priority")),
					}

					result, _, err := api.Issues.List(c.String("owner"), c.String("repo"), newIssueOpts)
					if err != nil {
						return err
					}

					fmt.Println(result)
					return nil
				},
			},
		},
	}
}
