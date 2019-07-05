package bitbucket

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepositoryCloneLinks_GetLinks(t *testing.T) {
	httpName := "https"
	httpHREF := "www.this.isthe.https.link.com"
	sshName := "ssh"
	sshHREF := "www.this.isthe.ssh.link.com"

	cloneLinks := &RepositoryCloneLinks{Values: []*Link{
		{
			Name: &httpName,
			HRef: &httpHREF,
		},
		{
			Name: &sshName,
			HRef: &sshHREF,
		},
	}}

	assert.Equal(t, "www.this.isthe.https.link.com", cloneLinks.GetHTTPS())
	assert.Equal(t, "www.this.isthe.ssh.link.com", cloneLinks.GetSSH())
}

func TestRepositoryCloneLinks_GetLinks_Missing(t *testing.T) {
	httpName := "https"
	httpHREF := "www.this.isthe.https.link.com"

	cloneLinks := &RepositoryCloneLinks{Values: []*Link{
		{
			Name: &httpName,
			HRef: &httpHREF,
		},
	}}

	assert.Equal(t, "www.this.isthe.https.link.com", cloneLinks.GetHTTPS())
	assert.Equal(t, "", cloneLinks.GetSSH())
}
