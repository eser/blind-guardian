package fetcher

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var (
	ErrFailedToParseUri              = errors.New("failed to parse URI")
	ErrInvalidScheme                 = errors.New("invalid URL scheme - must be HTTPS")
	ErrFailedToCreateRequest         = errors.New("failed to create request")
	ErrFailedToDownloadRemoteContent = errors.New("failed to download remote content")
	ErrUnexpectedStatusCode          = errors.New("unexpected status code")
)

const defaultTimeout = 30 * time.Second

type Fetcher struct {
	Timeout time.Duration
}

func NewFetcher() *Fetcher {
	return &Fetcher{Timeout: defaultTimeout}
}

func (f *Fetcher) WithTimeout(timeout time.Duration) *Fetcher {
	f.Timeout = timeout

	return f
}

func (f *Fetcher) Download(ctx context.Context, uri string) (io.ReadCloser, func() error, error) {
	parsedURL, err := url.Parse(uri)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", ErrFailedToParseUri, err)
	}

	if parsedURL.Scheme != "https" {
		return nil, nil, fmt.Errorf("%w: %s", ErrInvalidScheme, parsedURL.Scheme)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", ErrFailedToCreateRequest, err)
	}

	client := &http.Client{ //nolint:exhaustruct
		Timeout: f.Timeout,
	}

	resp, err := client.Do(req) //nolint:bodyclose
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", ErrFailedToDownloadRemoteContent, err)
	}

	cleanup := func() error {
		return resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("%w: %d", ErrUnexpectedStatusCode, resp.StatusCode)
	}

	return resp.Body, cleanup, nil
}
