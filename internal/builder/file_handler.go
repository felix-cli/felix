package builder

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func copyToTempDir(tmpDir string, reader *zip.ReadCloser) (string, error) {
	var rootDir string
	for _, f := range reader.Reader.File {

		zipped, err := f.Open()
		if err != nil {
			return "", err
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
				return "", err
			}

			defer writer.Close()

			if _, err = io.Copy(writer, zipped); err != nil {
				return "", err
			}
		}
	}

	return rootDir, nil
}

func (b *Builder) writeToLocal(reader *zip.ReadCloser) error {
	for _, f := range reader.Reader.File {
		fpath := f.Name
		if f.FileInfo().IsDir() {
			if fpath == b.rootDir {
				continue
			}

			truePath := strings.Replace(fpath, b.rootDir, "", 1)
			os.MkdirAll(truePath, f.Mode())

			continue
		}

		truePath := strings.Replace(fpath, b.rootDir, "", 1)

		err := b.writeFile(truePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Builder) writeFile(fileName string) error {
	file, err := b.Box.FindString(fileName)
	if err != nil {
		return err
	}

	parsedFile, err := b.parseTemplate(file)
	if err != nil {
		fmt.Println("err parsing: ", err)
	}

	err = ioutil.WriteFile(fileName, parsedFile, 0644)
	if err != nil {
		return err
	}

	return nil
}
