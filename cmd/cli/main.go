package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	repoCmd := flag.NewFlagSet("repo", flag.ExitOnError)
	fmt.Println(os.Args)

	switch os.Args[1] {
	case "repo":
		path := repoCmd.String("path", "", "path name")
		model := repoCmd.String("model", "", "model name")
		// 解析命令行参数
		err := repoCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Parse:", err)
		}
		fullPath := "internal/repository"
		if *path != "" {
			fullPath = fullPath + "/" + *path
		}
		fmt.Println("fullPath:", fullPath)

		fmt.Println("name:", *model)
		if *model == "" {
			fmt.Println("Usage: repo-gen [command]")
			fmt.Println("Available commands:")
			fmt.Println("  repo -name <name>")
			os.Exit(1)
		}
		generateRepository(fullPath, *model)

	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

func generateRepository(path, name string) {
	fmt.Printf("Generating repository file for %s...\n", name)

	lastSegment := filepath.Base(path)
	fmt.Println("package:", lastSegment)

	var modelName string
	var prefix, suffix string
	dotIndex := strings.Index(name, ".")
	fmt.Println("dotIndex:", dotIndex)
	if dotIndex != -1 {
		prefix = name[:dotIndex+1]
		suffix = name[dotIndex+1:]
		modelName = prefix + "." + strings.ToLower(suffix)
	} else {
		modelName = strings.ToLower(name)
	}
	fmt.Println("modelName:", modelName)
}
