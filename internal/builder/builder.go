package builder

import (
	"io/fs"
	"os"

	"github.com/pkg/errors"
)

// Builder is the struct that holds the box on it with all template files.
type Builder struct {
	Source      string
	Destination string
}

func New(src, dst string) *Builder {
	return &Builder{
		Source:      src,
		Destination: dst,
	}
}

func (b *Builder) Fixit() error {
	dstDir := b.Destination

	// If the destination is blank, write the template to the current directory.
	if dstDir == "" {
		var err error

		dstDir, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	// Create the destination directory if it does not already exist.
	if !IsDirectory(dstDir) {
		err := os.MkdirAll(dstDir, fs.FileMode(0755))
		if err != nil {
			return errors.Wrap(err, "creating output directory")
		}
	}

	srcDir, err := CurateSource(b.Source)
	if err != nil {
		return err
	}
	defer os.RemoveAll(srcDir)

	fields, err := ParseFelixConfig(srcDir)
	if err != nil {
		return errors.Wrap(err, "parsing felix config")
	}

	if fields == nil {
		err = CopyDirectory(srcDir, dstDir)

		return errors.Wrap(err, "copying to dst")
	}

	fields, err = PromptForMissingConfigValues(fields)
	if err != nil {
		return errors.Wrap(err, "reading config values from user")
	}

	err = WriteTemplates(srcDir, dstDir, fields)
	return errors.Wrap(err, "writing template to dst")
}
