package cmd

import (
	"github.com/piaverous/pira/pira"
	"github.com/spf13/cobra"
)

func New(app *pira.App) *cobra.Command {
	return buildPiraCommand(app)
}

func buildPiraCommand(app *pira.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pira",
		Short: "pira helps you get info from your Jira projects",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return app.Config.Load(cmd.Flags())
		},
	}

	cmd.AddCommand(buildIssuesCommand(app))
	cmd.AddCommand(buildConfigCommand(app))
	cmd.AddCommand(buildSprintCommand(app))

	return cmd
}
