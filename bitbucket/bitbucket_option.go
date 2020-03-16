package bitbucket

import "fmt"

// Option is a functional option for configuring the API http.
type Option func(*Client) error

// baseURL allows overriding of the default base API URL.
func BaseURL(baseURL string) Option {
	return func(c *Client) error {
		// Validate that there is no trailing slashes before setting the custom baseURL
		if baseURL[len(baseURL)-1:] == "/" {
			return fmt.Errorf("custom base URL cannot contain a trailing slash")
		}

		c.baseURL = baseURL
		return nil
	}
}