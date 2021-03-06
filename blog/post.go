package blog

import (
	"fmt"

	"github.com/fiskeben/microdotblog"
)

// PostError is returned when unable to publish a post.
type PostError struct {
	msg string
}

func (e PostError) Error() string {
	return e.msg
}

func newPostError(msg string, err error) PostError {
	reason := fmt.Sprintf("%s Reason: %s", msg, err.Error())
	return PostError{msg: reason}
}

// Poster defines how to post.
type Poster interface {
	Post(message string) (*microdotblog.Post, error)
}

// Post publishes a post to micro.blog. If fromFile is specified the contents of
// that file will be published, otherwise an editor will open and let the user
// type a post.
func Post(client Poster, message string, dryRun bool) (*string, error) {
	if dryRun {
		result := doDryRun(message)
		return &result, nil
	}

	post, err := client.Post(message)
	if err != nil {
		return nil, err
	}

	return &post.URL, nil
}
