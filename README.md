## Usage

```yaml
version: "2"
plugins:
  - name: py
    wasm:
      url: <to check at release page>
      sha256: <to check at release page>
sql:
  - schema: "schema.sql"
    queries: "queries.sql"
    engine: postgresql
    codegen:
      - out: src/querries
        plugin: py
```

## Structure
```
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── internal
│   ├── caseconv.go
│   ├── caseconv_test.go
│   ├── handler.go
│   ├── query.tmpl        -- querier template
│   ├── render.go
│   ├── resolve.go
│   └── resolve_test.go
└── plugin
    └── codegen.go        -- entrypoint
```

## Options:
To be done later...
