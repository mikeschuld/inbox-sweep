package gmailapi

import "context"

// API required for interacting with Gmail
type API interface{}

// Config configures a new Service
type Config struct {
	GmailAPI API
}

// NewService creates a new Service
func NewService(cfg Config) (*Service, error) {
	return &Service{
		GmailAPI: cfg.GmailAPI,
	}, nil
}

// Service is a gmail API service
type Service struct {
	GmailAPI API
}

// ListenAndServe starts the service listening
func (s *Service) ListenAndServe(ctx context.Context) {

}
