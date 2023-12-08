package model

type MediaEntities struct {
	Media        []*Media       `json:"media,omitempty"`
	UserMentions []*UserMention `json:"user_mentions,omitempty"`
	Urls         []*Url         `json:"urls,omitempty"`
	Hashtags     []*Hashtag     `json:"hashtags,omitempty"`
	Symbols      []*Symbol      `json:"symbols,omitempty"`
}

type Media struct {
	DisplayUrl           string                `json:"display_url"`
	ExpandedUrl          string                `json:"expanded_url"`
	IdStr                string                `json:"id_str"`
	Indices              []int                 `json:"indices"`
	MediaKey             string                `json:"media_key"`
	MediaUrlHttps        string                `json:"media_url_https"`
	Type                 string                `json:"type"`
	Url                  string                `json:"url"`
	AdditionalMediaInfo  *AdditionalMediaInfo  `json:"additional_media_info,omitempty"`
	ExtMediaAvailability *ExtMediaAvailability `json:"ext_media_availability"`
	Sizes                *Sizes                `json:"sizes"`
	OriginalInfo         *OriginalInfo         `json:"original_info"`
	VideoInfo            *VideoInfo            `json:"video_info,omitempty"`
	SourceStatusIdStr    string                `json:"source_status_id_str,omitempty"`
	SourceUserIdStr      string                `json:"source_user_id_str,omitempty"`
	Features             *Features             `json:"features,omitempty"`
	ExtAltText           string                `json:"ext_alt_text,omitempty"`
}

type AdditionalMediaInfo struct {
	Monetizable bool        `json:"monetizable"`
	SourceUser  *SourceUser `json:"source_user,omitempty"`
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description,omitempty"`
	Embeddable  bool        `json:"embeddable,omitempty"`
}

type SourceUser struct {
	UserResults *UserResult `json:"user_results"`
}

type ExtMediaAvailability struct {
	Status string `json:"status"`
}

type OriginalInfo struct {
	Height     int          `json:"height"`
	Width      int          `json:"width"`
	FocusRects []*FacesSpec `json:"focus_rects"`
}

type VideoInfo struct {
	AspectRatio    []int      `json:"aspect_ratio"`
	DurationMillis int        `json:"duration_millis,omitempty"`
	Variants       []*Variant `json:"variants"`
}

type Variant struct {
	Bitrate     int    `json:"bitrate,omitempty"`
	ContentType string `json:"content_type"`
	Url         string `json:"url"`
}

type Sizes struct {
	Large  *SizeSpec `json:"large"`
	Medium *SizeSpec `json:"medium"`
	Small  *SizeSpec `json:"small"`
	Thumb  *SizeSpec `json:"thumb"`
}

type Features struct {
	Large  *Feature `json:"large"`
	Medium *Feature `json:"medium"`
	Small  *Feature `json:"small"`
	Orig   *Feature `json:"orig"`
}

type Feature struct {
	Faces []*FacesSpec `json:"faces"`
}

type SizeSpec struct {
	H      int    `json:"h"`
	W      int    `json:"w"`
	Resize string `json:"resize"`
}

type FacesSpec struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

type InlineMediaData struct {
	InlineMedia []*InlineMedia `json:"inline_media,omitempty"`
}

type InlineMedia struct {
	MediaId string `json:"media_id"`
	Index   int    `json:"index"`
}
