# go-bitbucket

<a class="repo-badge" href="https://godoc.org/github.com/davidji99/bitbucket-go"><img src="https://godoc.org/github.com/davidji99/bitbucket-go?status.svg" alt="bitbucket-go?status"></a>
<a href="https://goreportcard.com/report/github.com/davidji99/go-bitbucket"><img class="badge" tag="github.com/davidji99/bitbucket-go" src="https://goreportcard.com/badge/github.com/davidji99/bitbucket-go"></a>

> Bitbucket APIv2 library for Golang.

## NOTE: 
This repository was originally a fork of this [repository](https://github.com/ktrysmt/go-bitbucket). I have introduced breaking changes to my original fork which made it impossible to track upstream/parent. So, I have decided it is best to move my fork to a standalone repository. All the history from the fork remains intact. 

I am aiming to implement all available API resources Bitbucket APIv2 provides to its users publically.

This fork also takes several 'inspirations' (with due credit) from the [go-github library](https://github.com/google/go-github) so if you see similarities, you know why =).

## Install

```sh
go get github.com/davidji99/go-bitbucket
```

## Usage

### Sample: 
```go
package main

import (
        "fmt"

        "github.com/davidji99/bitbucket-go/bitbucket"
)

func main() {
        client := bitbucket.NewBasicAuth("<USERNAME>", "<PASSWORD>")
        
        title := "new issue"
        description := "new issue description"
        baseBranch := "master"
        sourceBranch := "bugFix/fix-this-issue"
        closeSourceBranch := true

        createOpts := &bitbucket.CreatePullRequestOpts{
            Title:       &title,
            Description: &description,
            Destination: &bitbucket.NewPullRequestDestinationOpts{
                Branch: &bitbucket.Branch{Name: &baseBranch},
            },
            Source: &bitbucket.NewPullRequestSourceOpts{
                Branch: &bitbucket.Branch{Name: &sourceBranch},
            },
            CloseSourceBranch: &closeSourceBranch,
        }

        newPullRequest, response, createErr := client.PullRequests.Create("<ORG>", "<REPO_SLUG>", createOpts)
        if createErr != nil {
        	panic(createErr)
        }
        
        if response.StatusCode == 201 {
            fmt.Println("Pull request created!")
        }

        fmt.Println(newPullRequest.GetLinks().GetSelf().GetHRef())
}
```

### Query Parameters:
In addition to resource specific query parameters, Bitbucket offers what I like to call 'generic' query parameters that 
are not tied to a specific resource. These query parameters are:
- List `?page=1&pagelen=35`
- [Filter & Sort](https://developer.atlassian.com/bitbucket/api/2/reference/meta/filtering#query-sort) `?q=source.repository.full_name+%21%3D+%22main%2Frepo%22`
- [Partial Response](https://developer.atlassian.com/bitbucket/api/2/reference/meta/partial-response) `?fields=values.id,values.reviewers.username`

Not all of the above will work with every resource/endpoint so please refer to the official [Bitbucket APIv2 documentation](https://developer.atlassian.com/bitbucket/api/2/reference).

Users can mix and match generic parameters with resource parameters or even create their own `struct` for query parameters.

Example usage of query parameters:
```go
api := bitbucket.NewBasicAuth("<USER>", "<APP_PASSWORD>")

opts1 := &bitbucket.PartialRespOpts{
    Fields: "-values.links",
}

opts2 := &bitbucket.FilterSortOpts{
    Query: "destination.branch.name = \"master\"",
}

opts3 := &bitbucket.PullRequestListOpts{
    State: []string{"OPEN"},
}

result, _, err := api.PullRequests.List(c.String("<ORG>", "<REPO_SLUG>", opts1, opts2, opts3)
if err != nil {
    return err
}

for _, i := range result.Values {
    fmt.Println(i.GetPriority())
    fmt.Println(i.GetLinks().GetSelf().GetHRef())
}
```
which will return all pull requests that are `open`, with no links in the results, and whose destination branch in `master`.

## FAQ
- Only supports Bitbucket APIv2.

## Author

[davidji99](https://github.com/davidji99)
