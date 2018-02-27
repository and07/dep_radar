package html

import (
	"context"
	"sort"
	"strings"

	"github.com/stamm/dep_radar/src"
	"github.com/stamm/dep_radar/src/fill"
	i "github.com/stamm/dep_radar/src/interfaces"
	"github.com/stamm/dep_radar/src/providers"
	versionpkg "github.com/stamm/dep_radar/src/version"
)

const (
	classMandatoryOk  = "table-success"
	classMandatoryBad = "bg-danger"
	classExclude      = "table-warning"
	classNeedVersion  = "table-warning"
)

type templateStruct struct {
	LibView []libView
	AppView []appView
}

type libView struct {
	Name  string
	Hints string
}

type appView struct {
	Name string
	Libs map[string]appLibView
}

type appLibView struct {
	Ok      bool
	Class   string
	Title   string
	Version string
}

// AppsHTML return html with table. In the head apps, on the left side - libs
func prepare(ctx context.Context, apps <-chan i.IApp, detector *providers.Detector, rec src.MapRecommended) templateStruct {
	appsMap, libsMap := fill.GetTags(ctx, apps, detector)
	appsMap = fill.AddVersionLibToApp(appsMap, libsMap)

	// Sort libraries
	pkgsKey := make(sort.StringSlice, 0, len(libsMap))
	for pkgKey := range libsMap {
		pkgsKey = append(pkgsKey, string(pkgKey))
	}
	for pkg := range rec {
		if _, ok := libsMap[pkg]; !ok {
			pkgsKey = append(pkgsKey, string(pkg))
		}
	}
	sort.Strings(pkgsKey)

	// Sort apps
	appsKey := make(sort.StringSlice, 0, len(apps))
	for appPkg := range appsMap {
		appsKey = append(appsKey, string(appPkg))
	}
	sort.Strings(appsKey)

	// Prepair data for render
	pkgsViewData := make([]libView, 0, len(pkgsKey))
	appViewData := make([]appView, 0, len(appsKey))
	for _, pkgName := range pkgsKey {
		var hints []string
		recLib, ok := rec[i.Pkg(pkgName)]
		if ok {
			if recLib.Recommended != "" {
				hints = append(hints, recLib.Recommended)
			}
			if recLib.Mandatory {
				hints = append(hints, "mandatory")
			}
			if recLib.Exclude {
				hints = append(hints, "exclude")
			}
			if recLib.NeedVersion {
				hints = append(hints, "only version")
			}
		}
		pkgsViewData = append(pkgsViewData, libView{
			Name:  pkgName,
			Hints: strings.Join(hints, ", "),
		})
	}

	for _, appKey := range appsKey {
		appData := appView{
			Name: appKey,
			Libs: make(map[string]appLibView, len(pkgsViewData)),
		}
		for _, pkgData := range pkgsViewData {
			libs := appsMap[i.Pkg(appKey)]
			libPkg := i.Pkg(pkgData.Name)
			opt, okOpt := rec[libPkg]

			dep, ok := libs[libPkg]
			if !ok {
				// If library not in app
				appLibData := appLibView{
					Version: "—",
					Ok:      true,
				}
				if okOpt {
					if opt.Mandatory {
						appLibData.Class = classMandatoryBad
						appLibData.Ok = false
						appLibData.Title = "Mandatory to use"
					}
					if opt.Exclude {
						appLibData.Class = classMandatoryOk
					}
				}
				appData.Libs[pkgData.Name] = appLibData
				continue
			}

			appLibData := appLibView{
				Version: dep.Version,
				Ok:      true,
			}
			if appLibData.Version == "" {
				appLibData.Version = string(libs[libPkg].Hash)[0:8]
			}
			if okOpt {
				if opt.Exclude {
					appLibData.Class = classExclude
					appLibData.Ok = false
					appLibData.Title = "Need to exclude this library"
				} else {
					goodVersion, _ := versionpkg.Compare(opt.Recommended, dep.Version)
					if goodVersion {
						appLibData.Class = classMandatoryOk
					} else {
						appLibData.Class = classMandatoryBad
						appLibData.Ok = false
						appLibData.Title = "Need to change version to " + opt.Recommended
					}
				}
				if opt.NeedVersion && dep.Version == "" {
					appLibData.Class = classNeedVersion
					appLibData.Ok = false
					appLibData.Title = "Use version instead of revision"
				}
			}
			appData.Libs[pkgData.Name] = appLibData
		}
		appViewData = append(appViewData, appData)
	}

	return templateStruct{
		LibView: pkgsViewData,
		AppView: appViewData,
	}
}
