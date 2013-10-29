package main

import (
        . "dsp/bannerstore"
        "flag"
        "github.com/golang/glog"
        "labix.org/v2/mgo"
        //"labix.org/v2/mgo/bson"
)

var store *BannerStore
var session *mgo.Session
var geo, macros, site, template, tnsAdvertiser, tnsArticle, tnsBrand *mgo.Collection

func main() {
        var err error

        flag.Parse()
        glog.V(2).Infoln("Application starting...")

        session, err = mgo.Dial("localhost:27017")
        if err != nil {
                glog.Fatal(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)

        glog.V(2).Infoln("Connected to the database...")

        store, err = NewBannerStore("https://bayan2cdn.xmlrpc.http.yandex.net:35999",
                "sber-labs",
                "EADCA566-A8BF-403A-950A-0B82B526D2D1")
        if err != nil {
                glog.Fatal(err)
        }

        glog.V(2).Infoln("Connected to Yandex BannerStore...")

        logonId := store.CreateLogon()
        if logonId != "" {
                glog.V(2).Infof("  logonId: %s\n", logonId)
        } else {
                glog.Fatal("BannerStore authentication error.")
        }

        initCollections()
        ensureIndexes()
        //createAllDictionaries()
}

func initCollections() {
        geo = session.DB("creakl").C("geo")
        macros = session.DB("creakl").C("macros")
        site = session.DB("creakl").C("site")
        template = session.DB("creakl").C("template")
        tnsAdvertiser = session.DB("creakl").C("tnsadvertiser")
        tnsArticle = session.DB("creakl").C("tnsarticle")
        tnsBrand = session.DB("creakl").C("tnsbrand")
}

func ensureIndexes() {
        index1 := mgo.Index{
                Key:        []string{"nmb"},
                Unique:     true,
                DropDups:   true,
                Background: true,
                Sparse:     true,
        }
        index2 := mgo.Index{
                Key:        []string{"name"},
                Unique:     false,
                DropDups:   false,
                Background: true,
                Sparse:     true,
        }

        geo.EnsureIndex(index1)
        geo.EnsureIndex(index2)
        macros.EnsureIndex(index1)
        macros.EnsureIndex(index2)
        site.EnsureIndex(index1)
        site.EnsureIndex(index2)
        template.EnsureIndex(index1)
        template.EnsureIndex(index2)
        tnsAdvertiser.EnsureIndex(index1)
        tnsAdvertiser.EnsureIndex(index2)
        tnsArticle.EnsureIndex(index1)
        tnsArticle.EnsureIndex(index2)
        tnsBrand.EnsureIndex(index1)
        tnsBrand.EnsureIndex(index2)
}

func createAllDictionaries() {
        glog.V(2).Infoln("Updating All Dictionaries...")

        geoList := store.GetGeo()
        if len(geoList) == 0 {
                glog.Errorln("GeoInfo list is empty.")
        } else {
                for i := range geoList {
                        geo.Insert(&geoList[i])
                }
                glog.V(2).Infoln("  geo dictionary is updated.")
        }

        macrosList := store.GetMacros()
        if len(macrosList) == 0 {
                glog.Errorln("MacrosInfo list is empty.")
        } else {
                for i := range macrosList {
                        macros.Insert(&macrosList[i])
                }
                glog.V(2).Infoln("  macros dictionary is updated.")
        }

        siteList := store.GetSite()
        if len(siteList) == 0 {
                glog.Errorln("SiteInfo list is empty.")
        } else {
                for i := range siteList {
                        site.Insert(&siteList[i])
                }
                glog.V(2).Infoln("  site dictionary is updated.")
        }

        templateList := store.GetTemplate()
        if len(templateList) == 0 {
                glog.Errorln("TemplateInfo list is empty.")
        } else {
                for i := range templateList {
                        template.Insert(&templateList[i])
                }
                glog.V(2).Infoln("  template dictionary is updated.")
        }

        tnsAdvertiserList := store.GetTnsAdvertiser()
        if len(tnsAdvertiserList) == 0 {
                glog.Errorln("TnsAdvertiserInfo list is empty.")
        } else {
                for i := range tnsAdvertiserList {
                        tnsAdvertiser.Insert(&tnsAdvertiserList[i])
                }
                glog.V(2).Infoln("  tnsadvertiser dictionary is updated.")
        }

        tnsArticleList := store.GetTnsArticle()
        if len(tnsArticleList) == 0 {
                glog.Errorln("TnsArticleInfo list is empty.")
        } else {
                for i := range tnsArticleList {
                        tnsArticle.Insert(&tnsArticleList[i])
                }
                glog.V(2).Infoln("  tnsarticle is updated.")
        }

        tnsBrandList := store.GetTnsBrand()
        if len(tnsBrandList) == 0 {
                glog.Errorln("TnsBrandInfo list is empty.")
        } else {
                for i := range tnsBrandList {
                        tnsBrand.Insert(&tnsBrandList[i])
                }
                glog.V(2).Infoln("  tnsbrand dictionary is updated.")
        }
}
