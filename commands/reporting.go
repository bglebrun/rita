package commands

import (
	"github.com/ocmdev/rita/database"
	"github.com/ocmdev/rita/reporting"
	"github.com/urfave/cli"
)

func init() {
	command := cli.Command{

		Name: "html-report",
		Usage: "Create an html report for an analyzed database. " +
			"If no database is specified, a report will be created for every database.",
		ArgsUsage: "[database]",
		Flags: []cli.Flag{
			configFlag,
		},
		Action: func(c *cli.Context) error {
			res := database.InitResources(c.String("config"))
			databaseName := c.Args().Get(0)
			var databases []string
			if databaseName != "" {
				databases = append(databases, databaseName)
			} else {
				databases = res.MetaDB.GetAnalyzedDatabases()
			}
			err := reporting.PrintHTML(databases, res)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil
		},
	}
	bootstrapCommands(command)
}
