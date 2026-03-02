package internal

import (
	"bytes"
	"fmt"
	"text/template"

	_ "embed"

	"github.com/sqlc-dev/plugin-sdk-go/plugin"
)

// tmplRaw contains the embedded query template source.
// It is loaded at compile time via go:embed.
//
//go:embed query.tmpl
var tmplRaw string

// tmpl is the parsed text/template instance used to render
// Python query code. All helper functions are registered
// before parsing to ensure they are available inside the template.
var tmpl = template.Must(
	template.New("query.tmpl").
		Funcs(template.FuncMap{
			"resolve":      resolve,
			"toSnakeCase":  toSnakeCase,
			"toUpperCase":  toUpperCase,
			"toCamelCase":  toCamelCase,
			"toPascalCase": toPascalCase,
		}).
		Parse(tmplRaw),
)

// render executes the embedded query template using the provided
// sqlc plugin queries and returns the generated Python source code.
//
// It returns an error if template execution fails.
func render(queries []*plugin.Query) ([]byte, error) {
	var out bytes.Buffer

	if err := tmpl.Execute(&out, queries); err != nil {
		return nil, fmt.Errorf("tmpl.Execute: %w", err)
	}

	return out.Bytes(), nil
}
