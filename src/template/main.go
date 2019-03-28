package template

import (
	"bytes"
	"html/template"
)

// EmailData is the superset of allowed data in any of the passed templates
type EmailData struct {
	Username string
}

const bucket = "go-mail-template"

// Render Renders a template based on a passed name
// We accept data that is a string string map.  It is expected that you cast to a string on the way in
func Render(templateString string, data map[string]string) (string, error) {
	t, err := template.New("mailer").Parse(templateString)

	var buf bytes.Buffer
	err = t.Execute(&buf, data)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
