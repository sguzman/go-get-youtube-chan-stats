package structs

func point(str string) *string {
    if len(str) == 0 {
        return nil
    } else {
        return &str
    }
}

func points(str []string) *[]string {
    if len(str) == 0 {
        return nil
    } else {
        return &str
    }
}

func structToInsert(resp ItemType) InsertType {
    return InsertType{
        id: resp.Id,
        title: resp.Snippet.Title,
        description: point(resp.Snippet.Description),
        customUrl: point(resp.Snippet.CustomUrl),
        publishedAt: resp.Snippet.PublishedAt,
        defaultUrl: resp.Snippet.Thumbnails.Default.Url,
        defaultWidth: resp.Snippet.Thumbnails.Default.Width,
        defaultHeight: resp.Snippet.Thumbnails.Default.Height,
        mediumUrl: resp.Snippet.Thumbnails.Medium.Url,
        mediumWidth: resp.Snippet.Thumbnails.Medium.Width,
        mediumHeight: resp.Snippet.Thumbnails.Medium.Height,
        highUrl: resp.Snippet.Thumbnails.High.Url,
        highWidth: resp.Snippet.Thumbnails.High.Width,
        highHeight: resp.Snippet.Thumbnails.High.Height,
        localizedTitle: resp.Snippet.Localized.Title,
        localizedDescription: resp.Snippet.Localized.Description,
        uploadsPlaylist: resp.ContentDetails.RelatedPlaylists.Uploads,
        watchHistory: resp.ContentDetails.RelatedPlaylists.WatchHistory,
        watchLater: resp.ContentDetails.RelatedPlaylists.WatchLater,
        topicIds: resp.TopicDetails.TopicIds,
        topicCategories: resp.TopicDetails.TopicCategories,
        privacyStatus: resp.Status.PrivacyStatus,
        isLinked: resp.Status.IsLinked,
        longUploadsStatus: resp.Status.LongUploadsStatus,
        defaultTab: resp.BrandingSettings.Channel.DefaultTab,
        trackingAnalyticsAccountId: point(resp.BrandingSettings.Channel.TrackingAnalyticsAccountId),
        moderateComments: resp.BrandingSettings.Channel.ModerateComments,
        showRelatedChannels: resp.BrandingSettings.Channel.ShowRelatedChannels,
        showBrowseView: resp.BrandingSettings.Channel.ShowBrowseView,
        featuredChannelsTitle: resp.BrandingSettings.Channel.FeaturedChannelsTitle,
        featuredChannelsUrls: points(resp.BrandingSettings.Channel.FeaturedChannelsUrls),
        unsubscribedTrailer: resp.BrandingSettings.Channel.UnsubscribedTrailer,
        profileColor: resp.BrandingSettings.Channel.ProfileColor,
        bannerImageUrl: resp.BrandingSettings.Image.BannerImageUrl,
        bannerMobileImageUrl: resp.BrandingSettings.Image.BannerMobileImageUrl,
        bannerTabletLowImageUrl: resp.BrandingSettings.Image.BannerTabletLowImageUrl,
        bannerTabletImageUrl: resp.BrandingSettings.Image.BannerTabletImageUrl,
        bannerTabletHdImageUrl: resp.BrandingSettings.Image.BannerTabletHdImageUrl,
        bannerTabletExtraHdImageUrl: resp.BrandingSettings.Image.BannerTabletExtraHdImageUrl,
        bannerMobileLowImageUrl: resp.BrandingSettings.Image.BannerMobileLowImageUrl,
        bannerMobileMediumHdImageUrl: resp.BrandingSettings.Image.BannerMobileMediumHdImageUrl,
        bannerMobileHdImageUrl: resp.BrandingSettings.Image.BannerMobileHdImageUrl,
        bannerMobileExtraHdImageUrl: resp.BrandingSettings.Image.BannerMobileExtraHdImageUrl,
        bannerTvImageUrl: resp.BrandingSettings.Image.BannerTvImageUrl,
        bannerTvLowImageUrl: resp.BrandingSettings.Image.BannerTvLowImageUrl,
        bannerTvMediumImageUrl: resp.BrandingSettings.Image.BannerTvMediumImageUrl,
        bannerTvHighImageUrl: resp.BrandingSettings.Image.BannerTvHighImageUrl,,
    }
}
