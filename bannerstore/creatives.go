package bannerstore

// Yandex RTB BannerStore
//

import "mapstructure"

//import "fmt"

// Creatives

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

func (store *BannerStore) CreateCreative(data CreateCreativeInfo, result *int64) error {
        params := make([]interface{}, 2)
        params = append(params, store.CreateLogon())
        params = append(params, data)
        store.client.Call("BannerStore.CreateCreative", params, &result)

        return nil
}

type UpdateCreativeSignedExpireDateInfo struct {
        CreativeNmb int64
        ExpireDate  string  // "20120802Т19:30:00"
}

func (store *BannerStore) GetCreativeByNmb(nmb int64, result *CreativeInfo) error {
        var m interface{}

        params := make([]interface{}, 2)
        params = append(params, store.CreateLogon())
        params = append(params, nmb)
        store.client.Call("BannerStore.GetCreativeByNmb", params, &m)

        return mapstructure.Decode(m, result)
}

func (store *BannerStore) GetCreativeByTag(tag string, result []CreativeInfo) error {
        var m []interface{}

        params := make([]interface{}, 2)
        params = append(params, store.CreateLogon())
        params = append(params, tag)
        store.client.Call("BannerStore.GetCreativeByTag", params, &m)

        result = []CreativeInfo{}
        for i := range m {
                var s CreativeInfo
                err := mapstructure.Decode(m[i], &s)
                if err != nil {
                        return err
                }
                result = append(result, s)
        }

        return nil
}

func (store *BannerStore) GetCreativeMarcos(creativeNmb int64, result []CreativeMacrosInfo) error {
        var m []interface{}

        params := make([]interface{}, 2)
        params = append(params, store.CreateLogon())
        params = append(params, creativeNmb)
        store.client.Call("BannerStore.GetCreativeMacros", params, &m)

        result = []CreativeMacrosInfo{}
        for i := range m {
                var s CreativeMacrosInfo
                err := mapstructure.Decode(m[i], &s)
                if err != nil {
                        return err
                }
                result = append(result, s)
        }

        return nil
}

/*
func (store BannerStore) UpdateCreative(logonId string, data UpdateCreativeInfo) {

}

func (store BannerStore) UpdateCreativeSignedExpireDate(logonId string, data UpdateCreativeSignedExpireDateInfo) {

}

func (store BannerStore) UpdateCreativeMacros(logonId string, data UpdateCreativeMacrosInfo) {

}

func UpdateCreativeGeo(logonId string, data UpdateCreativeGeoInfo) {

}

func (store *BannerStore) UploadFile(data UploadFileInfo) int {
	*store.CreateLogon()

	return 0
}
*/
