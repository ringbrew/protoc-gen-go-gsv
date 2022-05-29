package domain

import (
	"bytes"
	"log"
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
