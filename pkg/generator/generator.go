package generator

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/genesysflow/iconify/pkg/api/iconify"
	"github.com/iancoleman/strcase"
)

func Generate(api string) {
	// Do Stuff Here
	collections, err := iconify.GetCollections(api)
	if err != nil {
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for _, collection := range collections {
		iconCollection, err := iconify.GetIconCollection(api, collection.Key)
		if err != nil {
			return
		}

		filePath, err := generateIconPackage(cwd, iconCollection.Prefix)
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

		for key, value := range iconCollection.Categories {
			fmt.Println(key)
			for _, icon := range value {
				body, err := iconify.GetIcon(api, iconCollection.Prefix, icon)
				if err != nil {
					return
				}
				generateIcon(file, cwd, icon, body)
			}
		}

		for _, icon := range iconCollection.Uncategorized {
			body, err := iconify.GetIcon(api, iconCollection.Prefix, icon)
			if err != nil {
				return
			}
			generateIcon(file, cwd, icon, body)
		}
	}
}

var iconFunctionTemplate = ""

func generateIcon(file *os.File, cwd, icon, body string) {
	functionName := getIconFunctionName(icon)
	if iconFunctionTemplate == "" {
		data, err := os.ReadFile(cwd + "/pkg/generator/icon_function.tpl.txt")
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

func generateIconPackage(cwd, prefix string) (string, error) {
	data, err := os.ReadFile(cwd + "/pkg/generator/icon_package.tpl.txt")
	if err != nil {
		return "", err
	}

	goFileName := "" + strcase.ToSnake(prefix) + ".templ"

	packageTmpl := strings.ReplaceAll(string(data), "$PACKAGE-NAME$", strcase.ToDelimited(prefix, 0))
	fileDir := cwd + "/" + strcase.ToSnake(prefix)
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
