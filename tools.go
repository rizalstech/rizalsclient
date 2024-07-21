//go:build tools

package rizalsclient

// ignore the warnings, this is for the go mod system to know that this is a tool
// / but vscode doesn't know that so it will show warnings
import (
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
)
