package blog

import (
	"fmt"

	api "github.com/fiskeben/microdotblog"
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
	Post(message string) (*api.Post, error)
}

// Post publishes a post to micro.blog. If fromFile is specified the contents of
// that file will be published, otherwise an editor will open and let the user
// type a post.
func Post(client Poster, post string, dryRun bool) (*string, error) {
	if dryRun {
		result := doDryRun(post)
		return &result, nil
	}

	_, err := client.Post(post)
	if err != nil {
		return nil, err
	}

	// Currently the Micro.blog API doesn't return the created resource.
	result := "Your post was created"
	return &result, nil
}
