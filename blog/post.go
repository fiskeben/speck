package blog

import (
	"fmt"
	"strings"

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
		response := []string{}
		response = append(response, "Dry run mode - your post will not be published.")
		response = append(response, "Here's what your post would be like:")
		response = append(response, "")
		response = append(response, post)
		response = append(response, "Run again without the --dry-run flag to publish your words.")
		result := strings.Join(response, "\n")
		return &result, nil
	}

	posted, err := client.Post(post)
	if err != nil {
		return nil, err
	}

	result := fmt.Sprintf("Your post was created: %s\n", posted.URL)
	return &result, nil
}
