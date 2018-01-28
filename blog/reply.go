package blog

import (
	"net/url"
)

// Replyer defines how to reply to other messages.
type Replyer interface {
	Reply(ID int64, message string) (*url.URL, error)
}

// Reply posts a message as a reply to another message identified by an ID.
func Reply(client Replyer, ID int64, post string, dryRun bool) (*string, error) {
	if dryRun {
		result := doDryRun(post)
		return &result, nil
	}

	url, err := client.Reply(ID, post)
	if err != nil {
		return nil, err
	}

	result := url.String()
	return &result, nil
}
