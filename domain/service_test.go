package domain

import (
	"bytes"
	"google.golang.org/protobuf/compiler/protogen"
	"log"
	"regexp"
	"strings"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	tmpl, err := template.New("serviceGenImpl").Delims("[[", "]]").Parse(serviceGenImpl)
	if err != nil {
		t.Error(err)
		return
	}

	var tmplResult bytes.Buffer

	if err := tmpl.Execute(&tmplResult, map[string]interface{}{
		"packageName": "packageDemo",
		"serviceName": "serviceDemo",
	}); err != nil {
		t.Error(err)
		return
	}

	log.Println(tmplResult.String())
}

func TestGen(t *testing.T) {
	s := NewServiceGen()
	if err := s.Generate(&protogen.Plugin{
		Files: []*protogen.File{
			{
				GoImportPath:            protogen.GoImportPath("test/abc"),
				GeneratedFilenamePrefix: "test",
				GoPackageName:           "testPackageName",
				Services: []*protogen.Service{
					{
						GoName: "testGoName",
					},
				},
			},
		},
	}); err != nil {
		t.Error(err)
		return
	}
}

func TestParam(t *testing.T) {
	params := make(map[string]string)
	for _, v := range strings.Split("module=demo", ",") {
		var param string
		var value string
		if i := strings.Index(v, "="); i >= 0 {
			value = v[i+1:]
			param = v[0:i]
		}

		params[param] = value
	}

	log.Println(params)
}

func TestCamelToSnake(t *testing.T) {
	packageName := "example"
	serviceName := "ExampleXXXXService"

	if len(serviceName) > len(packageName) && strings.HasPrefix(strings.ToLower(serviceName), strings.ToLower(packageName)) {
		serviceName = serviceName[len(packageName):]
	}

	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	toSnakeCase := func(str string) string {
		snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
		snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
		return strings.ToLower(snake)
	}

	log.Println(serviceName)
	log.Println(toSnakeCase(serviceName))
}
