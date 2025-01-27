package filesystem

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrFailedToCheckFile   = errors.New("failed to check file")
	ErrInvalidFilePath     = errors.New("invalid file path")
	ErrFailedToCreateFile  = errors.New("failed to create file")
	ErrFailedToWriteToFile = errors.New("failed to write to file")
	ErrFailedToGetFiles    = errors.New("failed to get files")
	ErrFailedToReadFile    = errors.New("failed to read file")
)

type Filesystem struct{}

func NewFilesystem() *Filesystem {
	return &Filesystem{}
}

func (f *Filesystem) Exists(targetFile string) (bool, error) {
	_, err := os.Stat(targetFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, fmt.Errorf("%w: %w", ErrFailedToCheckFile, err)
	}

	return true, nil
}

func (f *Filesystem) WriteToFile(targetFile string, content io.Reader) (func() error, error) {
	cleanPath := filepath.Clean(targetFile)
	if filepath.IsAbs(cleanPath) || strings.Contains(cleanPath, "..") {
		return nil, fmt.Errorf("%w: %s", ErrInvalidFilePath, cleanPath)
	}

	out, err := os.Create(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToCreateFile, err)
	}

	cleanup := func() error {
		return out.Close()
	}

	_, err = io.Copy(out, content)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToWriteToFile, err)
	}

	return cleanup, nil
}

func (f *Filesystem) ListFiles(targetDir string, includeDotFiles bool) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(targetDir, "*"))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToGetFiles, err)
	}

	if includeDotFiles {
		return files, nil
	}

	// Remove dotfiles from the list
	filteredFiles := make([]string, 0, len(files))

	for _, file := range files {
		base := filepath.Base(file)
		if !strings.HasPrefix(base, ".") {
			filteredFiles = append(filteredFiles, file)
		}
	}

	return filteredFiles, nil
}

func (f *Filesystem) Cat(file string) ([]string, error) {
	lines := []string{}

	content, err := os.OpenFile(file, os.O_RDONLY, 0) //nolint:gosec
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToReadFile, err)
	}
	defer content.Close() //nolint:errcheck

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func (f *Filesystem) Diff(firstFile string, secondFile string) ([]string, error) {
	firstContent, err := os.OpenFile(firstFile, os.O_RDONLY, 0) //nolint:gosec
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToReadFile, err)
	}
	defer firstContent.Close() //nolint:errcheck

	secondContent, err := os.OpenFile(secondFile, os.O_RDONLY, 0) //nolint:gosec
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToReadFile, err)
	}
	defer secondContent.Close() //nolint:errcheck

	firstScanner := bufio.NewScanner(firstContent)
	secondScanner := bufio.NewScanner(secondContent)

	diff := []string{}

	for secondScanner.Scan() {
		currentLine := secondScanner.Text()

		for firstScanner.Scan() {
			if currentLine == firstScanner.Text() {
				break
			}

			diff = append(diff, firstScanner.Text())
		}
	}

	return diff, nil
}
