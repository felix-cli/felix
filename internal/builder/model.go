package builder

import (
	"github.com/gobuffalo/packr"
)

// Builder is the struct that holds the box on it with all template files.
type Builder struct {
	Box       packr.Box
	rootDir   string
	template  *Template
	felixYaml map[interface{}]interface{}
}
