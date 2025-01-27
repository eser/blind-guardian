package main

import (
	"github.com/eser/blind-guardian/cmd/bg/subcommands"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "bg",
		Short: "Blind Guardian",
		Long:  `Blind Guardian is a tool for managing safety of your system.`,
	}

	cmd.AddCommand(subcommands.CmdUpdate())
	cmd.AddCommand(subcommands.CmdDiff())

	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
