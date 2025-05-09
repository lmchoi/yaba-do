/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

type TemplateData struct {
	Namespace   string
}

func createFiles(projectName string) {
	data := TemplateData{
		Namespace:   projectName,
	}

	// Define template files to generate
	templates := []struct {
		templatePath string
		outputPath   string
	}{
		{
			"templates/clojure/base/src/core.clj",
			filepath.Join(projectName, "src", "core.clj"),
		},
		{
			"templates/clojure/base/test/core_test.clj",
			filepath.Join(projectName, "test", "core_test.clj"),
		},
	}

	// Create files
	for _, tmpl := range templates {
		// Parse and execute template
		t, err := template.ParseFiles(tmpl.templatePath)
		if err != nil {
			fmt.Printf("Error parsing template: %v\n", err)
			return
		}

		file, err := os.Create(tmpl.outputPath)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return
		}
		defer file.Close()

		if err := t.Execute(file, data); err != nil {
			fmt.Printf("Error executing template: %v\n", err)
			return
		}
	}
}

func createDirectories(cmd *cobra.Command, args []string) {
	projectName := args[0]

	// Check if project directory already exists
	if _, err := os.Stat(projectName); err == nil {
		fmt.Printf("aborting: directory '%s' already exists\n", projectName)
		return
	}
	
	// Create directories
	dirs := []string{
		filepath.Join(projectName, "src"),
		filepath.Join(projectName, "test"),
	}
	
	// Set file permissions so that owner have full access and group and others can read and execute
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			return
		}
	}

	// Create files
	createFiles(projectName)
}

// cljCmd represents the clj command
var cljCmd = &cobra.Command{
	Use:   "clj [project-name]",
	Short: "Create a new Clojure (deps.edn) project",
	Long: `Creates a basic Clojure project structure with:
- src/
- test/`,
	Args: cobra.ExactArgs(1), // Requires exactly 1 argument (project name)
	Run: createDirectories,
}

func init() {
	rootCmd.AddCommand(cljCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cljCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cljCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
