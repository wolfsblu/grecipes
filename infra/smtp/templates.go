package smtp

import (
	"bytes"
	"fmt"
	"html/template"
)

type PasswordResetTemplate struct {
	ResetLink string
}

func buildTemplate(path string, data any) (string, error) {
	t := template.New(path)
	t, err := t.ParseFS(templateFS, fmt.Sprintf("templates/%s", path))
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.Execute(&tpl, data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}
