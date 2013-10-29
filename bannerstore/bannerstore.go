package bannerstore

import (
        "github.com/kolo/xmlrpc"
        "mapstructure"
        "reflect"
        "time"
)

type BannerStore struct {
        client        *xmlrpc.Client
        url           string
        username      string
        password      string
        logonId       string
        lastLoginTime time.Time

        // Dictionaries
        GetGeo           func() []GeoInfo
        GetMacros        func() []MacrosInfo
        GetSite          func() []SiteInfo
        GetTemplate      func() []TemplateInfo
        GetTnsAdvertiser func() []TnsAdvertiserInfo
        GetTnsArticle    func() []TnsArticleInfo
        GetTnsBrand      func() []TnsBrandInfo

        // Creatives
        CreateCreative                 func(CreateCreativeInfo) int64
        GetCreativeByNmb               func(int64) *CreativeInfo
        GetCreativeByTag               func(string) []CreativeInfo
        GetCreativeMacros              func(int64) []CreativeMacrosInfo
        GetCreativeTnsArticle          func(int64) []TnsArticleInfo
        GetCreativeTnsBrand            func(int64) []TnsBrandInfo
        UpdateCreative                 func(UpdateCreativeInfo)
        UpdateCreativeGeo              func(UpdateCreativeGeoInfo)
        UpdateCreativeMacros           func(UpdateCreativeMacrosInfo)
        UpdateCreativeDynamicMacros    func(UpdateCreativeDynamicMacrosInfo)
        UpdateCreativeSignedExpireDate func(UpdateCreativeSignedExpireDateInfo)
        UpdateCreativeSite             func(UpdateCreativeSiteInfo)
        UpdateCreativeTnsArticle       func(UpdateCreativeTnsArticleInfo)
        UpdateCreativeTnsBrand         func(UpdateCreativeTnsBrandInfo)

        // Files
        GetFileByNmb func(int64) *FileInfo
        GetFileByTag func(string) []FileInfo
        UploadFile   func(UploadFileInfo) int64

        // Moderation
        RequestCreativeEdit       func(RequestCreativeEditInfo)
        RequestCreativeModeration func(RequestCreativeModerationInfo)
}

func NewBannerStore(url, username, password string) (*BannerStore, error) {
        client, err := xmlrpc.NewClient(url, nil)
        store := BannerStore{client, url, username, password, "",
                time.Date(2009, time.November, 10, 15, 0, 0, 0, time.Local),
                nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
                nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}

        store.makeRequestFunc("GetGeo", &store.GetGeo)
        store.makeRequestFunc("GetMacros", &store.GetMacros)
        store.makeRequestFunc("GetSite", &store.GetSite)
        store.makeRequestFunc("GetTemplate", &store.GetTemplate)
        store.makeRequestFunc("GetTnsAdvertiser", &store.GetTnsAdvertiser)
        store.makeRequestFunc("GetTnsArticle", &store.GetTnsArticle)
        store.makeRequestFunc("GetTnsBrand", &store.GetTnsBrand)

        store.makeRequestFunc("CreateCreative", &store.CreateCreative)
        store.makeRequestFunc("GetCreativeByNmb", &store.GetCreativeByNmb)
        store.makeRequestFunc("GetCreativeByTag", &store.GetCreativeByTag)
        store.makeRequestFunc("GetCreativeMacros", &store.GetCreativeMacros)
        store.makeRequestFunc("GetCreativeByNmb", &store.GetCreativeByNmb)
        store.makeRequestFunc("GetCreativeTnsArticle", &store.GetCreativeTnsArticle)
        store.makeRequestFunc("GetCreativeTnsBrand", &store.GetCreativeTnsBrand)
        store.makeRequestFunc("UpdateCreative", &store.UpdateCreative)
        store.makeRequestFunc("UpdateCreativeGeo", &store.UpdateCreativeGeo)
        store.makeRequestFunc("UpdateCreativeMacros", &store.UpdateCreativeMacros)
        store.makeRequestFunc("UpdateCreativeDynamicMacros", &store.UpdateCreativeDynamicMacros)
        store.makeRequestFunc("UpdateCreativeSignedExpireDate", &store.UpdateCreativeSignedExpireDate)
        store.makeRequestFunc("UpdateCreativeSite", &store.UpdateCreativeSite)
        store.makeRequestFunc("UpdateCreativeTnsArticle", &store.UpdateCreativeTnsArticle)
        store.makeRequestFunc("UpdateCreativeTnsBrand", &store.UpdateCreativeTnsArticle)

        store.makeRequestFunc("GetFileByNmb", &store.GetFileByNmb)
        store.makeRequestFunc("GetFileByTag", &store.GetFileByTag)
        store.makeRequestFunc("UploadFile", &store.UploadFile)

        store.makeRequestFunc("RequestCreativeEdit", &store.RequestCreativeEdit)
        store.makeRequestFunc("RequestCreativeModeration", &store.RequestCreativeModeration)

        return &store, err
}

func (store *BannerStore) makeRequestFunc(req string, fptr interface{}) {
        baseRequestFunc := func(params []reflect.Value) []reflect.Value {
                var res []interface{}

                switch req {

                // Dictionaries
                case "GetGeo":
                        var retval []GeoInfo
                        store.client.Call("BannerStore."+req, nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}

                case "GetMacros":
                        var retval []MacrosInfo
                        store.client.Call("BannerStore."+req, nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}

                case "GetSite":
                        var retval []SiteInfo
                        store.client.Call("BannerStore."+req, nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}

                case "GetTemplate":
                        var retval []TemplateInfo
                        store.client.Call("BannerStore."+req, nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}

                case "GetTnsAdvertiser":
                        var retval []TnsAdvertiserInfo
                        store.client.Call("BannerStore."+req, nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}

                case "GetTnsArticle":
                        var retval []TnsArticleInfo
                        store.client.Call("BannerStore."+req, nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}

                case "GetTnsBrand":
                        var retval []TnsBrandInfo
                        store.client.Call("BannerStore."+req, nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}

                // Creatives
                case "CreateCreative":
                        var r int64
                        p := make([]interface{}, 2, 2)
                        p = append(p, store.CreateLogon())
                        p = append(p, params[0].Interface())
                        store.client.Call("BannerStore."+req, p, &r)
                        return []reflect.Value{reflect.ValueOf(r)}

                case "GetCreativeByNmb":
                        var retval CreativeInfo
                        p := make([]interface{}, 2, 2)
                        p = append(p, store.CreateLogon())
                        p = append(p, params[0].Interface())
                        store.client.Call("BannerStore."+req, p, &res)
                        mapstructure.Decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(&retval)}

                case "GetCreativeByTag":
                        var retval []CreativeInfo
                        p := make([]interface{}, 2, 2)
                        p = append(p, store.CreateLogon())
                        p = append(p, params[0].Interface())
                        store.client.Call("BannerStore."+req, p, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}

                case "GetCreativeMacros":
                        var retval []CreativeMacrosInfo
                        p := make([]interface{}, 2, 2)
                        p = append(p, store.CreateLogon())
                        p = append(p, params[0].Interface())
                        store.client.Call("BannerStore."+req, p, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}

                case "GetCreativeTnsArticle":
                        var retval []TnsArticleInfo
                        p := make([]interface{}, 2, 2)
                        p = append(p, store.CreateLogon())
                        p = append(p, params[0].Interface())
                        store.client.Call("BannerStore."+req, p, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}

                case "GetCreativeTnsBrand":
                        var retval []TnsBrandInfo
                        p := make([]interface{}, 2, 2)
                        p = append(p, store.CreateLogon())
                        p = append(p, params[0].Interface())
                        store.client.Call("BannerStore."+req, p, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}

                case "UpdateCreative",
                        "UpdateCreativeGeoInfo",
                        "UpdateCreativeMacrosInfo",
                        "UpdateCreativeDynamicMacrosInfo",
                        "UpdateCreativeSignedExpireDate",
                        "UpdateCreativeSite",
                        "UpdateCreativeTnsArticle",
                        "UpdateCreativeTnsBrand",
                        "RequestCreativeEdit",
                        "RequestCreativeModeration":
                        p := make([]interface{}, 2, 2)
                        p = append(p, store.CreateLogon())
                        p = append(p, params[0].Interface())
                        store.client.Call("BannerStore."+req, p, &res)
                        return nil

                // Files
                case "GetFileByNmb":
                        var retval FileInfo
                        p := make([]interface{}, 2, 2)
                        p = append(p, store.CreateLogon())
                        p = append(p, params[0].Interface())
                        store.client.Call("BannerStore."+req, p, &res)
                        mapstructure.Decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(&retval)}

                case "GetFileByTag":
                        var retval []FileInfo
                        p := make([]interface{}, 2, 2)
                        p = append(p, store.CreateLogon())
                        p = append(p, params[0].Interface())
                        store.client.Call("BannerStore."+req, p, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}

                case "UploadFile":
                        var r int64
                        p := make([]interface{}, 2, 2)
                        p = append(p, store.CreateLogon())
                        p = append(p, params[0].Interface())
                        store.client.Call("BannerStore."+req, p, &r)
                        return []reflect.Value{reflect.ValueOf(r)}

                default:
                        return nil
                }
        }

        fn := reflect.ValueOf(fptr).Elem()
        reqFun := reflect.MakeFunc(fn.Type(), baseRequestFunc)
        fn.Set(reqFun)
}

func decode(data interface{}, result interface{}) {
        val := reflect.ValueOf(result).Elem()
        dataVal := reflect.Indirect(reflect.ValueOf(data))

        valType := val.Type()
        valElemType := valType.Elem()

        // Make a new slice to hold our result, same size as the original data.
        sliceType := reflect.SliceOf(valElemType)
        valSlice := reflect.MakeSlice(sliceType, dataVal.Len(), dataVal.Len())

        for i := 0; i < dataVal.Len(); i++ {
                currentData := dataVal.Index(i).Interface()
                currentElem := valSlice.Index(i).Addr()
                mapstructure.DecodeIntoValue(currentData, currentElem)
        }

        val.Set(valSlice)
}

func (store *BannerStore) CreateLogon() string {
        if time.Now().After(store.lastLoginTime.Add(3 * time.Second)) {
                result := ""
                store.lastLoginTime = time.Now()
                store.client.Call("BannerStore.CreateLogon", xmlrpc.Struct{"name": store.username, "password": store.password}, &result)
                store.logonId = result
        }
        return store.logonId
}
