package main

import (
        "github.com/gorilla/mux"
        "github.com/nu7hatch/gouuid"

        "dsp/bannerstore"

        "encoding/json"
        "fmt"
        "html"
        "log"
        "net/http"
)

// Bid request

type Banner struct {
        W, H int
}

type ImpressionObject struct {
        Id      string
        Banner  Banner
}

type Site struct {
        Id, Referer, Rereferer string
}

type Geo struct {
        Country, Region, City string
}

type Device struct {
        Ua, Ip, Userdata string
        Geo              Geo
}

type BidRequestMessage struct {
        Id      string
        Imp     []ImpressionObject
        Site    Site
        Device  Device
        Cur     []string
}

// Bid Response

type DSPParams struct {
        Url_param1  string  `json:"url_param1,omitempty"`
        Json_param1 string  `json:"json_param1,omitempty"`
}

type BidParameters struct {
        Id          string    `json:"id"`
        Adid        string    `json:"adid"`
        Price       int       `json:"price"`
        Adm         string    `json:"adm"`
        Properties  string    `json:"properties"`
        Token       string    `json:"token"`
        View_notice string    `json:"view_notice"`
        Dsp_params  DSPParams `json:"dsp_params"`
}

type Bid struct {
        Bid []BidParameters `json:"bid"`
}

type BidResponseMessage struct {
        Bidid       string  `json:"bidid"`
        Id          string  `json:"id"`
        Cur         string  `json:"cur"`
        Units       int     `json:"units"`
        Setuserdata string  `json:"units"`
        Bidset      []Bid   `json:"bidset"`
}

func rtb_host() {
        r := mux.NewRouter()
        r.HandleFunc("/", HomeHandler).Methods("GET")
        r.HandleFunc("/", DSPHandler).Methods("POST")

        http.Handle("/", r)

        fmt.Printf("Listening on port 8080...\n")
        log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func DSPHandler(w http.ResponseWriter, r *http.Request) {
        var m BidRequestMessage

        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&m)

        if err != nil {
                panic(err)
        }

        fmt.Printf("REQUEST %v\n===========\n%+v\n===========\n\n", m.Id, m)

        guuid, err := uuid.NewV4()
        response := BidResponseMessage{Bidid: guuid.String(), Id: m.Id, Cur: "USD", Units: 0, Setuserdata: "",
                Bidset: []Bid{{
                        Bid: []BidParameters{{
                                Id:     "", Adid: "", Price: 1, Adm: "", Properties: "", Token: "", View_notice: "",
                                Dsp_params: DSPParams{Url_param1: "", Json_param1: ""}}}}}}

        fmt.Printf("RESPONSE %v\n===========\n%+v\n===========\n\n", response.Id, response)

        err = json.NewEncoder(w).Encode(&response)

        if err != nil {
                panic(err)
        }

}

func main() {
        store, err := bannerstore.NewBannerStore("https://bayan2cdn.xmlrpc.http.yandex.net:35999", "sber-labs", "EADCA566-A8BF-403A-950A-0B82B526D2D1")

        if err != nil {
                panic(err)
        }

        logon := store.CreateLogon()
        fmt.Printf("logonId:: %s\n", logon)

        // store.GetSite()
        map1 := bannerstore.MapGeoNameNmb(store.GetGeo())
        for k, v := range map1 {
                fmt.Printf("%v: %v\n", k, v)
        }
}
