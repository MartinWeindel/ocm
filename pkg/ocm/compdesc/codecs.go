// Copyright 2020 Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compdesc

import (
	"encoding/json"

	"sigs.k8s.io/yaml"
)

// Codec describes an encoding for en object
type Codec interface {
	Decoder
	StrictDecoder
	Encoder
}

// CodecWrapper is a simple struct that implements the Codec interface
type CodecWrapper struct {
	Decoder
	StrictDecoder
	Encoder
}

// ApplyEncodeOption applies the configured codec.
func (o CodecWrapper) ApplyEncodeOption(options *EncodeOptions) {
	options.Codec = o
}

// ApplyDecodeOption applies the configured codec.
func (o CodecWrapper) ApplyDecodeOption(options *DecodeOptions) {
	options.Codec = o
}

// Decoder defines a decoder for an object.
type Decoder interface {
	Decode(data []byte, obj interface{}) error
}

// StrictDecoder defines a decoder for an object.
type StrictDecoder interface {
	DecodeStrict(data []byte, obj interface{}) error
}

// Encoder defines a encoder for an object.
type Encoder interface {
	Encode(obj interface{}) ([]byte, error)
}

// DecoderFunc is a simple function that implements the Decoder interface.
type DecoderFunc func(data []byte, obj interface{}) error

// Decode is the Decode implementation of the Decoder interface.
func (e DecoderFunc) Decode(data []byte, obj interface{}) error {
	return e(data, obj)
}

// StrictDecoderFunc is a simple function that implements the StrictDecoder interface.
type StrictDecoderFunc func(data []byte, obj interface{}) error

// StrictDecode is the Decode implementation of the Decoder interface.
func (e StrictDecoderFunc) DecodeStrict(data []byte, obj interface{}) error {
	return e(data, obj)
}

// EncoderFunc is a simple function that implements the Encoder interface.
type EncoderFunc func(obj interface{}) ([]byte, error)

// Encode is the Encode implementation of the Encoder interface.
func (e EncoderFunc) Encode(obj interface{}) ([]byte, error) {
	return e(obj)
}

// DefaultYAMLCodec implements Codec interface with the yaml decoder encoder.
var DefaultYAMLCodec = CodecWrapper{
	Decoder:       DecoderFunc(func(data []byte, obj interface{}) error { return yaml.Unmarshal(data, obj) }),
	StrictDecoder: StrictDecoderFunc(func(data []byte, obj interface{}) error { return yaml.UnmarshalStrict(data, obj) }),
	Encoder:       EncoderFunc(yaml.Marshal),
}

// DefaultJSONCodec implements Codec interface with the json decoder encoder.
var DefaultJSONLCodec = CodecWrapper{
	Decoder:       DecoderFunc(json.Unmarshal),
	StrictDecoder: StrictDecoderFunc(func(data []byte, obj interface{}) error { return json.Unmarshal(data, obj) }),
	Encoder:       EncoderFunc(json.Marshal),
}
