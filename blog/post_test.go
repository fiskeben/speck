package blog

import (
	"fmt"
	"testing"

	api "github.com/fiskeben/microdotblog"
)

type posterMock struct {
	result *api.Post
}

func (p posterMock) Post(message string) (*api.Post, error) {
	if p.result == nil {
		return nil, fmt.Errorf("Posting failed")
	}
	return p.result, nil
}

func makePosterMock(result string) Poster {
	return posterMock{
		result: &api.Post{
			ContentHTML: result,
		},
	}
}

func TestPost(t *testing.T) {
	postData := "This is the post"
	dryrun := false

	mock := makePosterMock(postData)

	_, err := Post(mock, postData, dryrun)
	if err != nil {
		t.Fatal("Error posting", err.Error())
	}

}
