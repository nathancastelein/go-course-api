# From OpenAPI

This package provides HTTP handlers for the Pet Shelter API.

HTTP handlers have been generated using the [OAPI Codegen](https://github.com/deepmap/oapi-codegen) tool.

## Install the tool

```bash
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
```

## Generate handlers

```bash
oapi-codegen -package openapi ./openapi.yaml > petshelter.gen.go
```