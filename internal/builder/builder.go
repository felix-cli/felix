package builder

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/gobuffalo/packr"
)

const (
	defaultTemplateURL = "https://github.com/felix-cli/grpc-service.felix"
)

// Init fetches the default template repo and installs it to a users computer
func Init(template *Template) error {
	tmpDir, err := ioutil.TempDir("", "felix")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	templateURL := template.URL
	if templateURL == "" {
		templateURL = defaultTemplateURL
	}

	cmd := exec.Command("git", "clone", templateURL, tmpDir)
	err = cmd.Run()
	if err != nil {
		return err
	}

	box := packr.NewBox(tmpDir)
	builder := &Builder{
		Box:      box,
		rootDir:  tmpDir,
		template: template,
	}

	err = builder.writeToLocal()
	if err != nil {
		return err
	}

	return nil
}
