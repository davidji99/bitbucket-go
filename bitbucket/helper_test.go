package bitbucket

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseForResourceId_Valid(t *testing.T) {
	regexExpression := `http[sS]?:\/\/.*\/2.0\/repositories\/.*\/.*\/versions/(\d+)`
	expected := int64(354035)
	result := parseForResourceID(regexExpression, "https://api.bitbucket.org/2.0/repositories/user/repo/versions/354035")

	assert.Equal(t, &expected, result)
}

func TestParseForResourceId_Invalid(t *testing.T) {
	regexExpression := `http[sS]?:\/\/.*\/2.0\/reposi1tories\/.*\/.*\/versions/(\d+)`
	expected := int64(0)
	result := parseForResourceID(regexExpression, "https://api.bitbucket.org/2.0/repositories/user/repo/versions/354035")

	assert.Equal(t, &expected, result)
}
