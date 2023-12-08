package model

type GraphUserTweets struct {
	Data   *UserTweetsData `json:"data"`
	Errors []*ErrorRsp     `json:"errors"`
}

type UserTweetsData struct {
	User *UserTweetsDataDetail `json:"user"`
}

type UserTweetsDataDetail struct {
	Result *UserTweetsDataDetailResult `json:"result"`
}

type UserTweetsDataDetailResult struct {
	Typename   string                                `json:"__typename"`
	TimelineV2 *UserTweetsDataDetailResultTimelineV2 `json:"timeline_v2"`
}

type UserTweetsDataDetailResultTimelineV2 struct {
	Timeline *Timeline `json:"timeline"`
}
