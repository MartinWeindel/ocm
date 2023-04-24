// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

//go:generate go-bindata -pkg jsonscheme ../../../../../../../resources/component-descriptor-v2-schema.yaml
//go:generate gofmt -s -w bindata.go

package jsonscheme

import (
	"github.com/ghodss/yaml"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

var Schema *jsonschema.Schema

func init() {
	dataBytes, err := ResourcesComponentDescriptorV2SchemaYamlBytes()
	if err != nil {
		panic(err)
	}

	Schema, err = jsonschema.CompileString("component-descriptor-v2-schema.yaml", string(dataBytes))
	if err != nil {
		panic(err)
	}
}

// Validate validates the given data against the component descriptor v2 jsonscheme.
func Validate(src []byte) error {
	data := map[any]any{}
	err := yaml.Unmarshal(src, data)
	if err != nil {
		return err
	}
	return Schema.Validate(data)
}
