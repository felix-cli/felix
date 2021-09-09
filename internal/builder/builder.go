package builder

import (
	"os"

	"github.com/gobuffalo/packr"
)

const (
	defaultTemplateURL = "https://github.com/felix-cli/grpc-service.felix"
)

// Init fetches the default template repo and installs it to a users computer
func Init(template *Template) error {
	templateURL := template.URL
	if templateURL == "" {
		templateURL = defaultTemplateURL
	}

	srcDir, err := CurateSource(templateURL)
	if err != nil {
		return err
	}
	defer os.RemoveAll(srcDir)

	box := packr.NewBox(srcDir)
	builder := &Builder{
		Box:      box,
		rootDir:  srcDir,
		template: template,
	}

	err = builder.writeToLocal()
	if err != nil {
		return err
	}

	return nil
}
