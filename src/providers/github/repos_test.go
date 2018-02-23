package github

import (
	"context"
	"errors"
	"testing"

	"github.com/stamm/dep_radar/interfaces/mocks"
	"github.com/stretchr/testify/require"
)

// Repos

func TestGithubRepos_Ok(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	mHttpClient := &mocks.IWebClient{}
	mHttpClient.On("Get", "https://api.github.com/orgs/dep-radar/repos").Return([]byte(`[
  {
    "full_name": "dep-radar/test_app"
  },
  {
    "full_name": "dep-radar/test_app2"
  }
]`), nil)
	tagsGetter := New(mHttpClient)

	pkgs, err := tagsGetter.GetAllOrgRepos(context.Background(), "dep-radar")
	require.NoError(err)
	require.Len(pkgs, 2, "Expect 2 repos in org")
	require.EqualValues("github.com/dep-radar/test_app", pkgs[0])
	require.EqualValues("github.com/dep-radar/test_app2", pkgs[1])
}

func TestGithubRepos_Error(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	mHttpClient := &mocks.IWebClient{}
	mHttpClient.On("Get", "https://api.github.com/orgs/dep-radar/repos").Return([]byte(``), errors.New("aaa"))
	tagsGetter := New(mHttpClient)

	pkgs, err := tagsGetter.GetAllOrgRepos(context.Background(), "dep-radar")
	require.Error(err)
	require.Len(pkgs, 0, "Expect 0 repos in org")
}
