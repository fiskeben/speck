package blog

import (
	"fmt"

	api "github.com/fiskeben/microdotblog"
)

// Replyer defines how to reply to other messages.
type Replyer interface {
	Reply(ID int64, message string) (*api.Post, error)
}

// Reply posts a message as a reply to another message identified by an ID.
func Reply(client Replyer, ID int64, post string, dryRun bool) (*string, error) {
	if dryRun {
		result := doDryRun(post)
		return &result, nil
	}

	replied, err := client.Reply(ID, post)
	if err != nil {
		return nil, err
	}

	result := fmt.Sprintf("Your reply was posted: %s", replied.URL)
	return &result, nil
}
