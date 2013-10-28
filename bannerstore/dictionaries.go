package bannerstore

// Yandex RTB BannerStore
//
// Dictionaries
//

type GeoInfo struct {
        Nmb     int64
        Name    string
}

type MacrosInfo struct {
        Nmb           int64
        Name          string
        MacrosTypeNmb int64
        Caption       string
        SortNmb       int64
        IsEnabled     bool
}

type SiteInfo struct {
        Nmb           int64
        Name          string
        IsSecure      bool
        SiteStatusNmb int64
}

type TemplateInfo struct {
        Nmb        int64
        Name       string
        Data       string
        IsEnabled  bool
        IsApproved bool
}

type TnsAdvertiserInfo struct {
        Nmb     int64
        Name    string
}

type TnsArticleInfo struct {
        Nmb     int64
        Name    string
}

type TnsBrandInfo struct {
        Nmb     int64
        Name    string
}
