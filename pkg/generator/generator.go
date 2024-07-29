package generator

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/genesysflow/iconify/pkg/api/iconify"
	"github.com/iancoleman/strcase"
)

type IconifyGenerator struct {
	API string
	CWD string
}

// walk through the collections and generate the icons
func Generate(api string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	generator := IconifyGenerator{
		API: api,
		CWD: cwd,
	}

	collections, err := iconify.GetCollections(api)
	if err != nil {
		return
	}

	for _, collection := range collections {
		iconCollection, err := iconify.GetIconCollection(api, collection.Key)
		if err != nil {
			return
		}

		for key, value := range iconCollection.Categories {
			generator.generateFromCategory(key, iconCollection.Prefix, value)
		}

		generator.generateFromCategory("Uncategorized", iconCollection.Prefix, iconCollection.Uncategorized)
	}
}

// generate the icons from a category
func (i *IconifyGenerator) generateFromCategory(category, prefix string, icons []string) {
	filePath, err := i.generateIconPackage(prefix, category)
	if err != nil {
		log.Fatal("Error generating icon package: ", err)
		return
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening package file: ", err)
		return
	}
	defer file.Close()

	for _, icon := range icons {
		body, err := iconify.GetIcon(i.API, prefix, icon)
		if err != nil {
			return
		}
		i.generateIcon(file, icon, body)
	}
}

var iconFunctionTemplate = ""

// add an icon function to the file
func (i *IconifyGenerator) generateIcon(file *os.File, icon, body string) {
	functionName := getIconFunctionName(icon)
	if iconFunctionTemplate == "" {
		data, err := os.ReadFile(i.CWD + "/pkg/generator/icon_function.tpl.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		iconFunctionTemplate = string(data)
	}
	iconFunction := strings.ReplaceAll(iconFunctionTemplate, "$FUNCION-NAME$", functionName)
	iconFunction = strings.ReplaceAll(iconFunction, "$FUNCTION-PARAMS$", "")
	iconFunction = strings.ReplaceAll(iconFunction, "$FUNCTION-BODY$", body)
	file.WriteString(iconFunction)
}

// create a file to store the icon functions
func (i *IconifyGenerator) generateIconPackage(prefix, category string) (string, error) {
	data, err := os.ReadFile(i.CWD + "/pkg/generator/icon_package.tpl.txt")
	if err != nil {
		return "", err
	}

	goFileName := cleanCategoryName(category) + "_" + strcase.ToSnake(prefix) + ".templ"

	packageTmpl := strings.ReplaceAll(string(data), "$PACKAGE-NAME$", strcase.ToSnake(prefix))
	fileDir := i.CWD + "/" + strcase.ToSnake(prefix)
	filePath := fileDir + "/" + goFileName
	fmt.Println("Generating package: ", fileDir)

	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directories:", err)
		return "", err
	}

	err = os.WriteFile(filePath, []byte(packageTmpl), 0644)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func getIconFunctionName(icon string) string {
	return "Icon" + strcase.ToCamel(icon)
}

func cleanCategoryName(category string) string {
	r := strings.NewReplacer("/", "", "+", "")
	return strcase.ToSnake(
		r.Replace(category),
	)
}
