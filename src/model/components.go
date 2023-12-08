package model

import (
	"encoding/json"
	"github.com/spf13/cast"
)

type Url struct {
	DisplayUrl  string `json:"display_url,omitempty"`
	ExpandedUrl string `json:"expanded_url,omitempty"`
	Url         string `json:"url,omitempty"`
	Indices     []int  `json:"indices,omitempty"`

	Type    string `json:"type,omitempty"`
	UrlType string `json:"urlType,omitempty"`
}

type UserMention struct {
	IdStr      string `json:"id_str"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Indices    []int  `json:"indices"`
}

type Hashtag struct {
	Indices []int  `json:"indices"`
	Text    string `json:"text"`
}

type Symbol struct {
	Indices []int  `json:"indices"`
	Text    string `json:"text"`
}

type Views struct {
	Count string `json:"count,omitempty"`
	State string `json:"state,omitempty"`
}

func (v *Views) CountNum() int64 {
	if v.Count == "" {
		return 0
	}
	return cast.ToInt64(v.Count)
}

type EditControl struct {
	EditTweetIds       []string `json:"edit_tweet_ids,omitempty"`
	EditableUntilMsecs string   `json:"editable_until_msecs,omitempty"`
	IsEditEligible     bool     `json:"is_edit_eligible,omitempty"`
	EditsRemaining     string   `json:"edits_remaining,omitempty"`

	InitialTweetId     string          `json:"initial_tweet_id,omitempty"`
	EditControlInitial json.RawMessage `json:"edit_control_initial,omitempty"` // 仅包含 EditControl 的结构去除 EditControlInitial 后的其他字段
}

// ---------------------------------------------------------------------------------------------------- //

type ClientEventInfo struct {
	Component string                  `json:"component"`
	Element   string                  `json:"element,omitempty"`
	Details   *ClientEventInfoDetails `json:"details,omitempty"`
}

type ClientEventInfoDetails struct {
	TimelinesDetails *TimelinesDetails `json:"timelinesDetails,omitempty"`
}

type TimelinesDetails struct {
	InjectionType  string `json:"injectionType,omitempty"`
	ControllerData string `json:"controllerData,omitempty"`
	SourceData     string `json:"sourceData,omitempty"`
}

type MetadataBody struct {
	ScribeConfig         *ScribeConfig         `json:"scribeConfig,omitempty"`
	ConversationMetadata *ConversationMetadata `json:"conversationMetadata,omitempty"`
}

type ScribeConfig struct {
	Page string `json:"page"`
}

type ConversationMetadata struct {
	AllTweetIds         []string `json:"allTweetIds"`
	EnableDeduplication bool     `json:"enableDeduplication"`
}

// ---------------------------------------------------------------------------------------------------- //

type Richtext struct {
	RichtextTags []*RichtextTag `json:"richtext_tags,omitempty"`
}

type RichtextTag struct {
	FromIndex     int      `json:"from_index,omitempty"`
	ToIndex       int      `json:"to_index,omitempty"`
	RichtextTypes []string `json:"richtext_types,omitempty"`
}

type TextReferenceComponent struct {
	DisplayType string              `json:"displayType,omitempty"`
	Text        string              `json:"text,omitempty"`
	Sticky      bool                `json:"sticky,omitempty"`
	LandingUrl  *Url                `json:"landingUrl,omitempty"`
	Entities    []*ReferenceElement `json:"entities"`
}

type ReferenceElement struct {
	FromIndex int  `json:"fromIndex,omitempty"`
	ToIndex   int  `json:"toIndex,omitempty"`
	Ref       *Url `json:"ref,omitempty"`
}

type GraphPublicMetrics struct {
	BookmarkCount int `json:"bookmark_count,omitempty"`
	FavoriteCount int `json:"favorite_count,omitempty"`
	QuoteCount    int `json:"quote_count,omitempty"`
	ReplyCount    int `json:"reply_count,omitempty"`
	RetweetCount  int `json:"retweet_count,omitempty"`
}

type SocialContext struct {
	Type        string `json:"type"`
	ContextType string `json:"contextType"`
	Text        string `json:"text"`
}

type QuickPromoteEligibility struct {
	Eligibility string `json:"eligibility"`
}
