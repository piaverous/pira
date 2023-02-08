package cmd

import (
	"github.com/piaverous/pira/pira"
	"github.com/spf13/cobra"
)

func buildConfigCommand(app *pira.App) *cobra.Command {
	config := &cobra.Command{
		Use:   "config",
		Short: "Tools to configure pira",
	}

	config.AddCommand(buildConfigShowCommand(app))
	return config
}
