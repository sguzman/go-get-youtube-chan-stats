package structs

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