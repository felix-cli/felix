package builder

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func CurateSource(path string) (string, error) {
	tmpDir, err := ioutil.TempDir("", "felix")
	if err != nil {
		return "", err
	}

	if IsDirectory(path) {
		err = CopyDirectory(path, tmpDir)
	} else {
		err = GitClone(path, tmpDir)
	}

	if err != nil {
		return "", nil
	}

	return tmpDir, nil
}

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

	return err
}

func GitClone(src, dst string) error {
	cmd := exec.Command("git", "clone", src, dst)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
