package main

import (
	"github.com/genesysflow/iconify/pkg/generator"
	"github.com/spf13/cobra"
)

var api string
var generatorCmd = &cobra.Command{
	Use:   "generator",
	Short: "Generate templ icons using the specified template",
	Long:  `Generate templ icons using the specified template by going through the api provided by iconify`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		generator.Generate(api)
	},
}

func main() {
	generatorCmd.PersistentFlags().StringVarP(&api, "api", "a", "http://localhost:3000", "Api path to use when generating icons")
	generatorCmd.Execute()
}
