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

func CmdDiff() *cobra.Command {
	cmd := &cobra.Command{ //nolint:exhaustruct
		Use:   "diff",
		Short: "Diffs two blocklists",
		Long:  `Diffs two blocklists`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDiff(cmd.Context())
		},
	}

	return cmd
}

func runDiff(ctx context.Context) error {
	appConfig, err := bg.Init()
	if err != nil {
		return err //nolint:wrapcheck
	}

	fetcher := fetcher.NewFetcher()
	filesystem := filesystem.NewFilesystem()

	updater := updater.NewService(&appConfig.Update, slog.Default(), fetcher, filesystem)

	err = updater.Diff(ctx)
	if err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
