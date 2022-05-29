package domain

import (
	"bytes"
	"google.golang.org/protobuf/compiler/protogen"
	"log"
	"text/template"
)

type ServiceGen struct {
	module string
}

func NewServiceGen() ServiceGen {
	return ServiceGen{}
}

func (sg ServiceGen) Generate(plugin *protogen.Plugin) error {
	for _, f := range plugin.Files {
		if len(f.Services) == 0 {
			continue
		}

		for _, s := range f.Services {
			log.Println(" f.GeneratedFilenamePrefix:", f.GeneratedFilenamePrefix)
			log.Println(" f.GoImportPath:", f.GoImportPath)
			log.Println(" f.GoPackageName:", f.GoPackageName)
			log.Println(" s.GoName:", s.GoName)

			fileName := f.GeneratedFilenamePrefix + ".impl.go"

			t := plugin.NewGeneratedFile(fileName, f.GoImportPath)

			tmpl, err := template.New("serviceGenImpl").Delims("[[", "]]").Parse(serviceGenImpl)
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
