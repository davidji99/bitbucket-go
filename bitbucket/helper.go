package bitbucket

import (
	"regexp"
	"strconv"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// parseForResourceID takes a resource's self link and appropriate regex to parse out the resource ID.
// This is needed because there are few resource APIs that don't return the ID.
// Returns an int64 value for the resource's ID.
func parseForResourceID(regexExpression, selfLink string) *int64 {
	id := int64(0)
	regex := regexp.MustCompile(regexExpression)
	result := regex.FindStringSubmatch(selfLink)

	if len(result) == 2 {
		integer, _ := strconv.Atoi(result[1])
		id = int64(integer)
	}

	return &id
}

// toSnakeCase takes a string and converts it to snake case.
func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
