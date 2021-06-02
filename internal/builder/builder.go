package builder

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/gobuffalo/packr"
)

const (
	defaultTemplateURL = "http://github.com/felix-cli/grpc-service.felix/archive/master.zip"
)

// Init fetches the default template repo and installs it to a users computer
func Init(template *Template) error {
	tmpDir, err := ioutil.TempDir("", "felix")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	file, err := ioutil.TempFile(tmpDir, "templates-zip")
	if err != nil {
		return err
	}

	templateURL := template.URL
	if templateURL == "" {
		templateURL = defaultTemplateURL
	}

	cmd := exec.Command("curl", "-L", templateURL, "-o", file.Name())
	cmd.Run()

	reader, err := zip.OpenReader(file.Name())
	if err != nil {
		return err
	}
	defer reader.Close()

	rootDir, err := copyToTempDir(tmpDir, reader)
	if err != nil {
		return err
	}

	box := packr.NewBox(fmt.Sprintf("%s/%s", tmpDir, rootDir))
	builder := &Builder{
		Box:      box,
		rootDir:  rootDir,
		template: template,
	}

	err = builder.writeToLocal(reader)
	if err != nil {
		return err
	}

	if err := os.Remove(file.Name()); err != nil {
		return err
	}

	return nil
}
