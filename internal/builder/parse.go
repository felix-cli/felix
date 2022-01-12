package builder

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	FelixConfigNames = []string{"felix.yaml", "felix.yml"}
)

// ParseFelixConfig checks the directory for a felix.yaml file and parses it for use
// in a Go template.
func ParseFelixConfig(srcDir string) (map[interface{}]interface{}, error) {
	for _, name := range FelixConfigNames {
		n := path.Join(srcDir, name)
		configFile, err := os.Open(n)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}

			return nil, err
		}
		defer configFile.Close()

		d := yaml.NewDecoder(configFile)

		fields := make(map[interface{}]interface{})
		err = d.Decode(fields)
		if err != nil {
			return nil, err
		}

		return fields, nil
	}

	return nil, nil
}

// PromptForMissingConfigValues reads any missing config values from stdin.
func PromptForMissingConfigValues(fields map[interface{}]interface{}) (map[interface{}]interface{}, error) {
	reader := bufio.NewReader(os.Stdin)
	for k, v := range fields {
		fmt.Printf(`%s [%s]: `, k, v)
		text, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		text = strings.TrimSuffix(text, "\n")
		if text != "" {
			text = strings.Replace(text, " ", "_", -1)

			fields[k] = strings.TrimSuffix(text, "\n")
		}
	}

	return fields, nil
}
