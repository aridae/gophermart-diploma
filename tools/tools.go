//go:build tools
// +build tools

package tools

import (
	_ "github.com/aridae/gophermart-diploma/internal/logger"
	_ "github.com/caarlos0/env/v6"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
)
