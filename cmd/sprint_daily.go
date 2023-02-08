package cmd

import (
	"fmt"
	"os"

	"github.com/piaverous/pira/pira"
	"github.com/piaverous/pira/pira/types"
	"github.com/spf13/cobra"
)

func listContains(list []string, str string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

func buildSprintDailyCommand(app *pira.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:        "daily",
		Short:      "Daily Sprint recap",
		Args:       cobra.MinimumNArgs(1),
		ArgAliases: []string{"sprint"},
		RunE: func(cmd *cobra.Command, args []string) error {
			sprint := args[0]
			response, err := app.ListJiraIssues(sprint)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			sprintReport := map[string][]types.JiraIssue{}
			for _, category := range app.Config.Jira.SprintConfig.TicketStatuses {
				sprintReport[category.Name] = []types.JiraIssue{}
			}

			for _, issue := range response.Issues {
				for _, category := range app.Config.Jira.SprintConfig.TicketStatuses {
					if listContains(category.Statuses, issue.Fields.Status.Name) {
						sprintReport[category.Name] = append(sprintReport[category.Name], issue)
					}
				}
			}

			for _, category := range app.Config.Jira.SprintConfig.TicketStatuses {
				categoryStoryPoints := 0
				for _, issue := range sprintReport[category.Name] {
					points, err := app.StoryPointsFromIssue(issue)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						os.Exit(1)
					}
					categoryStoryPoints += points
				}
				fmt.Printf("| %15s |-> %2d issues | %2d story points\n", category.Name, len(sprintReport[category.Name]), categoryStoryPoints)
			}

			return nil
		},
	}

	return cmd
}
