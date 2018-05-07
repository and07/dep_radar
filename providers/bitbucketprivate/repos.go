package bitbucketprivate

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stamm/dep_radar"
)

type reposResponse struct {
	IsLastPage    bool        `json:"isLastPage"`
	Values        []RepoValue `json:"values"`
	NextPageStart int         `json:"nextPageStart"`
}

// RepoValue sub response from bitbucket for repo
type RepoValue struct {
	Slug string `json:"slug"`
}

// GetAllRepos get repos
func (p *Provider) GetAllRepos(ctx context.Context, project string) ([]dep_radar.Pkg, error) {
	var (
		resultRepos []dep_radar.Pkg
		start       int
		isLastPage  bool
	)

	for !isLastPage {
		repos, err := p.getRepos(ctx, project, start)
		if err != nil {
			return resultRepos, err
		}
		for _, repo := range repos.Values {
			resultRepos = append(resultRepos, dep_radar.Pkg(p.goGetURL+"/"+repo.Slug))
		}
		isLastPage = repos.IsLastPage
		start = repos.NextPageStart
	}
	return resultRepos, nil
}

func (p *Provider) getRepos(ctx context.Context, project string, start int) (reposResponse, error) {
	var repos reposResponse
	url := p.getReposURL(p.gitDomain, project, start)
	reposResponse, err := p.httpClient.Get(ctx, url)
	if err != nil {
		return repos, err
	}
	err = json.Unmarshal(reposResponse, &repos)

	return repos, err
}

func (p *Provider) getReposURL(domain, project string, start int) string {
	return fmt.Sprintf("https://%s/rest/api/1.0/projects/%s/repos?start=%d", domain, project, start)
}
