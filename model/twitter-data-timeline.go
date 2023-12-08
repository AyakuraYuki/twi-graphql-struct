package model

type Timeline struct {
	Instructions []*Instruction `json:"instructions"`
	Metadata     *MetadataBody  `json:"metadata"`
}

type Instruction struct {
	Type    string              `json:"type"`
	Entry   *InstructionEntry   `json:"entry"`
	Entries []*InstructionEntry `json:"entries"`
}

type InstructionEntry struct {
	EntryId   string                   `json:"entryId"`
	SortIndex string                   `json:"sortIndex"`
	Content   *InstructionEntryContent `json:"content"`
}

type InstructionEntryContent struct {
	EntryType       string                                 `json:"entryType"`
	Typename        string                                 `json:"__typename"`
	ItemContent     *ItemContent                           `json:"itemContent,omitempty"`
	Header          *TextReferenceComponent                `json:"header,omitempty"`
	Footer          *TextReferenceComponent                `json:"footer,omitempty"`
	Items           []*InstructionEntryContentItemsElement `json:"items,omitempty"`
	ClientEventInfo *ClientEventInfo                       `json:"clientEventInfo,omitempty"`
	Metadata        *MetadataBody                          `json:"metadata,omitempty"`
	DisplayType     string                                 `json:"displayType,omitempty"`
	Value           string                                 `json:"value,omitempty"`
	CursorType      string                                 `json:"cursorType,omitempty"`
}

type ItemContent struct {
	ItemType            string         `json:"itemType"`
	Typename            string         `json:"__typename"`
	TweetResults        *TweetResults  `json:"tweet_results"`
	TweetDisplayType    string         `json:"tweetDisplayType"`
	HasModeratedReplies bool           `json:"hasModeratedReplies,omitempty"`
	SocialContext       *SocialContext `json:"socialContext,omitempty"`
	UserResults         *UserResult    `json:"user_results,omitempty"`
	UserDisplayType     string         `json:"userDisplayType,omitempty"`
}

type InstructionEntryContentItemsElement struct {
	EntryId string                                   `json:"entryId"`
	Item    *InstructionEntryContentItemsElementItem `json:"item"`
}

type InstructionEntryContentItemsElementItem struct {
	ItemContent     *ItemContent     `json:"itemContent"`
	ClientEventInfo *ClientEventInfo `json:"clientEventInfo"`
}
