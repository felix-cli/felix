package builder

import (
	"io/ioutil"
	"os"

	"github.com/gobuffalo/packr"
)

// Builder is the struct that holds the box on it with all template files.
type Builder struct {
	Box packr.Box
}

// Init creates new directories with golang files.
func Init() error {
	// set up a new box by giving it a (relative) path to a folder on disk:
	box := packr.NewBox("./templates")
	builder := &Builder{
		Box: box,
	}

	if err := builder.buildMain(); err != nil {
		return err
	}

	if err := os.Mkdir("cmd", os.ModePerm); err != nil {
		return err
	}

	if err := os.Mkdir("internal", os.ModePerm); err != nil {
		return err
	}

	if err := builder.buildFile("go.mod"); err != nil {
		return err
	}

	if err := builder.buildFile("go.sum"); err != nil {
		return err
	}

	if err := os.Mkdir("internal/config", os.ModePerm); err != nil {
		return err
	}
	if err := builder.buildFile("internal/config/config.go"); err != nil {
		return err
	}

	if err := os.Mkdir("internal/handler", os.ModePerm); err != nil {
		return err
	}
	if err := builder.buildFile("internal/handler/handler.go"); err != nil {
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

func (b *Builder) buildMain() error {
	main, err := b.Box.Find("main.go")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("main.go", main, 0644)
	if err != nil {
		return err
	}

	return nil
}
