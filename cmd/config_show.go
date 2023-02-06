package cmd

import (
	"github.com/piaverous/pira/pira"
	"github.com/spf13/cobra"
)

func buildConfigShowCommand(app *pira.App) *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Print current config",
		RunE: func(cmd *cobra.Command, args []string) error {
			app.PrintConfig()
			return nil
		},
	}
}
