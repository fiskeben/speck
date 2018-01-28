package blog

import (
	"fmt"
	"net/url"
	"testing"
)

type posterMock struct {
	result *url.URL
}

func (p posterMock) Post(message string) (*url.URL, error) {
	if p.result == nil {
		return nil, fmt.Errorf("Posting failed")
	}
	return p.result, nil
}

func makePosterMock(result string) Poster {
	u, _ := url.Parse(result)
	return posterMock{
		result: u,
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
