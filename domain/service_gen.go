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

func NewServiceGen(module string) ServiceGen {
	return ServiceGen{
		module: module,
	}
}

func (sg ServiceGen) Generate(plugin *protogen.Plugin) error {
	for k, v := range GetParamSet().data {
		log.Println("k:", k, " v:", v)
	}

	for _, f := range plugin.Files {
		if len(f.Services) == 0 {
			continue
		}

		for _, s := range f.Services {
			log.Println(" f.module:", sg.module)
			log.Println(" f.GeneratedFilenamePrefix:", f.GeneratedFilenamePrefix)
			log.Println(" f.GoImportPath:", f.GoImportPath)
			log.Println(" f.GoPackageName:", f.GoPackageName)
			log.Println(" s.GoName:", s.GoName)

			fileName := sg.module + "internal/delivery/" + string(f.GoPackageName) + "/" + s.GoName + ".impl.go"

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
