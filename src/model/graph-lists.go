package model

type GraphListLatestTweetsTimeline struct {
	Data   *ListLatestTweetsTimelineData `json:"data"`
	Errors []*ErrorRsp                   `json:"errors"`
}

type ListLatestTweetsTimelineData struct {
	List *ListLatestTweetsTimelineDataList `json:"list"`
}

type ListLatestTweetsTimelineDataList struct {
	TweetsTimeline *ListLatestTweetsTimelineTweetsTimeline `json:"tweets_timeline"`
}

type ListLatestTweetsTimelineTweetsTimeline struct {
	Timeline *Timeline `json:"timeline"`
}
