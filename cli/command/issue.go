package command

import (
	"fmt"
	"github.com/davidji99/go-bitbucket/bitbucket"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func IssueList() cli.Command {
	return cli.Command{
		Name: "issue:list",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: GetFlagCliString("owner"),
			},
			cli.StringFlag{
				Name: GetFlagCliString("repo"),
			},
		},
		Action: func(c *cli.Context) error {
			api := bitbucket.NewBasicAuth(os.Getenv("BB_USER"), os.Getenv("BB_PASS"))

			result, _, err := api.Issues.List(c.String("owner"), c.String("repo"))
			if err != nil {
				return err
			}

			for _, i := range result.Values {
				fmt.Println(i)
			}

			return nil
		},
	}
}
