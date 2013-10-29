package bannerstore

type FileInfo struct {
        FileNmb                 int64
        FileName                string
        Data                    []byte
        Tag                     string
        CdnNmb                  int64
        CdnUrl                  string
        Height                  int64
        Width                   int64
        Size                    int64
        MimeTypeNmb             int64
        FileModerationStatusNmb int64
        FileModeratedDate       string
        FileModeratedInfo       string
        Version                 int64
        CdnRequestData          string
        CdnResponseDate         string
}

type UploadFileInfo struct {
        CdnNmb         int64
        CdnRequestData string
        FileName       string
        Bytes          []byte
        Tag            string
}
