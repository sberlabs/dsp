package bannerstore

// Yandex RTB BannerStore
//
// Creatives
//

type CreateCreativeInfo struct {
        CreativeName     string
        TnsAdvertiserNmb int64
        TemplateNmb      int64
        ExpireDate       string  // "20120802Т19:30:00"
        Таg              string
        Note             string
}

type CreativeInfo struct {
        CreativeNmb      int64
        CreativeName     string
        TnsAdvertiserNmb int64
        TemplateData     string
        Tag              string
        Moderation       CreativeModeratedInfo
        IsDeployed       bool
        Token            string
        Data             string
        Properties       string
        ExpireDate       string  // "20120802Т19:30:00"
        Note             string
}

type CreativeModeratedInfo struct {
        StatusNmb     int64
        ModeratedDate string
        Message       string
        RequestDate   string
        Log           []CreativeModeratingLogInfo
}

type CreativeModeratingLogInfo struct {
        Date      string
        StatusNmb int64
        Message   string
}

type CreativeMacrosInfo struct {
        MacrosNmb  int64
        MacrosName string
        Value      string
}

type UpdateCreativeInfo struct {
        CreativeNmb      int64
        CreativeName     string
        TnsAdvertiserNmb int64
        TemplateNmb      int64
        ExpireDate       string  // "20120802Т19:30:00"
        Tag              string
        Note             string
}

type UpdateCreativeGeoInfo struct {
        CreativeNmb int64
        GeoNmb      int64
        Exclude     bool
}

type UpdateCreativeMacrosInfo struct {
        CreativeNmb int64
        MacrosNmb   int64
        Value       string
}

type UpdateCreativeDynamicMacrosInfo struct {
        CreativeNmb int64
        MacrosNmb   int64
        Values      []string
}

type UpdateCreativeSiteInfo struct {
        CreativeNmb int64
        SiteNmb     int64
        Exclude     bool
}

type UpdateCreativeTnsArticleInfo struct {
        CreativeNmb int64
        Article     []int64
}

type UpdateCreativeTnsBrandInfo struct {
        CreativeNmb int64
        Brand       []int64
}

type UpdateCreativeSignedExpireDateInfo struct {
        CreativeNmb int64
        ExpireDate  string  // "20120802Т19:30:00"
}

type RequestCreativeEditInfo struct {
        CreativeNmb int64
}

type RequestCreativeModerationInfo struct {
        CreativeNmb int64
}
