package builder

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

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
		return "", err
	}

	return tmpDir, nil
}

func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func GitClone(src, dst string) error {
	cmd := exec.Command("git", "clone", src, dst)
	err := cmd.Run()
	if err != nil {
		return errors.Wrap(err, "cloning git repo")
	}

	return nil
}
