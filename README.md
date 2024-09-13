# zlogres

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Watsize-Library](https://img.shields.io/badge/Watsize-Library-289548)](https://github.com/owlsome-official)
[![CodeQL](https://github.com/owlsome-official/zlogres/actions/workflows/codeql.yml/badge.svg)](https://github.com/owlsome-official/zlogres/actions/workflows/codeql.yml)

zlogres is a middleware for GoFiber that logging about api elapsed time since request to response.

## Table of Contents

- [zlogres](#zlogres)
  - [Table of Contents](#table-of-contents)
  - [Run this on first time only](#run-this-on-first-time-only)
  - [Installation](#installation)
  - [Signatures](#signatures)
  - [Examples](#examples)
  - [Config](#config)
  - [Default Config](#default-config)
  - [Dependencies](#dependencies)
  - [Example Usage](#example-usage)

## Installation

```bash
  go get -u github.com/owlsome-official/zlogres
```

## Signatures

```go
func New(config ...Config) fiber.Handler
```

## Examples

Import the middleware package that is part of the Fiber web framework

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/owlsome-official/zlogres"
)
```

After you initiate your Fiber app, you can use the following possibilities:

```go
// Default
app.Use(zlogres.New())

// this middleware supported the `requestid` middleware
app.Use(requestid.New())
app.Use(zlogres.New())

// Or extend your config for customization
app.Use(requestid.New(requestid.Config{
  ContextKey: "transaction-id",
}))
app.Use(zlogres.New(zlogres.Config{
  RequestIDContextKey: "transaction-id",
}))
```

## Config

```go
// Config defines the config for middleware.
type Config struct {
  // Optional. Default: nil
  Next func(c *fiber.Ctx) bool

  // Optional. Default: "requestid"
  RequestIDContextKey string
}
```

## Default Config

```go
var ConfigDefault = Config{
  Next:                nil,
  RequestIDContextKey: "requestid",
}
```

## Dependencies

- [Zerolog](https://github.com/rs/zerolog)
- [Fiber](https://github.com/gofiber/fiber)

## Example Usage

Please go to [example/main.go](./example/main.go)

**Don't forget to run:**

```bash
  go mod tidy
```

Note: Custom usage please focus on `Custom` section
