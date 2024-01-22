package wx_request

type UploadMediaRequest struct {
	// 类型 枚举值有图片（image）、语音（voice）、视频（video）和缩略图（thumb）
	Type         string `json:"type"`
	Media        []byte `json:"media"`
	FileName     string `json:"fileName"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
}

type UploadMediaResponse struct {
	Errcode int         `json:"errcode"`
	MediaId interface{} `json:"media_id"`
	Url     interface{} `json:"url"`
}

type GetMediaListRequest struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Count  int    `json:"count"`
}

// GetMediaNewsListResponse 获取图文素材列表
type GetMediaNewsListResponse struct {
	Errcode    int `json:"errcode"`
	TotalCount int `json:"total_count"`
	ItemCount  int `json:"item_count"`
	Item       []struct {
		MediaId string `json:"media_id"`
		Content struct {
			NewsItem []struct {
				Title            string `json:"title"`
				ThumbMediaId     string `json:"thumb_media_id"`
				ShowCoverPic     int    `json:"show_cover_pic"`
				Author           string `json:"author"`
				Digest           string `json:"digest"`
				Content          string `json:"content"`
				Url              string `json:"url"`
				ContentSourceUrl string `json:"content_source_url"`
			} `json:"news_item"`
		} `json:"content"`
		UpdateTime string `json:"update_time"`
	} `json:"item"`
}

// GetMediaOtherListResponse 获取其它素材列表
type GetMediaOtherListResponse struct {
	Errcode    int         `json:"errcode"`
	TotalCount interface{} `json:"total_count"`
	ItemCount  interface{} `json:"item_count"`
	Item       []struct {
		MediaId    interface{} `json:"media_id"`
		Name       interface{} `json:"name"`
		UpdateTime interface{} `json:"update_time"`
		Url        interface{} `json:"url"`
	} `json:"item"`
}
