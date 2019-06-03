# go-bitbucket

<a class="repo-badge" href="https://godoc.org/github.com/davidji99/go-bitbucket"><img src="https://godoc.org/github.com/davidji99/go-bitbucket?status.svg" alt="go-bitbucket?status"></a>
<a href="https://goreportcard.com/report/github.com/davidji99/go-bitbucket"><img class="badge" tag="github.com/davidji99/go-bitbucket" src="https://goreportcard.com/badge/github.com/davidji99/go-bitbucket"></a>

> Bitbucket APIv2 library for Golang.

## NOTE: 
This repository does not follow its [parent](https://github.com/ktrysmt/go-bitbucket) anymore as I have introduced major breaking changes to this fork. I aim to implement all available API resources Bitbucket APIv2 provides to its users publically.

This fork also takes several 'inspirations' (with due credit) from the [go-github library](https://github.com/google/go-github) so if you see similarities, you know why =).

## Install

```sh
go get github.com/davidji99/go-bitbucket
```

## Usage

```go
package main

import (
        "fmt"

        "github.com/davidji99/go-bitbucket"
)

func main() {

        c := bitbucket.NewBasicAuth("username", "password")

        opt := &bitbucket.PullRequestsOptions{
                Owner:             "your-team",
                RepoSlug:          "awesome-project",
                SourceBranch:      "develop",
                DestinationBranch: "master",
                Title:             "fix bug. #9999",
                CloseSourceBranch: true,
        }

        res, err := c.Repositories.PullRequests.Create(opt)
        if err != nil {
                panic(err)
        }

        fmt.Println(res) 
}
```

## FAQ
- Only supports Bitbucket APIv2.

## Author

[davidji99](https://github.com/davidji99)
