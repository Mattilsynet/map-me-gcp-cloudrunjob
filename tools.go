//go:build tools
// +build tools

package tools

//go:generate go get -u ./...
//go:generate go mod tidy
