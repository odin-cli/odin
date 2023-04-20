package utils

import (
	"bytes"
	"text/template"
)

func GetTemplateValues(str string, params map[string]string) string {
	var result bytes.Buffer
	tpl, _ := template.New("").Parse(str)
	tpl.Execute(&result, &params)
	return result.String()
}
