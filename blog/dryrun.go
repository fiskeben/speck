package blog

import "strings"

func doDryRun(post string) string {
	response := []string{}
	response = append(response, "Dry run mode - your post will not be published.")
	response = append(response, "Here's what your post would be like:")
	response = append(response, "")
	response = append(response, post)
	response = append(response, "Run again without the --dry-run flag to publish your words.")
	return strings.Join(response, "\n")
}
