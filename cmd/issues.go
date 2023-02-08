package cmd

import (
	"github.com/piaverous/pira/pira"
	"github.com/spf13/cobra"
)

func buildIssuesCommand(app *pira.App) *cobra.Command {
	issues := &cobra.Command{
		Use:   "issues",
		Short: "Manage issues in your Jira instance.",
	}

	issues.AddCommand(buildIssuesListCommand(app))
	issues.AddCommand(buildIssuesGetCommand(app))

	return issues
}
