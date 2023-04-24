// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package spiff

import (
	"github.com/mandelsoft/spiff/spiffing"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"sigs.k8s.io/yaml"

	"github.com/open-component-model/ocm/pkg/errors"
)

func ValidateSourceByScheme(src spiffing.Source, schemedata []byte) error {
	data, err := src.Data()
	if err != nil {
		return err
	}
	return ValidateByScheme(data, schemedata)
}

func ValidateByScheme(src []byte, schemedata []byte) error {
	schema, err := jsonschema.CompileString("schema.yaml", string(schemedata))
	if err != nil {
		return errors.Wrapf(err, "invalid scheme")
	}

	data := map[any]any{}
	err = yaml.Unmarshal(src, data)
	if err != nil {
		return err
	}
	return schema.Validate(data)
}
