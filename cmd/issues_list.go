package cmd

import (
	"fmt"
	"os"

	"github.com/piaverous/pira/pira"
	"github.com/spf13/cobra"
)

func buildIssuesListCommand(app *pira.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:        "list [flags] sprint",
		Short:      "List issues",
		Args:       cobra.MinimumNArgs(1),
		ArgAliases: []string{"sprint"},
		RunE: func(cmd *cobra.Command, args []string) error {
			sprint := args[0]
			response, err := app.ListJiraIssues(sprint)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			for _, issue := range response.Issues {
				fmt.Printf("%s : %s\n", issue.Key, issue.Fields.Summary)
				if issue.Fields.Assignee.DisplayName != "" {
					fmt.Printf("\t%-15s : %s\n", "assigned to", issue.Fields.Assignee.DisplayName)
				}
				fmt.Printf("\t%-15s : %s\n", "status", issue.Fields.Status.Name)

				for _, cField := range issue.Fields.CustomFields {
					fmt.Printf("\t%-15s : '%s' - %s\n", "custom field", cField.Alias, cField.Value)
				}
			}
			return nil
		},
	}

	return cmd
}
