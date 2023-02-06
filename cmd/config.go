package cmd

import (
	"github.com/piaverous/pira/pira"
	"github.com/spf13/cobra"
)

func buildConfigCommand(app *pira.App) *cobra.Command {
	list := &cobra.Command{
		Use:   "config",
		Short: "Tools to configure pira",
	}

	list.AddCommand(buildConfigShowCommand(app))
	return list
}
