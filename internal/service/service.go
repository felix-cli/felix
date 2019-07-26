package service

import (
	"io"
	"os"
	"path/filepath"
)

// Init creates new directories with golang files.
func Init() error {
	if err := buildMain(); err != nil {
		return err
	}

	if err := os.Mkdir("cmd", os.ModePerm); err != nil {
		return err
	}

	return nil
}

func buildMain() error {
	absPath, _ := filepath.Abs("../felix/internal/service/sample/main.go")
	srcFile, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create("main.go") // creates if file doesn't exist
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
