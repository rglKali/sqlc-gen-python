package internal

import (
	"testing"

	"github.com/sqlc-dev/plugin-sdk-go/plugin"
)

type resolveTestCase struct {
	name  string
	ident *plugin.Identifier
	want  string
}

func TestResolve(t *testing.T) {
	tests := []resolveTestCase{
		// Integers
		{"int4", &plugin.Identifier{Name: "int4"}, "int"},
		{"pg_catalog.int4", &plugin.Identifier{Schema: "pg_catalog", Name: "int4"}, "int"},
		{"serial", &plugin.Identifier{Name: "serial"}, "int"},

		// Floats
		{"float8", &plugin.Identifier{Name: "float8"}, "float"},
		{"pg_catalog.float8", &plugin.Identifier{Schema: "pg_catalog", Name: "float8"}, "float"},

		// Numeric
		{"numeric", &plugin.Identifier{Name: "numeric"}, "Decimal"},
		{"pg_catalog.numeric", &plugin.Identifier{Schema: "pg_catalog", Name: "numeric"}, "Decimal"},

		// Boolean
		{"bool", &plugin.Identifier{Name: "bool"}, "bool"},

		// Binary
		{"bytea", &plugin.Identifier{Name: "bytea"}, "bytes"},

		// Date/Time
		{"date", &plugin.Identifier{Name: "date"}, "date"},
		{"time", &plugin.Identifier{Schema: "pg_catalog", Name: "time"}, "time"},
		{"timestamp", &plugin.Identifier{Schema: "pg_catalog", Name: "timestamp"}, "datetime"},
		{"timestamptz", &plugin.Identifier{Name: "timestamptz"}, "datetime"},
		{"interval", &plugin.Identifier{Name: "interval"}, "timedelta"},

		// Text
		{"text", &plugin.Identifier{Name: "text"}, "str"},
		{"varchar", &plugin.Identifier{Schema: "pg_catalog", Name: "varchar"}, "str"},
		{"bpchar", &plugin.Identifier{Schema: "pg_catalog", Name: "bpchar"}, "str"},
		{"jsonb", &plugin.Identifier{Name: "jsonb"}, "str"},

		// UUID
		{"uuid", &plugin.Identifier{Name: "uuid"}, "UUID"},

		// Network & misc (mapped to str or Any)
		{"inet", &plugin.Identifier{Name: "inet"}, "Any"},
		{"cidr", &plugin.Identifier{Name: "cidr"}, "Any"},
		{"point", &plugin.Identifier{Name: "point"}, "Any"},

		// Unknown type falls back to str
		{"unknown", &plugin.Identifier{Name: "weird_type"}, "str"},
		{"schema.unknown", &plugin.Identifier{Schema: "schema", Name: "weird"}, "str"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolve(tt.ident)
			if got != tt.want {
				t.Fatalf("resolve(DataType(%v)) = %q, want %q", tt.ident, got, tt.want)
			}
		})
	}
}
