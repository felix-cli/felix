package builder

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

// Template is the default struct for defining template variables
type Template struct {
	Org  string
	Proj string
}

func (b *Builder) parseTemplate(file string) ([]byte, error) {
	tmpl, err := template.New("felix").Parse(file)
	if err != nil {
		return nil, err
	}

	var newFile bytes.Buffer
	err = tmpl.Execute(&newFile, b.felixYaml)
	if err != nil {
		return nil, err
	}

	return newFile.Bytes(), nil
}

func (b *Builder) updateTemplateFromFelixYaml() {
	felixYaml, err := b.Box.Find("felix.yaml")
	if err != nil {
		return
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(felixYaml), &m)
	if err != nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for k, v := range m {
		fmt.Printf(`%s [%s]: `, k, v)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		if text != "" {
			text = strings.Replace(text, " ", "_", -1)

			m[k] = strings.TrimSuffix(text, "\n")
		}
	}

	b.felixYaml = m

	return
}
