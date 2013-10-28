package bannerstore_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "rtbc/bannerstore"
)

var _ = Describe("Creatives", func() {
	var (
		store *BannerStore
		err   error
	)

	BeforeEach(func() {
		store, err = NewBannerStore("https://bayan2cdn.xmlrpc.http.yandex.net:35999",
			"sber-labs",
			"EADCA566-A8BF-403A-950A-0B82B526D2D1")
	})

	XDescribe("Creating new Creative in Yandex RTB banner store", func() {
		Context("when passed correct parameters", func() {
			It("should return correct id", func() {
				var id int64
				err := store.CreateCreative(CreateCreativeInfo{
					"Advertizing",
					67660,
					1,
					"20120802T19:30:00",
					"main",
					"Продвижение услуг рекламного агентства"}, &id)
				Expect(err).NotTo(HaveOccured())
				Expect(id).ShouldNot(BeZero())
			})
		})
	})

	XDescribe("Getting Creative by number from Yandex RTB banner store", func() {
		Context("when passed correct parameters", func() {
			It("should return correct CreativeInfo", func() {
				var info CreativeInfo
				err := store.GetCreativeByNmb(int64(67660), &info)
				Expect(err).NotTo(HaveOccured())
				Expect(info).ShouldNot(BeNil())
				Expect(info.CreativeNmb).To(Equal(int64(67660)))
			})
		})
	})

	XDescribe("Getting Creative by tag from Yandex RTB banner store", func() {
		Context("when passed correct parameters", func() {
			It("should return correct CreativeInfo", func() {
				var info []CreativeInfo
				err := store.GetCreativeByTag("main", info)
				Expect(err).NotTo(HaveOccured())
				Expect(info).ShouldNot(BeNil())
				Expect(len(info)).ShouldNot(BeZero())
			})
		})
	})

})
