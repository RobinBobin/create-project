package eslint

import (
	"bytes"
	"maps"
	"path"
	"regexp"
	"slices"
	"strings"

	"github.com/robinbobin/create-project/assets"
	"github.com/robinbobin/create-project/utils"
)

func getConfigData() (configNames []string, dependencies []string) {
	entries, err := assets.ReadDir(utils.ESLINT)

	utils.PanicOnError(err)

	uniqueImports := make(map[string]struct{})
	re := regexp.MustCompile(`from\s+'([^']+)'`)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		configNames = append(configNames, strings.TrimSuffix(entry.Name(), ".ts"))

		buf := assets.ReadFile(path.Join(utils.ESLINT, entry.Name()))

		matches := re.FindAllSubmatch(buf, -1)

		slash := []byte(`/`)

		for _, match := range matches {
			if len(match) <= 1 {
				continue
			}

			dependency := match[1]
			firstByte := dependency[0]

			if firstByte == '.' {
				continue
			}

			if firstByte == '@' {
				parts := bytes.SplitN(dependency, slash, 3)
				dependency = bytes.Join(parts[:2], slash)
			} else {
				dependency, _, _ = bytes.Cut(dependency, slash)
			}

			uniqueImports[string(dependency)] = struct{}{}
		}
	}

	dependencies = slices.Collect(maps.Keys(uniqueImports))

	return
}
