// SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors.
//
// SPDX-License-Identifier: Apache-2.0

package envtest_test

import (
	"context"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gardener/component-cli/ociclient/test/envtest"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "envtest Test Suite")
}

var _ = Describe("Test Environment", func() {

	It("should run and stop a test registry", func() {
		ctx := context.Background()
		defer ctx.Done()
		testenv := envtest.New(envtest.Options{
			RegistryBinaryPath: filepath.Join("../../../", envtest.DefaultRegistryBinaryPath),
		})
		Expect(testenv.Start(ctx)).To(Succeed())

		Expect(testenv.Close()).To(Succeed())
	})

})
