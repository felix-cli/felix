package builder

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

// CopyDirectory copies all files in the src directory to the dst directory.
func CopyDirectory(src, dst string) error {
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if path == src {
			return nil
		}

		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		dstPath := strings.Replace(path, src, dst, 1)

		switch info.Mode() & os.ModeType {
		case os.ModeDir:
			err := os.MkdirAll(dstPath, info.Mode())
			if err != nil {
				return err
			}

		case os.ModeSymlink:
			// TODO: handle symlinks

		default:
			file, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(dstPath, file, info.Mode())
			if err != nil {
				return err
			}
		}

		return nil
	})

	return errors.Wrap(err, "cloning git repo")
}

// WriteTemplates will read each file in src into a template, execute the template with
// the variables specified in fields, and output the files to the dst directory.
// Neither src nor dst should be empty strings.
func WriteTemplates(src, dst string, fields map[interface{}]interface{}) error {
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if path == src {
			return nil
		}

		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		dstPath := strings.Replace(path, src, dst, 1)

		switch info.Mode() & os.ModeType {
		case os.ModeDir:
			err := os.MkdirAll(dstPath, info.Mode())
			if err != nil {
				return err
			}

		case os.ModeSymlink:
			// TODO: handle symlinks

		default:
			tmpl, err := template.ParseFiles(path)
			if err != nil {
				return err
			}

			newFile, err := os.Create(dstPath)
			if err != nil {
				return err
			}

			newFile.Chmod(info.Mode())

			err = tmpl.Execute(newFile, fields)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
