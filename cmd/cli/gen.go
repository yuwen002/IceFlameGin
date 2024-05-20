package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
	"unicode"
)

func main() {
	repoCmd := flag.NewFlagSet("repo", flag.ExitOnError)
	fmt.Println(os.Args)

	switch os.Args[1] {
	case "repo":
		path := repoCmd.String("path", "", "path name")
		model := repoCmd.String("model", "", "model name")
		structDescription := repoCmd.String("struct", "", "struct description")
		// 解析命令行参数
		err := repoCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Parse:", err)
		}
		fullPath := "cmd/gen/repo"
		if *path != "" {
			fullPath = fullPath + "/" + *path
		}
		fmt.Println("fullPath:", fullPath)

		fmt.Println("name:", *model)
		if *model == "" {
			fmt.Println("Usage: repo-gen [command]")
			fmt.Println("Available commands:")
			fmt.Println("  repo -model <name>")
			os.Exit(1)
		}
		generateRepository(fullPath, *model, *structDescription)

	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

func camelToSnake(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}

func generateRepository(path, name, structDescription string) {
	fmt.Printf("Generating repository file for %s...\n", name)

	lastSegment := filepath.Base(path)
	fmt.Println("package:", lastSegment)

	dotIndex := strings.Index(name, ".")
	if dotIndex == -1 {
		fmt.Println("name is not a dot")
		return
	}

	// 引入包名字
	importModelName := name[:dotIndex]
	// struct名字
	structName := name[dotIndex+1:]
	// 表名
	tableName := camelToSnake(structName)
	// 完整的model名字
	fullModelName := name

	data := map[string]any{
		"package_name":       lastSegment,
		"import_model_name":  importModelName,
		"struct_name":        structName,
		"table_name":         tableName,
		"model_name":         fullModelName,
		"struct_description": structDescription,
		"date":               time.Now().Format("2006-01-02 15:04:05"),
	}

	executable, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
		return
	}

	dir := filepath.Dir(executable)
	fmt.Println("Current directory:", dir)

	// 创建并解析模板
	tmpl, err := template.New("repo.tmpl").ParseFiles(filepath.Join(dir, "tmpl", "repo.tmpl"))
	if err != nil {
		fmt.Println("ParseFiles:", err)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		fmt.Println("Execute:", err)
		return
	}

	repoDir := filepath.Join(dir, "gen", "repo")
	if _, err := os.Stat(repoDir); os.IsNotExist(err) {
		if err := os.MkdirAll(repoDir, 0755); err != nil {
			// 处理创建目录错误
			fmt.Println("MkdirAll:", err)
			return
		}
	}

	repoName := filepath.Join(dir, "gen", "repo", tableName+".gen_repo.go")
	if err := os.WriteFile(repoName, buf.Bytes(), 0644); err != nil {
		fmt.Println("Error writing to file:", err)
	}
	fmt.Println("Generated repository file:", repoName)
}
