package internal

import (
	"github.com/sqlc-dev/plugin-sdk-go/plugin"
	"github.com/sqlc-dev/plugin-sdk-go/sdk"
)

func resolve(ident *plugin.Identifier) string {
	switch sdk.DataType(ident) {

	// --- Integers ---
	case "serial", "serial4", "pg_catalog.serial4",
		"bigserial", "serial8", "pg_catalog.serial8",
		"smallserial", "serial2", "pg_catalog.serial2",
		"integer", "int", "int4", "pg_catalog.int4",
		"bigint", "int8", "pg_catalog.int8",
		"smallint", "int2", "pg_catalog.int2":
		return "int"

	// --- Floats ---
	case "float", "double precision", "float8", "pg_catalog.float8",
		"real", "float4", "pg_catalog.float4":
		return "float"

	// --- Exact numeric ---
	// asyncpg returns Decimal (not decimal.Decimal — imported as "from decimal import Decimal")
	case "numeric", "pg_catalog.numeric":
		return "Decimal"

	// --- Boolean ---
	case "boolean", "bool", "pg_catalog.bool":
		return "bool"

	// --- Binary ---
	case "bytea", "blob", "pg_catalog.bytea":
		return "bytes"

	// --- Date/Time ---
	case "date":
		return "date"
	case "pg_catalog.time", "pg_catalog.timetz":
		return "time"
	case "pg_catalog.timestamp", "pg_catalog.timestamptz", "timestamptz":
		return "datetime"
	case "interval", "pg_catalog.interval":
		return "timedelta"

	// --- Text ---
	case "text", "pg_catalog.varchar", "pg_catalog.bpchar", "string", "citext",
		"char", "name", "xml", "json", "jsonb":
		return "str"

	// --- UUID ---
	case "uuid":
		return "UUID"

	// --- Money ---
	case "money":
		fallthrough

	// --- ltree extension ---
	case "ltree", "lquery", "ltxtquery":
		fallthrough

	// --- Network types ---
	case "macaddr", "macaddr8":
		return "str"

	// asyncpg returns ipaddress.IPv4Interface/IPv6Interface for inet and
	case "inet", "cidr":
		fallthrough

	// --- Geometric types ---
	case "point", "box", "circle", "line", "lseg", "path", "polygon":
		fallthrough

	// --- Bit strings ---
	case "bit", "varbit":
		return "Any"

	default:
		return "str"
	}
}
