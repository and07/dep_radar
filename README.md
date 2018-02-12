[![Travis](https://img.shields.io/travis/stamm/dep_radar.svg?style=flat-square)](https://travis-ci.org/stamm/dep_radar)
[![Code Climate](https://img.shields.io/codeclimate/github/stamm/dep_radar.svg?style=flat-square)](https://codeclimate.com/github/stamm/dep_radar)
[![Code Climate](https://img.shields.io/codeclimate/coverage/github/stamm/dep_radar.svg?style=flat-square)](https://codeclimate.com/github/stamm/dep_radar/coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/stamm/dep_radar)](https://goreportcard.com/report/github.com/stamm/dep_radar)

## Dep radar
`dep radar` is a prototype to control Go dependencies in microservice world.
`dep radar` is not stable yet. It requires Go 1.9 or newer to compile.

# Screenshots
![Application screenshot](https://github.com/stamm/dep_radar/raw/master/docs/apps.png)

![Libraries screenshot](https://github.com/stamm/dep_radar/raw/master/docs/libs.png)

## How it works
You can't just run some binary. You have to write a bit of code.
Your code must implement:
* Get a list of packages of applications what dependencies you want to monitor.
* Init provider detector. It can be a default with support only github, but you can add you own provider.
* A http handler with calling method for generate html table with all apps and dependencies.

Simple example for create a table with dependencies for whole github organization [dep-radar](https://github.com/dep-radar):

```go
package main

import (
	"os"

	"github.com/stamm/dep_radar/src/helpers"
	"github.com/stamm/dep_radar/src"
)

func main() {
	recom := src.MapRecomended{
		"github.com/pkg/errors": src.Option{
			Recomended: ">=0.8.0",
			Mandatory:  true,
			Exclude: false,
			NeedVersion: true,
		},
	}
	helpers.GithubOrg(os.Getenv("GITHUB_TOKEN"), "dep-radar", ":8081", recom)
}

```
You can find more [examples](examples/)



## Supported code storage
* Github (env: GITHUB_TOKEN)
* Private Bitbucket (env: BB_GIT_URL, BB_GO_GET_URL, BB_USER, BB_PASSWORD)

## Supported dep tools
* [dep](https://github.com/golang/dep)
* [glide](https://github.com/Masterminds/glide)
