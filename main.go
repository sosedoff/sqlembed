package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	version = "0.1.1"
)

var (
	path        string
	packageName string
	showVersion bool
	goEmbed     bool
	goEmbedDir  string
)

type queryItem struct {
	File string
	Name string
	Data string
}

func init() {
	flag.StringVar(&path, "path", "", "Path to directory containing SQL files")
	flag.StringVar(&packageName, "package", "queries", "Output package name")
	flag.BoolVar(&showVersion, "v", false, "Show current version")
	flag.BoolVar(&goEmbed, "embed", false, "Use go:embed directives")
	flag.StringVar(&goEmbedDir, "embed-dir", "", "Relative directory for go:embed")
	flag.Parse()
}

func fatal(msg interface{}) {
	fmt.Fprintf(os.Stderr, "%s", msg)
	os.Exit(1)
}

func main() {
	if path == "" {
		fatal("path is not provided")
	}

	items := []queryItem{}

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || strings.ToLower(filepath.Ext(info.Name())) != ".sql" {
			return nil
		}

		baseName := filepath.Base(info.Name())
		baseName = strings.Replace(baseName, filepath.Ext(baseName), "", 1)
		constName := constantizeName(baseName)

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Go embed file path must be relative to the package directory
		if goEmbedDir != "" {
			path = filepath.Join(goEmbedDir, filepath.Base(info.Name()))
		}

		items = append(items, queryItem{
			File: path,
			Name: constName,
			Data: fmt.Sprintf("`%s`", strings.ReplaceAll(string(data), "\n", " ")),
		})

		return nil
	})

	targetTemplate := contentTemplate
	if goEmbed {
		targetTemplate = embedTemplate
	}

	tpl, err := template.New("main").Parse(targetTemplate)
	if err != nil {
		fatal(err)
	}

	err = tpl.Execute(os.Stdout, map[string]interface{}{
		"packageName": packageName,
		"items":       items,
	})
	if err != nil {
		fatal(err)
	}
}

func constantizeName(name string) string {
	parts := strings.Split(name, "_")
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, "")
}

const (
	contentTemplate = `
package {{ .packageName }}
const (
{{- range .items }}
	// {{ .Name }} is imported from {{ .File }}
	{{ .Name }} = {{ .Data }}
{{ end -}}
)
`

	embedTemplate = `
package {{ .packageName }}

import (
	_ "embed"
)

var (
{{- range .items }}
	//go:embed {{ .File }}
	{{ .Name }} string
{{ end -}}
)
`
)
