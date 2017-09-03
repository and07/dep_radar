package dep

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	i "github.com/stamm/dep_radar/interfaces"
)

const (
	file = "Gopkg.lock"
)

var (
	_ i.IDepStrategy = Tool
)

type rawLock struct {
	Projects []rawLockedProject `toml:"projects"`
}

type rawLockedProject struct {
	Name     string   `toml:"name"`
	Branch   string   `toml:"branch,omitempty"`
	Revision string   `toml:"revision"`
	Version  string   `toml:"version,omitempty"`
	Source   string   `toml:"source,omitempty"`
	Packages []string `toml:"packages"`
}

func Tool(a i.IApp) (i.AppDeps, error) {
	res := i.AppDeps{
		Manager: i.DepManager,
	}
	content, err := a.File(file)
	if err != nil {
		return res, err
	}
	if len(content) == 0 {
		return res, fmt.Errorf("File %s is empty", file)
	}
	// fmt.Printf("content = \n%+v\n---------------------\n", string(content))
	raw := rawLock{}
	err = toml.Unmarshal(content, &raw)
	if err != nil {
		return res, errors.Wrap(err, "Unable to parse the lock as TOML")
	}
	// fmt.Printf("lockFile = %+v\n", lockFile.Imports)
	res.Deps = make(map[i.Pkg]i.Dep, len(raw.Projects))
	for _, imp := range raw.Projects {
		// fmt.Printf("imp.Name = %+v\n", imp.Name)
		res.Deps[i.Pkg(imp.Name)] = i.Dep{
			Package: imp.Name,
			Hash:    i.Hash(imp.Revision),
		}
	}
	// fmt.Printf("deps = %+v\n", deps)
	return res, nil
}
