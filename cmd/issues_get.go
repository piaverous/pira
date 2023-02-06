package cmd

import (
	"fmt"

	"github.com/piaverous/pira/pira"
	"github.com/spf13/cobra"
)

func buildIssuesGetCommand(app *pira.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:        "get",
		Short:      "Get an issue",
		Args:       cobra.MinimumNArgs(1),
		ArgAliases: []string{"issue_id"},
		RunE: func(cmd *cobra.Command, args []string) error {
			issueId := args[0]
			response, err := app.GetJiraIssue(issueId)
			if err != nil {
				return err
			}
			if response.Key == "" {
				fmt.Printf("No issue for ID %s\n", issueId)
			} else {
				fmt.Printf("Found issue %s : %s\n", issueId, response.Fields.Summary)
			}
			return nil
		},
	}

	return cmd
}
