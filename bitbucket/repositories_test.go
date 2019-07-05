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

	cloneLinks := &RepositoryLinks{
		Clone: []*Link{
			{
				Name: &httpName,
				HRef: &httpHREF,
			},
			{
				Name: &sshName,
				HRef: &sshHREF,
			},
		},
	}

	assert.Equal(t, "www.this.isthe.https.link.com", cloneLinks.GetHTTPSCloneURL())
	assert.Equal(t, "www.this.isthe.ssh.link.com", cloneLinks.GetSSHCloneURL())
}

func TestRepositoryCloneLinks_GetLinks_Missing(t *testing.T) {
	httpName := "https"
	httpHREF := "www.this.isthe.https.link.com"

	cloneLinks := &RepositoryLinks{
		Clone: []*Link{
			{
				Name: &httpName,
				HRef: &httpHREF,
			},
		},
	}

	assert.Equal(t, "www.this.isthe.https.link.com", cloneLinks.GetHTTPSCloneURL())
	assert.Equal(t, "", cloneLinks.GetSSHCloneURL())
}

func TestRepositoryCloneLinks_GetLinks_None(t *testing.T) {
	cloneLinks := &RepositoryLinks{
		Clone: []*Link{},
	}

	assert.Equal(t, "", cloneLinks.GetHTTPSCloneURL())
	assert.Equal(t, "", cloneLinks.GetSSHCloneURL())
}
