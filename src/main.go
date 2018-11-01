package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "github.com/imroc/req"
    _ "github.com/lib/pq"
    "math/rand"
    "os"
    "runtime"
    "strings"
    "time"
)

const (
    defaultHost = "192.168.1.63"
    defaultPort = "30000"
)

func connStr() string {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")

    if len(host) == 0 || len(port) == 0 {
        return fmt.Sprintf("user=postgres dbname=youtube host=%s port=%s sslmode=disable", defaultHost, defaultPort)
    } else {
        return fmt.Sprintf("user=postgres dbname=youtube host=%s port=%s sslmode=disable", host, port)
    }
}

type LocalizedType struct {
    Title string `json:"title"`
    Description string `json:"description"`
}

type ThumbnailType struct {
    Url string `json:"url"`
    Width uint64 `json:"width"`
    Height uint64 `json:"height"`
}

type ThumbnailsType struct {
    Default ThumbnailType `json:"default"`
    Medium  ThumbnailType `json:"medium"`
    High    ThumbnailType `json:"high"`
}

type SnippetType struct {
    Title string `json:"title"`
    Description string `json:"description"`
    CustomUrl string `json:"customUrl"`
    PublishedAt string `json:"publishedAt"`
    Thumbnails ThumbnailsType `json:"thumbnails"`
    Localized LocalizedType `json:"localized"`
}

type RelatedPlaylistsType struct {
    Uploads string `json:"uploads"`
    WatchHistory string `json:"watchHistory"`
    WatchLater string `json:"watchLater"`
}

type ContentDetailsType struct {
    RelatedPlaylists RelatedPlaylistsType `json:"relatedPlaylists"`
}

type TopicDetailsType struct {
    TopicIds []string `json:"topicIds"`
    TopicCategories []string `json:"topicCategories"`
}

type StatusType struct {
    PrivacyStatus string `json:"privacyStatus"`
    IsLinked bool `json:"isLinked"`
    LongUploadsStatus string `json:"longUploadsStatus"`
}

type ChannelType struct {
    Title string `json:"title"`
    Description string `json:"description"`
    DefaultTab string `json:"defaultTab"`
    TrackingAnalyticsAccountId string `json:"trackingAnalyticsAccountId"`
    ModerateComments bool `json:"moderateComments"`
    ShowRelatedChannels bool `json:"showRelatedChannels"`
    ShowBrowseView bool `json:"showBrowseView"`
    FeaturedChannelsTitle string `json:"featuredChannelsTitle"`
    FeaturedChannelsUrls []string `json:"featuredChannelsUrls"`
    UnsubscribedTrailer string `json:"unsubscribedTrailer"`
    ProfileColor string `json:"profileColor"`
}

type ImageType struct {
    BannerImageUrl string `json:"bannerImageUrl"`
    BannerMobileImageUrl string `json:"bannerMobileImageUrl"`
    BannerTabletLowImageUrl string `json:"bannerTabletLowImageUrl"`
    BannerTabletImageUrl string `json:"bannerTabletImageUrl"`
    BannerTabletHdImageUrl string `json:"bannerTabletHdImageUrl"`
    BannerTabletExtraHdImageUrl string `json:"bannerTabletExtraHdImageUrl"`
    BannerMobileLowImageUrl string `json:"bannerMobileLowImageUrl"`
    BannerMobileMediumHdImageUrl string `json:"bannerMobileMediumHdImageUrl"`
    BannerMobileHdImageUrl string `json:"bannerMobileHdImageUrl"`
    BannerMobileExtraHdImageUrl string `json:"bannerMobileExtraHdImageUrl"`
    BannerTvImageUrl string `json:"bannerTvImageUrl"`
    BannerTvLowImageUrl string `json:"bannerTvLowImageUrl"`
    BannerTvMediumImageUrl string `json:"bannerTvMediumImageUrl"`
    BannerTvHighImageUrl string `json:"bannerTvHighImageUrl"`
}

type HintType struct {
    Property string `json:"property"`
    Value string `json:"value"`
}

type BrandingSettingsType struct {
    Channel ChannelType `json:"channel"`
    Image ImageType `json:"image"`
    Hints []HintType `json:"hints"`
}

type ItemType struct {
    Kind string `json:"kind"`
    Etag string `json:"etag"`
    Id string `json:"id"`
    Snippet SnippetType `json:"snippet"`
    ContentDetails ContentDetailsType `json:"contentDetails"`
    TopicDetails TopicDetailsType `json:"topicDetails"`
    Status StatusType `json:"status"`
    BrandingSettings BrandingSettingsType `json:"brandingSettings"`
}

type PageInfoType struct {
    TotalResults uint64 `json:"totalResults"`
    ResultsPerPage uint64 `json:"resultsPerPage"`
}

type ResponseType struct {
    Kind string `json:"kind"`
    Etag string `json:"etag"`
    PageInfo PageInfoType `json:"pageInfo"`
    Items []ItemType `json:"items"`
}

func connection() *sql.DB {
    db, err := sql.Open("postgres", connStr())
    if err != nil {
        panic(err)
    }

    return db
}

func channels() []string {
    sqlStr := "select serial from youtube.entities.channels LIMIT 50"
    db := connection()
    defer func() {
        err := db.Close()
        if err != nil {
            panic(err)
        }
    }()

    row, err := db.Query(sqlStr)
    if err != nil {
        panic(err)
    }

    serials := make([]string, 0)
    for row.Next() {
        var serial string

        err = row.Scan(&serial)
        if err != nil {
            panic(err)
        }

        serials = append(serials, serial)
    }

    return serials
}

func getKey() string {
    rawKey := os.Getenv("API_KEY")
    splitKeys := strings.Split(rawKey, "|")

    return splitKeys[rand.Intn(len(splitKeys))]
}

func getJson(cs []string) ResponseType {
    key := getKey()
    url := "https://www.googleapis.com/youtube/v3/channels"
    partStr := "snippet,contentDetails,brandingSettings,contentOwnerDetails,invideoPromotion,localizations,status,topicDetails"
    idStr := strings.Join(cs, ",")

    param := req.Param{
        "part":  partStr,
        "id": idStr,
        "key": key,
    }

    r, err := req.Get(url, param)
    if err != nil {
        panic(err)
    }

    str, err := r.ToBytes()
    if err != nil {
        panic(err)
    }

    var data ResponseType
    err = json.Unmarshal(str, &data)
    if err != nil {
        panic(err)
    }

    return data
}

func main() {
    rand.Seed(time.Now().Unix())
    for {
        chans := channels()
        data := getJson(chans)
        fmt.Println(data)

        runtime.GC()
    }
}
