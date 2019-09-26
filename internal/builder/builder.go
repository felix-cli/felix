package builder

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	defaultTemplateURL = "http://github.com/scottcrawford03/grpc-service.felix/archive/master.zip"
)

// Init fetches the default template repo and installs it to a users computer
func Init() error {
	tmpDir, err := ioutil.TempDir("", "felix")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	file, err := ioutil.TempFile(tmpDir, "templates-zip")
	if err != nil {
		return err
	}

	cmd := exec.Command("curl", "-L", defaultTemplateURL, "-o", file.Name())
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

	err = writeToLocal(tmpDir, rootDir, reader)
	if err != nil {
		return err
	}

	if err := os.Remove(file.Name()); err != nil {
		return err
	}

	return nil
}
