package bannerstore_test

import (
        . "dsp/bannerstore"
        . "github.com/onsi/ginkgo"
        . "github.com/onsi/gomega"
)

var _ = Describe("Creatives", func() {
        var (
                store   *BannerStore
                err     error
        )

        BeforeEach(func() {
                store, err = NewBannerStore("https://bayan2cdn.xmlrpc.http.yandex.net:35999",
                        "sber-labs",
                        "EADCA566-A8BF-403A-950A-0B82B526D2D1")
        })

        Describe("Creating new Creative in Yandex RTB banner store", func() {
                Context("when passed correct parameters", func() {
                        It("should return correct id", func() {
                                id := store.CreateCreative(CreateCreativeInfo{
                                        "Advertizing",
                                        67660,
                                        1,
                                        "20120802T19:30:00",
                                        "main",
                                        "Продвижение услуг рекламного агентства"})
                                Expect(id).ShouldNot(BeZero())
                        })
                })
        })

        Describe("Getting Creative by number from Yandex RTB banner store", func() {
                Context("when passed correct parameters", func() {
                        It("should return correct CreativeInfo", func() {
                                info := store.GetCreativeByNmb(int64(67660))
                                Expect(info).ShouldNot(BeNil())
                                Expect(info.CreativeNmb).To(Equal(int64(67660)))
                        })
                })
        })

        Describe("Getting Creative by tag from Yandex RTB banner store", func() {
                Context("when passed correct parameters", func() {
                        It("should return correct CreativeInfo", func() {
                                info := store.GetCreativeByTag("main")
                                Expect(info).ShouldNot(BeNil())
                                Expect(len(info)).ShouldNot(BeZero())
                        })
                })
        })

})
