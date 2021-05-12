package builder_test

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/felix-cli/felix/internal/builder"
)

var update = flag.Bool("update", false, "update .golden files")

func CompareDir(t *testing.T, expDir, testDir string) {
	t.Helper()

	expFiles, err := ioutil.ReadDir(expDir)
	require.NoError(t, err, "expFiles")

	testFiles, err := ioutil.ReadDir(testDir)
	require.NoError(t, err, "testFiles")

	assert.Len(t, testFiles, len(expFiles))
	for i, testFile := range testFiles {
		expFile := expFiles[i]

		expFileName := fmt.Sprintf("%s/%s", expDir, expFile.Name())
		testFileName := fmt.Sprintf("%s/%s", testDir, testFile.Name())

		if testFile.IsDir() {
			CompareDir(t, expFileName, testFileName)
		} else {
			expContent, err := ioutil.ReadFile(expFileName)
			require.NoError(t, err, "expContent")

			testContent, err := ioutil.ReadFile(testFileName)
			require.NoError(t, err, "testContent")

			assert.Equal(t, expContent, testContent)
		}
	}
}

func Test(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "GRPC Template",
			expected: "grpc-service.felix",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()

			curDir, err := os.Getwd()
			require.NoError(t, err)

			outDir, err := ioutil.TempDir("", "felix")
			require.NoError(t, err)
			defer os.RemoveAll(outDir)

			err = os.Chdir(outDir)
			require.NoError(t, err)
			defer func() {
				err = os.Chdir(curDir)
				require.NoError(t, err)
			}()

			tmp := builder.Template{
				Org:  "update",
				Proj: "update_me",
			}

			err = builder.Init(&tmp)
			require.NoError(t, err)

			expDir := fmt.Sprintf("%s/testdata/%s.golden", curDir, tc.expected)

			if *update {
				tempDir := fmt.Sprintf("%s.temp", expDir)
				err = os.Rename(expDir, tempDir)
				require.NoError(t, err)

				err = os.Rename(outDir, expDir)
				if err != nil {
					err = os.Rename(tempDir, expDir)
					require.NoError(t, err)
				}

				err = os.RemoveAll(tempDir)
				require.NoError(t, err)
			} else {
				CompareDir(t, expDir, outDir)
			}
		})
	}
}
