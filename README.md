# Iconify for go Templ

[![Go Reference](https://pkg.go.dev/badge/github.com/genesysflow/iconify.svg)](https://pkg.go.dev/github.com/genesysflow/iconify)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

`genesysflow/iconify` is a Go library that allows you to easily use [Iconify](https://iconify.design/) icons in your Templ templates. Iconify provides a large collection of icons from various icon sets, and this library integrates them seamlessly with the Templ templating engine.

## Features

- Easy integration of Iconify icons in Templ templates
- Supports all icons available in Iconify
- Lightweight and simple to use
- Supported icon sets:
    - mdi
    - fa

## Installation

To install genesysflow/iconify, use `go get`:

```sh
go get github.com/genesysflow/iconify
```

## Example

```go

package view

import (
    "github.com/genesysflow/iconify/mdi"
)

templ Example() {
    @mdi.IconAlien()
}