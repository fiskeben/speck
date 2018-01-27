package blog

import (
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

	_, err := client.Reply(ID, post)
	if err != nil {
		return nil, err
	}

	// Currently the Micro.blog API doesn't return the created resource.
	result := "Your reply was posted: %s"
	return &result, nil
}
