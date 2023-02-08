package cmd

import (
	"github.com/piaverous/pira/pira"
	"github.com/spf13/cobra"
)

func buildSprintCommand(app *pira.App) *cobra.Command {
	sprint := &cobra.Command{
		Use:   "sprint",
		Short: "Get sprint info from your Jira instance.",
	}

	sprint.AddCommand(buildSprintDailyCommand(app))

	return sprint
}
