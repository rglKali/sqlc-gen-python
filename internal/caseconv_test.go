package internal

import (
	"testing"
)

type caseConvTestCase struct {
	input string
	want  string
}

func TestToSnakeCase(t *testing.T) {
	tests := []caseConvTestCase{
		{"camelCase", "camel_case"},
		{"PascalCase", "pascal_case"},
		{"snake_case", "snake_case"},
		{"UPPER_CASE", "upper_case"},
		{"XMLParser", "xml_parser"},
		{"userID", "user_id"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toSnakeCase(tt.input)
			if got != tt.want {
				t.Fatalf("toSnakeCase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToUpperCase(t *testing.T) {
	tests := []caseConvTestCase{
		{"camelCase", "CAMEL_CASE"},
		{"snake_case", "SNAKE_CASE"},
		{"PascalCase", "PASCAL_CASE"},
		{"xmlParser", "XML_PARSER"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toUpperCase(tt.input)
			if got != tt.want {
				t.Fatalf("toUpperCase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []caseConvTestCase{
		{"snake_case", "snakeCase"},
		{"PascalCase", "pascalCase"},
		{"UPPER_CASE", "upperCase"},
		{"xml_parser", "xmlParser"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toCamelCase(tt.input)
			if got != tt.want {
				t.Fatalf("toCamelCase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []caseConvTestCase{
		{"snake_case", "SnakeCase"},
		{"camelCase", "CamelCase"},
		{"UPPER_CASE", "UpperCase"},
		{"xml_parser", "XmlParser"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toPascalCase(tt.input)
			if got != tt.want {
				t.Fatalf("toPascalCase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	tests := []caseConvTestCase{
		{"hello", "Hello"},
		{"HELLO", "Hello"},
		{"h", "H"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := capitalize(tt.input)
			if got != tt.want {
				t.Fatalf("capitalize(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
