package bannerstore_test

import (
        . "github.com/onsi/ginkgo"
        . "github.com/onsi/gomega"
        . "dsp/bannerstore"
)

var _ = Describe("Bannerstore", func() {
        var (
                store   *BannerStore
                err     error
        )

        BeforeEach(func() {
                store, err = NewBannerStore("https://bayan2cdn.xmlrpc.http.yandex.net:35999",
                        "sber-labs",
                        "EADCA566-A8BF-403A-950A-0B82B526D2D1")
        })

        Describe("Logging in to Yandex RTB banner store", func() {
                Context("when passed correct login and password parameters", func() {
                        It("should return correct LogonInfo", func() {
                                var logon string
                                logon = store.CreateLogon()
                                Expect(logon).ShouldNot(BeEmpty())
                        })
                })
        })

})
