package domain

import (
	"bytes"
	"errors"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"strings"
	"text/template"
)

type ServiceGen struct {
}

func NewServiceGen() ServiceGen {
	return ServiceGen{}
}

func (sg ServiceGen) Generate(plugin *protogen.Plugin) error {
	params := make(map[string]string)
	for _, v := range strings.Split(plugin.Request.GetParameter(), ",") {
		var param string
		var value string
		if i := strings.Index(v, "="); i >= 0 {
			value = param[i+1:]
			param = param[0:i]
		}

		params[param] = value
	}

	if params["module"] == "" {
		return errors.New("invalid module param")
	}

	for _, f := range plugin.Files {
		for _, s := range f.Services {
			fileName := params["module"] + "/internal/delivery/" + string(f.GoPackageName) + "/" + s.GoName + ".impl.go"

			// 已经创建的不管.
			if _, err := os.Stat(fileName); err == nil {
				continue
			}

			t := plugin.NewGeneratedFile(fileName, f.GoImportPath)

			tmpl, err := template.New("serviceGenImpl").Delims("[[", "]]").Parse(serviceGenImpl)
			if err != nil {
				return err
			}

			var tmplResult bytes.Buffer

			if err := tmpl.Execute(&tmplResult, map[string]interface{}{
				"module":      params["module"],
				"packageName": f.GoPackageName,
				"serviceName": s.GoName,
			}); err != nil {
				return err
			}

			t.P(tmplResult.String())
		}
	}

	return nil
}
