package cmd

import (
	"github.com/piaverous/pira/pira"
	"github.com/spf13/cobra"
)

func buildIssuesCommand(app *pira.App) *cobra.Command {
	// TODO: implement the "list" sub-command here.
	list := &cobra.Command{
		Use:   "issues",
		Short: "Manage issues in your Jira instance.",
	}

	list.AddCommand(buildIssuesListCommand(app))
	list.AddCommand(buildIssuesGetCommand(app))

	return list
}
