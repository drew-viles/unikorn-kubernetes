// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/unikorn-cloud/core/pkg/openapi"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "https://raw.githubusercontent.com/unikorn-cloud/core/main/pkg/openapi/common.spec.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}