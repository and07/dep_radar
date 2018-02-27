package providers

import (
	"context"
	"testing"

	i "github.com/stamm/dep_radar/src/interfaces"
	"github.com/stamm/dep_radar/src/providers/github"
	"github.com/stretchr/testify/assert"
)

func TestDetect_ExpectGithub(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	detect := DefaultDetector()
	prov, err := detect.Detect(context.Background(), i.Pkg("github.com/golang/dep"))
	assert.NoError(err)
	assert.IsType(&github.Provider{}, prov)
	assert.Implements((*i.IProvider)(nil), prov)
}

func TestDetect_ExpectError(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	detect := DefaultDetector()
	app, err := detect.Detect(context.Background(), i.Pkg("gopkg.in/yaml.v2"))
	assert.EqualError(err, "No provider")
	assert.Nil(app)
}
