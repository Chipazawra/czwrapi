// Package classification albums API.
//
// Documentation fo Product API
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	samplegin "github.com/Chipazawra/czwrapi/internal/sample-gin"
)

//go:generate swagger generate spec
func main() {
	samplegin.Run()
}
