package domain

import (
	"bytes"
	"errors"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"regexp"
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
			value = v[i+1:]
			param = v[0:i]
		}

		params[param] = value
	}

	if params["module"] == "" {
		return errors.New("invalid module param")
	}

	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	toSnakeCase := func(str string) string {
		snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
		snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
		return strings.ToLower(snake)
	}

	for _, f := range plugin.Files {
		for _, s := range f.Services {
			var tmplResult bytes.Buffer

			serviceName := s.GoName
			if len(serviceName) > len(string(f.GoPackageName)) && strings.HasPrefix(strings.ToLower(serviceName), strings.ToLower(string(f.GoPackageName))) {
				serviceName = serviceName[len(f.GoPackageName):]
			}

			deliveryFileName := params["module"] + "/internal/delivery/" + string(f.GoPackageName) + "/" + toSnakeCase(s.GoName) + ".impl.go"
			// 已经创建的不管.
			if _, err := os.Stat(deliveryFileName); err == nil {
				continue
			}
			deliveryFile := plugin.NewGeneratedFile(deliveryFileName, f.GoImportPath)
			tmpl, err := template.New("serviceGenImpl").Delims("[[", "]]").Parse(serviceGenImpl)
			if err != nil {
				return err
			}
			if err := tmpl.Execute(&tmplResult, map[string]interface{}{
				"module":           params["module"],
				"packageName":      f.GoPackageName,
				"serviceName":      serviceName,
				"protoServiceName": s.GoName,
			}); err != nil {
				return err
			}
			deliveryFile.P(tmplResult.String())

			serviceDefineFileName := string(f.GoImportPath) + "/" + toSnakeCase(s.GoName) + ".define.go"
			// 已经创建的不管.
			if _, err := os.Stat(deliveryFileName); err == nil {
				continue
			}
			serviceDefineFile := plugin.NewGeneratedFile(serviceDefineFileName, f.GoImportPath)
			defineTmpl, err := template.New("serviceDefineImpl").Delims("[[", "]]").Parse(serviceDefineImpl)
			if err != nil {
				return err
			}
			tmplResult.Reset()
			if err := defineTmpl.Execute(&tmplResult, map[string]interface{}{
				"module":           params["module"],
				"packageName":      f.GoPackageName,
				"serviceName":      serviceName,
				"protoServiceName": s.GoName,
			}); err != nil {
				return err
			}
			serviceDefineFile.P(tmplResult.String())
		}
	}

	return nil
}
