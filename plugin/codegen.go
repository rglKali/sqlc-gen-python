// Copyright © 2026 Vadim Tertilov <vadimtertilov@gmail.com>
// All rights reserved.
package main

import (
	"github.com/rglKali/sqlc-gen-python/internal"
	"github.com/sqlc-dev/plugin-sdk-go/codegen"
)

func main() {
	codegen.Run(internal.Handler)
}
