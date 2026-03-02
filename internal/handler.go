package internal

import (
	"context"
	"fmt"

	"github.com/sqlc-dev/plugin-sdk-go/plugin"
)

// Handler is the sqlc plugin entrypoint.
//
// It receives a GenerateRequest containing parsed SQL queries and
// produces a GenerateResponse with generated Python files.
//
// Queries are grouped by their source filename so that each .sql file
// produces a corresponding .py output file. Template rendering is
// delegated to the render function.
func Handler(ctx context.Context, req *plugin.GenerateRequest) (*plugin.GenerateResponse, error) {
	// Group queries by their originating SQL filename.
	groupedByFile := make(map[string][]*plugin.Query)
	for _, q := range req.Queries {
		groupedByFile[q.Filename] = append(groupedByFile[q.Filename], q)
	}

	var resp plugin.GenerateResponse

	// Generate one Python file per SQL file.
	for file, queries := range groupedByFile {
		content, err := render(queries)
		if err != nil {
			return nil, fmt.Errorf("render: %w", err)
		}

		resp.Files = append(resp.Files, &plugin.File{
			// Replace ".sql" suffix with ".py".
			Name:     file[:len(file)-3] + "py",
			Contents: content,
		})
	}

	return &resp, nil
}
