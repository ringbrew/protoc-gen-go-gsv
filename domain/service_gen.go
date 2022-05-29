package domain

import (
	"bytes"
	"google.golang.org/protobuf/compiler/protogen"
	"text/template"
)

type ServiceGen struct {
	module string
}

func NewServiceGen(module string) ServiceGen {
	return ServiceGen{
		module: module,
	}
}

func (sg ServiceGen) Generate(plugin *protogen.Plugin) error {
	for _, f := range plugin.Files {
		if len(f.Services) == 0 {
			continue
		}

		for _, s := range f.Services {
			fileName := "./internal/domain/delivery/" + f.GeneratedFilenamePrefix + "." + s.GoName + ".impl.go"

			t := plugin.NewGeneratedFile(fileName, f.GoImportPath)

			tmpl, err := template.New("serviceGenImpl").Parse(serviceGenImpl)
			if err != nil {
				return err
			}

			var tmplResult bytes.Buffer

			if err := tmpl.Execute(&tmplResult, map[string]interface{}{
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