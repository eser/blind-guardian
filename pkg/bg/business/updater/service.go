package updater

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"sort"
	"time"
)

type Fetcher interface {
	Download(ctx context.Context, uri string) (io.ReadCloser, func() error, error)
}

type Filesystem interface {
	Exists(targetFile string) (bool, error)
	WriteToFile(targetFile string, content io.Reader) (func() error, error)
	ListFiles(targetDir string, includeDotFiles bool) ([]string, error)

	Cat(file string) ([]string, error)
	Diff(firstFile string, secondFile string) ([]string, error)
}

type Service struct {
	config     *UpdateConfig
	logger     *slog.Logger
	fetcher    Fetcher
	filesystem Filesystem
}

var ErrUpdateFailed = errors.New("update failed")

func NewService(config *UpdateConfig, logger *slog.Logger, fetcher Fetcher, filesystem Filesystem) *Service { //nolint:lll
	return &Service{config: config, logger: logger, fetcher: fetcher, filesystem: filesystem}
}

func (s *Service) generateFilename() string {
	date := time.Now().Format("20060102")
	outputFilename := fmt.Sprintf(s.config.OutputFile, date)
	outputFile := filepath.Join(s.config.OutputDir, outputFilename)

	return outputFile
}

func (s *Service) Update(ctx context.Context, force bool) error {
	body, dCleanup, err := s.fetcher.Download(ctx, s.config.BlocklistUri)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrUpdateFailed, err)
	}

	defer dCleanup() //nolint:errcheck

	filename := s.generateFilename()

	exists, err := s.filesystem.Exists(filename)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrUpdateFailed, err)
	}

	if !exists || force {
		s.logger.Info("Updating blocklist")

		wtfCleanup, err := s.filesystem.WriteToFile(filename, body)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrUpdateFailed, err)
		}

		defer wtfCleanup() //nolint:errcheck
	}

	return nil
}

func (s *Service) Diff(ctx context.Context) error {
	files, err := s.getLatestFiles()
	if err != nil {
		return err
	}

	if len(files) < 1 {
		return nil
	}

	var lines []string
	if len(files) == 1 {
		lines, err = s.filesystem.Cat(files[0])
	} else {
		lines, err = s.filesystem.Diff(files[0], files[1])
	}

	if err != nil {
		return fmt.Errorf("%w: %w", ErrUpdateFailed, err)
	}

	for _, line := range lines {
		fmt.Println(line) //nolint:forbidigo
	}

	return nil
}

func (s *Service) getLatestFiles() ([]string, error) {
	files, err := s.filesystem.ListFiles(s.config.OutputDir, false)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUpdateFailed, err)
	}

	if len(files) < 2 { //nolint:gomnd,mnd
		return files, nil
	}

	// Sort files in descending order to get latest files first
	sort.Sort(sort.Reverse(sort.StringSlice(files)))

	return files[0:2], nil
}
