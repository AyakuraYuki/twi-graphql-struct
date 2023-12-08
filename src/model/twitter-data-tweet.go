package model

import "encoding/json"

const (
	TimelineClearCache  = "TimelineClearCache"  // 清理节点的标记，指向 Instructions[0].Type
	TimelinePinEntry    = "TimelinePinEntry"    // 时间线上指用户置顶的推文，指向 Instructions[x].Entry
	TimelineAddEntities = "TimelineAddEntities" // 时间线上的所有数据，都在这个类型的指向 Instructions[x].Entries[] 里

	EntryIdTweetPrefix        = "tweet-"
	EntryIdCursorTopPrefix    = "cursor-top-"
	EntryIdCursorBottomPrefix = "cursor-bottom-"
)

type TweetResults struct {
	Result *TweetResultsDetail `json:"result"`
}

type TweetResultsDetail struct {
	Typename                string                   `json:"__typename"`
	RestId                  string                   `json:"rest_id"`
	Core                    *TweetResultsResultCore  `json:"core"`
	EditControl             *EditControl             `json:"edit_control"`
	IsTranslatable          bool                     `json:"is_translatable"`
	Views                   *Views                   `json:"views"`
	Source                  string                   `json:"source"`
	Legacy                  *TweetLegacy             `json:"legacy"`
	QuotedStatusResult      *QuotedStatusResult      `json:"quoted_status_result,omitempty"`
	PreviousCounts          *GraphPublicMetrics      `json:"previous_counts,omitempty"`
	UnmentionData           json.RawMessage          `json:"unmention_data"`
	QuickPromoteEligibility *QuickPromoteEligibility `json:"quick_promote_eligibility,omitempty"`
}

type TweetResultsResultCore struct {
	UserResults *UserResult `json:"user_results"`
}

// ---------------------------------------------------------------------------------------------------- //
// 推文

type TweetLegacy struct {
	BookmarkCount             int                    `json:"bookmark_count"`
	Bookmarked                bool                   `json:"bookmarked"`
	CreatedAt                 string                 `json:"created_at"`
	ConversationIdStr         string                 `json:"conversation_id_str"`
	DisplayTextRange          []int                  `json:"display_text_range"`
	Entities                  *MediaEntities         `json:"entities,omitempty"`
	ExtendedEntities          *MediaEntities         `json:"extended_entities,omitempty"`
	FavoriteCount             int                    `json:"favorite_count"`
	Favorited                 bool                   `json:"favorited"`
	FullText                  string                 `json:"full_text"`
	IsQuoteStatus             bool                   `json:"is_quote_status"`
	Lang                      string                 `json:"lang"`
	PossiblySensitive         bool                   `json:"possibly_sensitive,omitempty"`
	PossiblySensitiveEditable bool                   `json:"possibly_sensitive_editable,omitempty"`
	QuoteCount                int                    `json:"quote_count"`
	ReplyCount                int                    `json:"reply_count"`
	RetweetCount              int                    `json:"retweet_count"`
	Retweeted                 bool                   `json:"retweeted"`
	UserIdStr                 string                 `json:"user_id_str"`
	IdStr                     string                 `json:"id_str"`
	QuotedStatusIdStr         string                 `json:"quoted_status_id_str,omitempty"`
	QuotedStatusPermalink     *QuotedStatusPermalink `json:"quoted_status_permalink,omitempty"`
	RetweetedStatusResult     *RetweetedStatusResult `json:"retweeted_status_result,omitempty"`
}

type QuotedStatusPermalink struct {
	Url      string `json:"url"`
	Expanded string `json:"expanded"`
	Display  string `json:"display"`
}

// ---------------------------------------------------------------------------------------------------- //
// 转推

type RetweetedStatusResult struct {
	Result *RetweetedStatusResultDetail `json:"result"`
}

type RetweetedStatusResultDetail struct {
	Typename           string                  `json:"__typename"`
	RestId             string                  `json:"rest_id"`
	Core               *TweetResultsResultCore `json:"core"`
	EditControl        *EditControl            `json:"edit_control"`
	IsTranslatable     bool                    `json:"is_translatable"`
	Views              *Views                  `json:"views"`
	Source             string                  `json:"source"`
	Legacy             *TweetLegacy            `json:"legacy"`
	PreviousCounts     *GraphPublicMetrics     `json:"previous_counts,omitempty"`
	QuotedStatusResult *QuotedStatusResult     `json:"quoted_status_result,omitempty"`
}

// ---------------------------------------------------------------------------------------------------- //
// 引用

type QuotedStatusResult struct {
	Result *QuotedStatusResultDetail `json:"result"`
}

type QuotedStatusResultDetail struct {
	Typename        string                  `json:"__typename"`
	RestId          string                  `json:"rest_id"`
	Core            *TweetResultsResultCore `json:"core"`
	EditControl     *EditControl            `json:"edit_control"`
	IsTranslatable  bool                    `json:"is_translatable"`
	Views           *Views                  `json:"views"`
	Source          string                  `json:"source"`
	NoteTweet       *NoteTweet              `json:"note_tweet,omitempty"`
	QuotedRefResult *QuotedRefResult        `json:"quotedRefResult,omitempty"`
	Legacy          *TweetLegacy            `json:"legacy"`
}

type QuotedRefResult struct {
	Result *QuotedRefResultDetail `json:"result"`
}

type QuotedRefResultDetail struct {
	Typename string `json:"__typename"`
	RestId   string `json:"rest_id"`
}

// ---------------------------------------------------------------------------------------------------- //

type NoteTweet struct {
	IsExpandable     bool              `json:"is_expandable"`
	NoteTweetResults *NoteTweetResults `json:"note_tweet_results"`
}

type NoteTweetResults struct {
	Result *NoteTweetResultsResult `json:"result"`
}

type NoteTweetResultsResult struct {
	Id        string           `json:"id"`
	Text      string           `json:"text"`
	EntitySet *MediaEntities   `json:"entity_set"`
	Richtext  *Richtext        `json:"richtext"`
	Media     *InlineMediaData `json:"media"`
}
