package builder

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr"
)

// Builder is the struct that holds the box on it with all template files.
type Builder struct {
	Box packr.Box
}

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

	var rootDir string
	for _, f := range reader.Reader.File {

		zipped, err := f.Open()
		if err != nil {
			return err
		}

		defer zipped.Close()

		// get the individual file name and extract the current directory
		path := filepath.Join(tmpDir, f.Name)
		if f.FileInfo().IsDir() {
			if len(strings.Split(f.Name, string(os.PathSeparator))) == 2 {
				rootDir = f.Name
			}
			os.MkdirAll(path, f.Mode())
		} else {
			writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, f.Mode())
			if err != nil {
				return err
			}

			defer writer.Close()

			if _, err = io.Copy(writer, zipped); err != nil {
				return err
			}
		}
	}

	box := packr.NewBox(fmt.Sprintf("%s/%s", tmpDir, rootDir))
	builder := &Builder{
		Box: box,
	}

	for _, f := range reader.Reader.File {
		fpath := f.Name
		if f.FileInfo().IsDir() {
			if fpath == rootDir {
				continue
			}

			truePath := strings.Replace(fpath, rootDir, "", 1)
			os.MkdirAll(truePath, f.Mode())

			continue
		}

		truePath := strings.Replace(fpath, rootDir, "", 1)
		err = builder.buildFile(truePath)
		if err != nil {
			return err
		}
	}

	if err := os.Remove(file.Name()); err != nil {
		return err
	}

	return nil
}

func (b *Builder) buildFile(fileName string) error {
	file, err := b.Box.Find(fileName)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
