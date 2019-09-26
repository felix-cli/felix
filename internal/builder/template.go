package builder

import (
	"bytes"
	"text/template"
)

// Template is the default struct for defining template variables
type Template struct {
	Org  string
	Proj string
}

func (b *Builder) parseTemplate(file string) ([]byte, error) {
	tmp := Template{
		Org:  "scottcrawford03",
		Proj: "test",
	}

	tmpl, err := template.New("felix").Parse(file)
	if err != nil {
		return nil, err
	}

	var newFile bytes.Buffer
	err = tmpl.Execute(&newFile, tmp)
	if err != nil {
		return nil, err
	}

	return newFile.Bytes(), nil
}
