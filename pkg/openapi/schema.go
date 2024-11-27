// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
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
	externalRef1 "github.com/unikorn-cloud/region/pkg/openapi"
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

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(path.Dir(pathToFile), "https://raw.githubusercontent.com/unikorn-cloud/core/main/pkg/openapi/common.spec.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef1.PathToRawSpec(path.Join(path.Dir(pathToFile), "https://raw.githubusercontent.com/unikorn-cloud/region/main/pkg/openapi/server.spec.yaml")) {
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
