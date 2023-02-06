package cmd

import (
	"github.com/piaverous/pira/pira"
	"github.com/spf13/cobra"
)

func buildSprintCommand(app *pira.App) *cobra.Command {
	// TODO: implement the "list" sub-command here.
	list := &cobra.Command{
		Use:   "sprint",
		Short: "Get sprint info from your Jira instance.",
	}

	list.AddCommand(buildSprintDailyCommand(app))

	return list
}
