package html

import (
	"bytes"
	"html/template"
	"io/ioutil"

	i "github.com/stamm/dep_radar/interfaces"
	"github.com/stamm/dep_radar/src"
	"github.com/stamm/dep_radar/src/html/templates"
	"github.com/stamm/dep_radar/src/providers"
)

// LibsHTML return html with table. In the head libs, on the left side - apps
func LibsHTML(apps <-chan i.IApp, detector *providers.Detector, rec src.MapRecomended) ([]byte, error) {
	var buf bytes.Buffer
	raw, err := templates.Asset("src/html/templates/libs.html")
	if err != nil {
		return buf.Bytes(), err
	}
	tmpl, err := template.New("libs").Parse(string(raw))
	if err != nil {
		return buf.Bytes(), err
	}
	data := prepare(apps, detector, rec)
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return buf.Bytes(), err
	}
	html, err := ioutil.ReadAll(&buf)
	if err != nil {
		return buf.Bytes(), err
	}
	return html, nil
}
