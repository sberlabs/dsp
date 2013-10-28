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

        GetGeo           func() []GeoInfo
        GetMacros        func() []MacrosInfo
        GetSite          func() []SiteInfo
        GetTemplate      func() []TemplateInfo
        GetTnsAdvertiser func() []TnsAdvertiserInfo
        GetTnsArticle    func() []TnsArticleInfo
        GetTnsBrand      func() []TnsBrandInfo
}

func NewBannerStore(url, username, password string) (*BannerStore, error) {
        client, err := xmlrpc.NewClient(url, nil)
        store := BannerStore{client, url, username, password, "",
                time.Date(2009, time.November, 10, 15, 0, 0, 0, time.Local),
                nil, nil, nil, nil, nil, nil, nil}

        store.makeRequestFunc("GetGeo", &store.GetGeo)
        store.makeRequestFunc("GetMacros", &store.GetMacros)
        store.makeRequestFunc("GetSite", &store.GetSite)
        store.makeRequestFunc("GetTemplate", &store.GetTemplate)
        store.makeRequestFunc("GetTnsAdvertiser", &store.GetTnsAdvertiser)
        store.makeRequestFunc("GetTnsArticle", &store.GetTnsArticle)
        store.makeRequestFunc("GetTnsBrand", &store.GetTnsBrand)

        return &store, err
}

func (store *BannerStore) makeRequestFunc(req string, fptr interface{}) {
        baseRequestFunc := func(params []reflect.Value) []reflect.Value {
                var res []interface{}

                switch req {
                // Dictionary methods, logonId is not required
                case "GetGeo":
                        var retval []GeoInfo
                        store.client.Call("BannerStore.GetGeo", nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}
                case "GetMacros":
                        var retval []MacrosInfo
                        store.client.Call("BannerStore.GetMacros", nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}
                case "GetSite":
                        var retval []SiteInfo
                        store.client.Call("BannerStore.GetSite", nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}
                case "GetTemplate":
                        var retval []TemplateInfo
                        store.client.Call("BannerStore.GetTemplate", nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}
                case "GetTnsAdvertiser":
                        var retval []TnsAdvertiserInfo
                        store.client.Call("BannerStore.GetTnsAdvertiser", nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}
                case "GetTnsArticle":
                        var retval []TnsArticleInfo
                        store.client.Call("BannerStore.GetTnsArticle", nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}
                case "GetTnsBrand":
                        var retval []TnsBrandInfo
                        store.client.Call("BannerStore.GetTnsBrand", nil, &res)
                        decode(res, &retval)
                        return []reflect.Value{reflect.ValueOf(retval)}
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
