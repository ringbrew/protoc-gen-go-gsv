package domain

import (
	"bytes"
	"google.golang.org/protobuf/compiler/protogen"
	"log"
	"text/template"
)

type ServiceGen struct {
}

func NewServiceGen() ServiceGen {
	return ServiceGen{}
}

func (sg ServiceGen) Generate(plugin *protogen.Plugin) error {
	ps := GetParamSet()
	for k, v := range ps.data {
		log.Println("k:", k, " v:", v)
	}

	for _, f := range plugin.Files {
		if len(f.Services) == 0 {
			continue
		}

		for _, s := range f.Services {
			log.Println(" f.module:", ps.Get("module"))
			log.Println(" f.GeneratedFilenamePrefix:", f.GeneratedFilenamePrefix)
			log.Println(" f.GoImportPath:", f.GoImportPath)
			log.Println(" f.GoPackageName:", f.GoPackageName)
			log.Println(" s.GoName:", s.GoName)

			fileName := ps.Get("module") + "internal/delivery/" + string(f.GoPackageName) + "/" + s.GoName + ".impl.go"

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
