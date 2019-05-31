package bitbucket

import (
	"regexp"
	"strconv"
)

// parseForResourceId takes a resource's self link and appropriate regex to parse out the resource ID.
// This is needed because there are few resource APIs that don't return the ID.
// Returns an int64 value for the resource's ID.
func parseForResourceId(regexExpression, selfLink string) *int64 {
	id := int64(0)
	regex := regexp.MustCompile(regexExpression)
	result := regex.FindStringSubmatch(selfLink)

	if len(result) == 2 {
		integer, _ := strconv.Atoi(result[1])
		id = int64(integer)
	}

	return &id
}
