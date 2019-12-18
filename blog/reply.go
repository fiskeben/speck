package blog

import "github.com/fiskeben/microdotblog"

// Replyer defines how to reply to other messages.
type Replyer interface {
	Reply(ID int64, message string) (*microdotblog.Post, error)
}

// Reply posts a message as a reply to another message identified by an ID.
func Reply(client Replyer, ID int64, message string, dryRun bool) (*string, error) {
	if dryRun {
		result := doDryRun(message)
		return &result, nil
	}

	post, err := client.Reply(ID, message)
	if err != nil {
		return nil, err
	}

	return &post.URL, nil
}
