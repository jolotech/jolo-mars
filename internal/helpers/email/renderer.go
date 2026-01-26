package email


import (
	"bytes"
	"html/template"
	"path/filepath"
)

func renderTemplate(templateName string, data any) (string, error) {
	path := filepath.Join("templates", templateName)

	

	tpl, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
