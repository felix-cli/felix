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

	fmt.Println("read items from tmp dir")
	filepath.WalkDir(b.rootDir, func(path string, d fs.DirEntry, _ error) error {
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
			fmt.Println("what is this? ", truePath)
			err := os.MkdirAll(truePath, fileInfo.Mode())
			if err != nil {
				fmt.Println("error making dir: ", err.Error())
				return err
			}

			return nil
		}

		truePath := strings.ReplaceAll(path, fmt.Sprintf("%s/", b.rootDir), destinationDir)

		if len(b.felixYaml) == 0 {
			b.updateTemplateFromFelixYaml()
		}

		b.writeFile(path, truePath)

		return nil
	})

	return nil
}

func (b *Builder) writeFile(tmpPath, newPath string) error {
	file, err := ioutil.ReadFile(tmpPath)
	if err != nil {
		fmt.Println("err reading file ", tmpPath)
		return err
	}

	parsedFile, err := b.parseTemplate(string(file))
	if err != nil {
		fmt.Println("err parsing: ", err)
	}

	fmt.Println("writing file to path", newPath)
	err = ioutil.WriteFile(newPath, parsedFile, 0644)
	if err != nil {
		fmt.Println("err writing file: ", err.Error())
		return err
	}

	return nil
}

// func (b *Builder) writeFile(destinationDir, fileName string) error {
// 	file, err := b.Box.FindString(fileName)
// 	if err != nil {
// 		return err
// 	}

// parsedFile, err := b.parseTemplate(file)
// if err != nil {
// 	fmt.Println("err parsing: ", err)
// }

// newFile := filepath.Join(destinationDir, fileName)
// err = ioutil.WriteFile(newFile, parsedFile, 0644)
// if err != nil {
// 	return err
// }

// 	return nil
// }
