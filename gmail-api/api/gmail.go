package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailAPI struct {
	client  *http.Client
	service *gmail.Service
}

func NewGmail(ctx context.Context, credentialsFile string) (*GmailAPI, error) {
	creds, err := os.ReadFile(credentialsFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read credentials secret file %q: %w", credentialsFile, err)
	}

	cfg, err := google.ConfigFromJSON(creds, gmail.GmailReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse credentials secret file to config: %w", err)
	}

	client, err := getClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to get http client from config: %w", err)
	}

	svc, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("failed to create gmail service")
	}

	return &GmailAPI{
		client:  client,
		service: svc,
	}, nil
}

func getClient(cfg *oauth2.Config) (*http.Client, error) {
	return http.DefaultClient, nil
}
