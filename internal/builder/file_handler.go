package builder

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func (b *Builder) writeToLocal() error {
	var destinationDir string
	if b.template.Name != "" {
		os.MkdirAll(b.template.Name, fs.FileMode(0755))
		destinationDir = fmt.Sprintf("%s/", b.template.Name)
	}

	err := filepath.WalkDir(b.rootDir, func(path string, d fs.DirEntry, _ error) error {
		if path == b.rootDir {
			return nil
		}

		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if fileInfo.Name() == ".git" || strings.Contains(path, ".git/") {
			return nil
		}

		if fileInfo.IsDir() {
			truePath := strings.Replace(path, fmt.Sprintf("%s/", b.rootDir), destinationDir, 1)
			err := os.MkdirAll(truePath, fileInfo.Mode())
			if err != nil {
				return err
			}

			return nil
		}

		truePath := strings.ReplaceAll(path, fmt.Sprintf("%s/", b.rootDir), destinationDir)

		// set felix yaml if it exists
		if len(b.felixYaml) == 0 {
			b.updateTemplateFromFelixYaml()
		}

		err = b.writeFile(path, truePath, fileInfo.Name())
		if err != nil {
			return fmt.Errorf("writing file %s, %w", fileInfo.Name(), err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (b *Builder) writeFile(tmpPath, newPath, fileName string) error {
	var file []byte
	var err error
	file, err = ioutil.ReadFile(tmpPath)
	if err != nil {
		return err
	}

	if len(b.felixYaml) != 0 {
		file, err = b.parseTemplate(fileName, string(file))
		if err != nil {
			return fmt.Errorf("parsing template: %w", err)
		}
	}

	err = ioutil.WriteFile(newPath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}
