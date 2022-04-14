// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package transfer_test

import (
	"bytes"

	. "github.com/open-component-model/ocm/cmds/ocm/testhelper"
	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/mime"
	"github.com/open-component-model/ocm/pkg/oci/repositories/ctf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const ARCH = "/tmp/ctf"
const VERSION = "v1"
const NS = "mandelsoft/test"
const OUT = "/tmp/res"

var _ = Describe("Test Environment", func() {
	var env *TestEnv

	BeforeEach(func() {
		env = NewTestEnv()
		env.OCICommonTransport(OUT, accessio.FormatDirectory)
	})

	AfterEach(func() {
		env.Cleanup()
	})

	It("transfers a named artefact", func() {
		env.OCICommonTransport(ARCH, accessio.FormatDirectory, func() {
			env.Namespace(NS, func() {
				env.Manifest(VERSION, func() {
					env.Config(func() {
						env.BlobStringData(mime.MIME_JSON, "{}")
					})
					env.Layer(func() {
						env.BlobStringData(mime.MIME_TEXT, "testdata")
					})
				})
			})
		})

		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("transfer", "artefact", ARCH+"//"+NS+":"+VERSION, "directory::"+OUT)).To(Succeed())
		Expect("\n" + buf.String()).To(Equal(
			`
copying mandelsoft/test:v1 to mandelsoft/test:v1...
copied 1 from 1 artefact(s)
`))
		Expect(env.ReadFile(OUT + "/" + ctf.ArtefactIndexFileName)).To(Equal([]byte("{\"schemaVersion\":1,\"artefacts\":[{\"repository\":\"mandelsoft/test\",\"tag\":\"v1\",\"digest\":\"sha256:2c3e2c59e0ac9c99864bf0a9f9727c09f21a66080f9f9b03b36a2dad3cce6ff9\"}]}")))
	})

	It("transfers an unnamed artefact set", func() {
		env.ArtefactSet(ARCH, accessio.FormatDirectory, func() {
			env.Manifest(VERSION, func() {
				env.Config(func() {
					env.BlobStringData(mime.MIME_JSON, "{}")
				})
				env.Layer(func() {
					env.BlobStringData(mime.MIME_TEXT, "testdata")
				})
			})
		})

		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("transfer", "artefact", ARCH, "directory::"+OUT+"//"+NS)).To(Succeed())
		Expect("\n" + buf.String()).To(Equal(
			`
copying :v1 to mandelsoft/test:v1...
copied 1 from 1 artefact(s)
`))
		Expect(env.ReadFile(OUT + "/" + ctf.ArtefactIndexFileName)).To(Equal([]byte("{\"schemaVersion\":1,\"artefacts\":[{\"repository\":\"mandelsoft/test\",\"tag\":\"v1\",\"digest\":\"sha256:2c3e2c59e0ac9c99864bf0a9f9727c09f21a66080f9f9b03b36a2dad3cce6ff9\"}]}")))
	})
})
