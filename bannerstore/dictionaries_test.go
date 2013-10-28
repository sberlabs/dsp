package bannerstore_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "dsp/bannerstore"
)

var _ = Describe("Dictionaries", func() {

	var (
		store *BannerStore
		err   error
	)

	BeforeEach(func() {
		store, err = NewBannerStore("https://bayan2cdn.xmlrpc.http.yandex.net:35999",
			"sber-labs",
			"EADCA566-A8BF-403A-950A-0B82B526D2D1")
	})

	Describe("loading GeoInfo from Yandex RTB banner store", func() {
		Context("when GeoInfo loads successfully", func() {
			It("should populate array of GeoInfo correctly", func() {
				info := store.GetGeo()
				Expect(info).ShouldNot(BeNil())
				Expect(len(info)).ShouldNot(BeZero())

				m := make(map[string]int64)
				for i := range info {
					m[info[i].Name] = info[i].Nmb
				}
				Expect(m["Гренландия"]).Should(Equal(int64(21567)))
			})
		})
	})

	Describe("loading MacrosInfo from Yandex RTB banner store", func() {
		Context("when MacrosInfo loads successfully", func() {
			It("should populate array of MacrosInfo correctly", func() {
				info := store.GetMacros()
				Expect(info).ShouldNot(BeNil())
				Expect(len(info)).ShouldNot(BeZero())

				m := make(map[string]int64)
				for i := range info {
					m[info[i].Name] = info[i].Nmb
				}
				Expect(m["{dynamic_banner10}"]).Should(Equal(int64(510)))
			})
		})
	})

	Describe("loading SiteInfo from Yandex RTB banner store", func() {
		Context("when SiteInfo loads successfully", func() {
			It("should populate array of SiteInfo correctly", func() {
				info := store.GetSite()
				Expect(info).ShouldNot(BeNil())
				Expect(len(info)).ShouldNot(BeZero())

				m := make(map[string]int64)
				for i := range info {
					m[info[i].Name] = info[i].Nmb
				}
				Expect(m["Inopressa"]).Should(Equal(int64(9000041)))
			})
		})
	})

	Describe("loading TemplateInfo from Yandex RTB banner store", func() {
		Context("when TemplteInfo loads successfully", func() {
			It("should populate array of TemplateInfo correctly", func() {
				info := store.GetTemplate()
				Expect(info).ShouldNot(BeNil())
				Expect(len(info)).ShouldNot(BeZero())

				m := make(map[string]int64)
				for i := range info {
					m[info[i].Name] = info[i].Nmb
				}
				Expect(m["Обычный с двумя пикселями статистики"]).Should(Equal(int64(13)))
			})
		})
	})

	Describe("loading TnsAdvertiserInfo from Yandex RTB banner store", func() {
		Context("when TnsAdvertiserInfo loads successfully", func() {
			It("should populate array of TnsAdvertiserInfo correctly", func() {
				info := store.GetTnsAdvertiser()
				Expect(info).ShouldNot(BeNil())
				Expect(len(info)).ShouldNot(BeZero())

				m := make(map[string]int64)
				for i := range info {
					m[info[i].Name] = info[i].Nmb
				}
				Expect(m["ТВОРЕЦ"]).Should(Equal(int64(576320)))
			})
		})
	})

	Describe("loading TnsBrandInfo from Yandex RTB banner store", func() {
		Context("when TnsBrandInfo loads successfully", func() {
			It("should populate array of TnsBrandInfo correctly", func() {
				info := store.GetTnsBrand()
				Expect(info).ShouldNot(BeNil())
				Expect(len(info)).ShouldNot(BeZero())

				m := make(map[string]int64)
				for i := range info {
					m[info[i].Name] = info[i].Nmb
				}
				Expect(m["РЕСПУБЛИКА ПОЕТ"]).Should(Equal(int64(769259)))
			})
		})
	})

})
