package cli

import "os"

const (
	BitbucketUserEnvVar = "BB_USER"
	BitbucketPassEnvVar = "BB_PASS"
	BitbucketHostEnvVar = "BB_HOST"
)

type BBCli struct {
	User *string
	Pass *string
	Host *string
}

func New() *BBCli {
	user := os.Getenv(BitbucketUserEnvVar)
	pass := os.Getenv(BitbucketPassEnvVar)
	host := os.Getenv(BitbucketHostEnvVar)

	return &BBCli{
		User: &user,
		Pass: &pass,
		Host: &host,
	}
}

// GetUser returns the User field if it's non-nil, empty string otherwise.
func (c *BBCli) GetUser() string {
	if c == nil || c.User == nil {
		return ""
	}
	return *c.User
}

// GetPass returns the Pass field if it's non-nil, empty string otherwise.
func (c *BBCli) GetPass() string {
	if c == nil || c.Pass == nil {
		return ""
	}
	return *c.Pass
}

// GetHost returns the Host field if it's non-nil, empty string otherwise.
func (c *BBCli) GetHost() string {
	if c == nil || c.Host == nil {
		return ""
	}
	return *c.Host
}
