package main

import (
	"context"

	gmailapi "github.com/mikeschuld/inbox-sweep/gmail-api"
	"github.com/mikeschuld/inbox-sweep/gmail-api/api"
)

func main() {
	ctx := context.Background()

	gmail, err := api.NewGmail(ctx, "credentials.json")
	if err != nil {
		panic(err)
	}
	svc, err := gmailapi.NewService(gmailapi.Config{
		GmailAPI: gmail,
	})
	if err != nil {
		panic(err)
	}

	svc.ListenAndServe(ctx)
}
