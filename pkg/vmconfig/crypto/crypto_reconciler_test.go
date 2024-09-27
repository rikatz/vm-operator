// Copyright (c) 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package crypto_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vmware-tanzu/vm-operator/pkg/constants/testlabels"
	"github.com/vmware-tanzu/vm-operator/pkg/vmconfig/crypto"
)

var _ = Describe("New", Label(testlabels.Crypto), func() {
	It("should return a reconciler", func() {
		Expect(crypto.New()).ToNot(BeNil())
	})
})

var _ = Describe("Name", Label(testlabels.Crypto), func() {
	It("should return crypto", func() {
		Expect(crypto.New().Name()).To(Equal("crypto"))
	})
})
