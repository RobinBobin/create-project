package eslintconfigexpo

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"regexp"
// 	"strings"

// 	"github.com/Masterminds/semver/v3"
// 	"github.com/charmbracelet/huh"
// 	"github.com/robinbobin/create-project/options"
// 	"github.com/robinbobin/create-project/utils"
// )

// const eslintConfigExpo = "eslint-config-expo"

// func HandleOutdated() {
// 	const eslintPluginReactHooks = "eslint-plugin-react-hooks"

// 	// Get current eslint-plugin-react-hooks version
// 	rawContent, err := os.ReadFile(
// 		filepath.Join(
// 			utils.NODE_MODULES,
// 			eslintPluginReactHooks,
// 			utils.PACKAGE_JSON,
// 		),
// 	)

// 	if errors.Is(err, os.ErrNotExist) {
// 		return
// 	}

// 	utils.PanicOnError(err)

// 	re := regexp.MustCompile(`"version"\s*:\s*"(.*)"`)

// 	matches := re.FindStringSubmatch(string(rawContent))

// 	if matches == nil {
// 		panic(fmt.Errorf("can't determine '%v' version", eslintPluginReactHooks))
// 	}

// 	versionFromPackageJSON := matches[1]

// 	// Get eslint-plugin-react-hooks version from eslint-config-expo package.json
// 	rawContent, err = os.ReadFile(
// 		filepath.Join(
// 			utils.NODE_MODULES,
// 			eslintConfigExpo,
// 			utils.PACKAGE_JSON,
// 		),
// 	)

// 	if errors.Is(err, os.ErrNotExist) {
// 		return
// 	}

// 	utils.PanicOnError(err)

// 	re = regexp.MustCompile(`"eslint-plugin-react-hooks"\s*:\s*"(.*)"`)

// 	matches = re.FindStringSubmatch(string(rawContent))

// 	if matches == nil {
// 		return
// 	}

// 	versionFromDependencies := matches[1]

// 	// Compare with the min compatible version
// 	minVersion := semver.MustParse("7.0.1")

// 	version, err := semver.StrictNewVersion(versionFromDependencies)

// 	if err == nil {
// 		if version.GreaterThanEqual(minVersion) {
// 			return
// 		}
// 	} else if errors.Is(err, semver.ErrInvalidCharacters) {
// 		constraints, err := semver.NewConstraint(versionFromDependencies)

// 		utils.PanicOnError(err)

// 		ok, errs := constraints.Validate(minVersion)

// 		if ok || strings.Contains(errs[0].Error(), "is less than") {
// 			return
// 		}
// 	} else {
// 		utils.PanicOnError(err)
// 	}

// 	fmt.Printf(
// 		"'%v' is using '%v@%v'.\n",
// 		eslintConfigExpo,
// 		eslintPluginReactHooks,
// 		versionFromDependencies,
// 	)

// 	fmt.Printf(
// 		"'%v@%v' is also installed.\n",
// 		eslintPluginReactHooks,
// 		versionFromPackageJSON,
// 	)

// 	fmt.Println("They might not work together.")

// 	actions := []*utils.Action[func()]{
// 		{
// 			Name: "Leave as is.",
// 		},
// 		{
// 			Fn:   commentOut,
// 			Name: fmt.Sprintf("Comment out '%v' usage.", eslintConfigExpo),
// 		},
// 		// {
// 		// 	Fn:   uninstall,
// 		// 	Name: fmt.Sprintf("Uninstall '%v'.", eslintConfigExpo),
// 		// },
// 	}

// 	action := actions[1]

// 	if !options.Options.ShouldUseDefaults {
// 		utils.PanicOnError(
// 			huh.NewSelect[*utils.Action[func()]]().
// 				Title("What would you like to do?").
// 				Options(huh.NewOptions(actions...)...).
// 				Value(&action).
// 				Run(),
// 		)
// 	}

// 	if action.Fn != nil {
// 		action.Fn()
// 	}
// }
