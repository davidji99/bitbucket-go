package bitbucket

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddQueryParams_NoOpts(t *testing.T) {
	urlStr := fmt.Sprintf("/repositories/%s/%s", "bOrg", "bRepo")
	urlStr, addOptErr := addQueryParams(urlStr)

	assert.Nil(t, addOptErr)
	assert.Equal(t, "/repositories/bOrg/bRepo", urlStr)
}

func TestAddQueryParams_OneOpt(t *testing.T) {
	opt := &ListOpts{
		Page:    int64(2),
		Pagelen: int64(5),
	}
	urlStr := fmt.Sprintf("/repositories/%s/%s", "bOrg", "bRepo")
	urlStr, addOptErr := addQueryParams(urlStr, opt)

	assert.Nil(t, addOptErr)
	assert.Equal(t, "/repositories/bOrg/bRepo?page=2&pagelen=5", urlStr)
}

func TestAddQueryParams_MultipleOpts(t *testing.T) {
	opt1 := &ListOpts{
		Page:    int64(2),
		Pagelen: int64(5),
	}

	opt2 := &FilterSortOpts{
		Query: "source.repository.full_name != \"main/repo\" AND state = \"OPEN\"",
		Sort:  "updated_on",
	}

	urlStr := fmt.Sprintf("/repositories/%s/%s", "bOrg", "bRepo")
	urlStr, addOptErr := addQueryParams(urlStr, opt1, opt2)

	assert.Nil(t, addOptErr)
	assert.Equal(t, "/repositories/bOrg/bRepo?page=2&pagelen=5&q=source.repository.full_name+%21%3D+%22main%2Frepo%22+AND+state+%3D+%22OPEN%22&sort=updated_on", urlStr)
}
