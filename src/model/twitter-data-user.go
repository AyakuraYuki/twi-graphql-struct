package model

import "encoding/json"

type UserResult struct {
	Result *RawUser `json:"result"`
}

type RawUser struct {
	Typename                        string                 `json:"__typename"`
	Legacy                          *UserLegacy            `json:"legacy"`
	RestId                          string                 `json:"rest_id"`
	Id                              string                 `json:"id"`
	AffiliatesHighlightedLabel      *AffiliatesHighlighted `json:"affiliates_highlighted_label"`
	IsBlueVerified                  bool                   `json:"is_blue_verified"`
	ProfileImageShape               string                 `json:"profile_image_shape"`
	LegacyExtendedProfile           json.RawMessage        `json:"legacy_extended_profile"`
	IsProfileTranslatable           bool                   `json:"is_profile_translatable"`
	HasGraduatedAccess              bool                   `json:"has_graduated_access,omitempty"`
	HasHiddenLikesOnProfile         bool                   `json:"has_hidden_likes_on_profile"`
	HasHiddenSubscriptionsOnProfile bool                   `json:"has_hidden_subscriptions_on_profile"`
	VerificationInfo                *VerificationInfo      `json:"verification_info"`
	HighlightsInfo                  *HighlightsInfo        `json:"highlights_info"`
	BusinessAccount                 json.RawMessage        `json:"business_account"`
	CreatorSubscriptionsCount       int                    `json:"creator_subscriptions_count"`
	Professional                    *Professional          `json:"professional,omitempty"`
	SuperFollowEligible             bool                   `json:"super_follow_eligible,omitempty"`
}

type UserLegacy struct {
	Following               bool            `json:"following,omitempty"`
	CanDm                   bool            `json:"can_dm,omitempty"`
	CanMediaTag             bool            `json:"can_media_tag,omitempty"`
	CreatedAt               string          `json:"created_at"`
	DefaultProfile          bool            `json:"default_profile"`
	DefaultProfileImage     bool            `json:"default_profile_image"`
	Description             string          `json:"description"`
	Entities                *UserEntities   `json:"entities"`
	FastFollowersCount      int             `json:"fast_followers_count"`
	FavouritesCount         int             `json:"favourites_count"`
	FollowersCount          int             `json:"followers_count"`
	FriendsCount            int             `json:"friends_count"`
	HasCustomTimelines      bool            `json:"has_custom_timelines"`
	IsTranslator            bool            `json:"is_translator"`
	ListedCount             int             `json:"listed_count"`
	Location                string          `json:"location"`
	MediaCount              int             `json:"media_count"`
	Name                    string          `json:"name"`
	NormalFollowersCount    int             `json:"normal_followers_count"`
	PinnedTweetIdsStr       []string        `json:"pinned_tweet_ids_str"`
	PossiblySensitive       bool            `json:"possibly_sensitive"`
	ProfileBannerUrl        string          `json:"profile_banner_url"`
	ProfileImageUrlHttps    string          `json:"profile_image_url_https"`
	ProfileInterstitialType string          `json:"profile_interstitial_type"`
	ScreenName              string          `json:"screen_name"`
	StatusesCount           int             `json:"statuses_count"`
	TranslatorType          string          `json:"translator_type"`
	Url                     string          `json:"url"`
	Verified                bool            `json:"verified"`
	VerifiedType            string          `json:"verified_type,omitempty"`
	WithheldInCountries     json.RawMessage `json:"withheld_in_countries"`
}

type UserEntities struct {
	Description *Urls `json:"description,omitempty"`
	Url         *Urls `json:"url,omitempty"`
}

type Urls struct {
	Urls []*Url `json:"urls,omitempty"`
}

type Professional struct {
	RestId           string        `json:"rest_id"`
	ProfessionalType string        `json:"professional_type"`
	Category         []interface{} `json:"category"`
}

type VerificationInfo struct {
	IsIdentityVerified bool                    `json:"is_identity_verified"`
	Reason             *VerificationInfoReason `json:"reason"`
}

type VerificationInfoReason struct {
	Description       *VerificationInfoReasonDescription `json:"description"`
	VerifiedSinceMsec string                             `json:"verified_since_msec"`
}

type VerificationInfoReasonDescription struct {
	Text     string                                     `json:"text"`
	Entities []*VerificationInfoReasonDescriptionEntity `json:"entities"`
}

type VerificationInfoReasonDescriptionEntity struct {
	FromIndex int                                         `json:"from_index"`
	ToIndex   int                                         `json:"to_index"`
	Ref       *VerificationInfoReasonDescriptionEntityRef `json:"ref"`
}

type VerificationInfoReasonDescriptionEntityRef struct {
	Url     string `json:"url"`
	UrlType string `json:"url_type"`
}

type HighlightsInfo struct {
	CanHighlightTweets bool   `json:"can_highlight_tweets"`
	HighlightedTweets  string `json:"highlighted_tweets"`
}

type AffiliatesHighlighted struct {
	Label *AffiliatesHighlightedLabel `json:"label"`
}

type AffiliatesHighlightedLabel struct {
	Url struct {
		Url     string `json:"url"`
		UrlType string `json:"urlType"`
	} `json:"url"`
	Badge struct {
		Url string `json:"url"`
	} `json:"badge"`
	Description          string `json:"description"`
	UserLabelType        string `json:"userLabelType"`
	UserLabelDisplayType string `json:"userLabelDisplayType"`
}
