package subcommands

import (
	"context"
	"log/slog"

	"github.com/eser/blind-guardian/pkg/bg"
	"github.com/eser/blind-guardian/pkg/bg/adapters/fetcher"
	"github.com/eser/blind-guardian/pkg/bg/adapters/filesystem"
	"github.com/eser/blind-guardian/pkg/bg/business/updater"
	"github.com/spf13/cobra"
)

func CmdUpdate() *cobra.Command {
	cmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "update",
		Short: "Updates blind guardian sources",
		Long:  `Updates blind guardian sources such as domain blocklist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			force, err := cmd.Flags().GetBool("force")
			if err != nil {
				return err //nolint:wrapcheck
			}

			return runUpdate(cmd.Context(), force)
		},
	}

	cmd.Flags().BoolP("force", "f", false, "Force update")

	return cmd
}

func runUpdate(ctx context.Context, force bool) error {
	appConfig, err := bg.Init()
	if err != nil {
		return err //nolint:wrapcheck
	}

	fetcher := fetcher.NewFetcher()
	filesystem := filesystem.NewFilesystem()

	updater := updater.NewService(&appConfig.Update, slog.Default(), fetcher, filesystem)

	err = updater.Update(ctx, force)
	if err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
