package model

type GraphTweetDetail struct {
	Data *TweetDetailData `json:"data"`
}

type TweetDetailData struct {
	ThreadedConversationWithInjectionsV2 *Timeline `json:"threaded_conversation_with_injections_v2"`
}

type GraphTweetResultByRestId struct {
	Data *TweetResultByRestIdData `json:"data"`
}

type TweetResultByRestIdData struct {
	TweetResults *TweetResults `json:"tweetResult"`
}
