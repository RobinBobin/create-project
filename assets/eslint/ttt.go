package eslint

// func createConfigs() {
// 	// Get the file list from `assets/eslint`
// 	entries, err := assets.ReadDir(utils.ESLINT)

// 	utils.PanicOnError(err)

// 	configs := []config{}

// 	for _, entry := range entries {
// 		if entry.IsDir() || entry.Name() == utils.ESLINT_CONFIG_TS {
// 			continue
// 		}

// 		configs = append(configs, config{
// 			fileName: entry.Name(),
// 		})
// 	}

// 	if len(configs) == 0 {
// 		return
// 	}

// 	options.Options.AddInclude(fmt.Sprintf("%v/**/*.js", utils.ESLINT))

// 	utils.PanicOnError(os.Mkdir(utils.ESLINT, 0775))

// 	re := regexp.MustCompile(`(?s)import\s+.*?\s+from\s+['"]([^'"]+)['"]`)

// 	// Read eslint.config.js and prepare the content for addition
// 	rawContent, err := os.ReadFile(utils.ESLINT_CONFIG_TS)

// 	utils.PanicOnError(err)

// 	content := string(rawContent)

// 	const separator = "\n"
// 	separatorIndex := strings.LastIndex(content, separator)
// 	separatorIndex = strings.LastIndex(content[:separatorIndex], separator)

// 	content = content[:separatorIndex]

// 	if content[len(content)-1] != ',' {
// 		content += ","
// 	}

// 	content += "\n"

// 	// Overwrite eslint.config.js
// 	file, err := os.Create(utils.ESLINT_CONFIG_TS)

// 	utils.PanicOnError(err)

// 	defer func() {
// 		_ = file.Close()
// 	}()

// 	var importNames []string

// 	for _, config := range configs {
// 		fileName := filepath.Join(utils.ESLINT, config.fileName)

// 		// Copy the file
// 		assets.CopyFile(fileName, fileName)

// 		// Read the file to get the imports
// 		rawData, err := os.ReadFile(fileName)

// 		utils.PanicOnError(err)

// 		data := string(rawData)

// 		// Get package names from imports
// 		matches := re.FindAllStringSubmatch(data, -1)
// 		packages := []string{}

// 		for _, match := range matches {
// 			parts := strings.Split(match[1], "/")

// 			if strings.HasPrefix(parts[0], ".") {
// 				// Relative import
// 				dir := filepath.Join(utils.ESLINT, parts[1])
// 				err := assets.CopyFS(dir, dir)

// 				if !errors.Is(err, os.ErrExist) {
// 					utils.PanicOnError(err)
// 				}

// 				continue
// 			}

// 			var packageName string

// 			if strings.HasPrefix(parts[0], "@") && len(parts) >= 2 {
// 				// Scoped package
// 				packageName = strings.Join(parts[:2], "/")
// 			} else {
// 				// Unscoped package
// 				packageName = parts[0]
// 			}

// 			packages = append(packages, packageName)
// 		}

// 		utils.RunCmd(
// 			fmt.Sprint(
// 				"pnpm install --save-dev ",
// 				strings.Join(packages, " "),
// 			),
// 		)

// 		// Add require
// 		parts := strings.Split(
// 			strings.TrimSuffix(config.fileName, ".js"),
// 			".",
// 		)

// 		for index := 1; index < len(parts); index++ {
// 			parts[index] = strings.ToUpper(parts[index][:1]) + parts[index][1:]
// 		}

// 		importName := fmt.Sprint(strings.Join(parts, ""), "Config")

// 		importNames = append(importNames, importName)

// 		_, err = fmt.Fprintf(file, "const %v = require('./%v')\n", importName, fileName)

// 		utils.PanicOnError(err)
// 	}

// 	_, err = file.WriteString(content)

// 	utils.PanicOnError(err)

// 	for _, importName := range importNames {
// 		_, err := fmt.Fprintf(file, "  %v,\n", importName)

// 		utils.PanicOnError(err)
// 	}

// 	_, err = file.WriteString("]);\n")

// 	utils.PanicOnError(err)
// }
