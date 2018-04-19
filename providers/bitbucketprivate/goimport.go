package bitbucketprivate

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/stamm/dep_radar"
	"github.com/stamm/dep_radar/goimport"
)

// GetProject get all projects from bitbucket
func GetProject(ctx context.Context, client dep_radar.IWebClient, pkg dep_radar.Pkg, prefix string) (string, error) {
	prefix = strings.Trim(prefix, "/")
	prefix = regexp.QuoteMeta(prefix)
	re := regexp.MustCompile(prefix + `/([^/]+)`)
	url := string(pkg)
	sources, err := goimport.GetImports(ctx, client, url)
	if err != nil {
		return "", err
	}
	urls := make([]string, 0, len(sources))
	for _, source := range sources {
		if url != source.Prefix {
			continue
		}
		matched := re.FindAllStringSubmatch(source.RepoRoot, -1)
		if len(matched) > 0 {
			return matched[0][1], nil
		}
		urls = append(urls, source.RepoRoot)
	}
	return "", fmt.Errorf("Can't find project in repo urls %s", strings.Join(urls, ", "))
}
