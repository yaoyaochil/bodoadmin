package xhs_request

type UserInfoByIDRequest struct {
	UserID                   string `json:"user_id"`
	ProfilePageHeadExpansion string `json:"profile_page_head_exp"`
}

type UserInfoByUserIDResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    struct {
		Collected               int    `json:"collected"`
		CollectedMovieNum       int    `json:"collected_movie_num"`
		RemarkName              string `json:"remark_name"`
		Imageb                  string `json:"imageb"`
		CollectedTagsNum        int    `json:"collected_tags_num"`
		CommunityRuleUrl        string `json:"community_rule_url"`
		IsRecommendLevelIllegal bool   `json:"is_recommend_level_illegal"`
		SellerInfo              struct {
			TabCodeNames       []string `json:"tab_code_names"`
			TabGoodsName       string   `json:"tab_goods_name"`
			IsTabGoodsFirst    bool     `json:"is_tab_goods_first"`
			CanProcessCoupon   bool     `json:"can_process_coupon"`
			TabGoodsApiVersion int      `json:"tab_goods_api_version"`
			StoreId            string   `json:"store_id"`
		} `json:"seller_info"`
		RedOfficialVerifyType int `json:"red_official_verify_type"`
		ShareInfoV2           struct {
			Content string `json:"content"`
			Title   string `json:"title"`
		} `json:"share_info_v2"`
		AvatarLikeStatus bool   `json:"avatar_like_status"`
		Fstatus          string `json:"fstatus"`
		CollectedBookNum int    `json:"collected_book_num"`
		ShareLink        string `json:"share_link"`
		TabPublic        struct {
			CollectionNote  bool `json:"collection_note"`
			CollectionBoard bool `json:"collection_board"`
			Seed            bool `json:"seed"`
			Collection      bool `json:"collection"`
		} `json:"tab_public"`
		DescAtUsers []interface{} `json:"desc_at_users"`
		NoteNumStat struct {
			Posted    int `json:"posted"`
			Liked     int `json:"liked"`
			Collected int `json:"collected"`
		} `json:"note_num_stat"`
		ViewerUserRelationShowTab bool   `json:"viewer_user_relation_show_tab"`
		ShowExtraInfoButton       bool   `json:"show_extra_info_button"`
		Desc                      string `json:"desc"`
		Gender                    int    `json:"gender"`
		RedOfficialVerified       bool   `json:"red_official_verified"`
		Result                    struct {
			Success bool   `json:"success"`
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"result"`
		ViewerUserRelationInfo struct {
			RelationInfo string   `json:"relation_info"`
			NickNames    []string `json:"nick_names"`
			Type         int      `json:"type"`
			HeadImage    []string `json:"head_image"`
		} `json:"viewer_user_relation_info"`
		RecommendInfo     string `json:"recommend_info"`
		RecommendInfoIcon string `json:"recommend_info_icon"`
		Interactions      []struct {
			Type      string `json:"type"`
			Name      string `json:"name"`
			Count     int    `json:"count"`
			IsPrivate bool   `json:"is_private"`
			Toast     string `json:"toast"`
		} `json:"interactions"`
		Blocking   bool   `json:"blocking"`
		IpLocation string `json:"ip_location"`
		Images     string `json:"images"`
		Level      struct {
			Image     string `json:"image"`
			ImageLink string `json:"image_link"`
			LevelName string `json:"level_name"`
			Number    int    `json:"number"`
		} `json:"level"`
		Location         string `json:"location"`
		ZhongTongBarInfo struct {
			Conversions []interface{} `json:"conversions"`
		} `json:"zhong_tong_bar_info"`
		CollectedNotesNum   int `json:"collected_notes_num"`
		CollectedProductNum int `json:"collected_product_num"`
		TabVisible          struct {
			Seed     bool `json:"seed"`
			Curation bool `json:"curation"`
			Note     bool `json:"note"`
			Collect  bool `json:"collect"`
			Like     bool `json:"like"`
			Goods    bool `json:"goods"`
		} `json:"tab_visible"`
		Blocked                  bool   `json:"blocked"`
		FeedbackAccountAppealUrl string `json:"feedback_account_appeal_url"`
		BannerInfo               struct {
			Image   string `json:"image"`
			BgColor string `json:"bg_color"`
		} `json:"banner_info"`
		Nboards                   int    `json:"nboards"`
		CollectedBrandNum         int    `json:"collected_brand_num"`
		Fans                      int    `json:"fans"`
		LocationJump              bool   `json:"location_jump"`
		RedOfficialVerifyBaseInfo string `json:"red_official_verify_base_info"`
		IdentityDeeplink          string `json:"identity_deeplink"`
		Ndiscovery                int    `json:"ndiscovery"`
		RedClubInfo               struct {
			RedClub      bool   `json:"red_club"`
			RedClubLevel int    `json:"red_club_level"`
			RedClubUrl   string `json:"red_club_url"`
			Redclubscore int    `json:"redclubscore"`
		} `json:"red_club_info"`
		ShareInfo struct {
			Title     string `json:"title"`
			Content   string `json:"content"`
			StoreLink string `json:"store_link"`
		} `json:"share_info"`
		DefaultCollectionTab     string `json:"default_collection_tab"`
		Nickname                 string `json:"nickname"`
		RedOfficialVerifyContent string `json:"red_official_verify_content"`
		Tags                     []struct {
			Icon    string `json:"icon"`
			Name    string `json:"name"`
			TagType string `json:"tag_type"`
		} `json:"tags"`
		Follows         int    `json:"follows"`
		RedId           string `json:"red_id"`
		Liked           int    `json:"liked"`
		CollectedPoiNum int    `json:"collected_poi_num"`
		Userid          string `json:"userid"`
	} `json:"data"`
	Code int `json:"code"`
}

type GetUserShopRequest struct {
	SellerId string `json:"seller_id"` // 店铺ID
	Sort     string `json:"sort"`      // 排序方式 新品：new_arrival 销量倒序:sales_qty
	Page     string `json:"page"`      // 页数
}

type GetUserShopResponse struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    struct {
		HasMore      bool `json:"has_more"`
		ListingItems []struct {
			SkuId          string `json:"sku_id"`
			Image          string `json:"image"`
			Buyable        bool   `json:"buyable"`
			TagStrategyMap struct {
				AfterPrice []struct {
					TagType    int `json:"tag_type"`
					TagContent struct {
						Content string `json:"content"`
					} `json:"tag_content"`
					Type string `json:"type"`
				} `json:"after_price,omitempty"`
				UponPrice []struct {
					TagType    int `json:"tag_type"`
					TagContent struct {
						FrameColorRgba        string  `json:"frameColorRgba"`
						LeftSpacing           int     `json:"leftSpacing"`
						FontSize              int     `json:"fontSize"`
						FontColor             string  `json:"fontColor"`
						FontColorRgba         string  `json:"fontColorRgba"`
						FrameColorDarkRgba    string  `json:"frameColorDarkRgba"`
						FrameDarkTransparency float64 `json:"frameDarkTransparency"`
						FrameTransparency     float64 `json:"frameTransparency"`
						FontDarkTransparency  float64 `json:"fontDarkTransparency"`
						FontTransparency      float64 `json:"fontTransparency"`
						FrameColor            string  `json:"frameColor"`
						FontColorDarkRgba     string  `json:"fontColorDarkRgba"`
						FontColorDark         string  `json:"fontColorDark"`
						TopSpacing            int     `json:"topSpacing"`
						Content               string  `json:"content"`
						FrameColorDark        string  `json:"frameColorDark"`
						LineHeight            int     `json:"lineHeight"`
					} `json:"tag_content"`
					Type string `json:"type"`
				} `json:"upon_price"`
				Behavior []struct {
					TagContent struct {
						Content       string `json:"content"`
						FontColor     string `json:"fontColor"`
						FontColorDark string `json:"fontColorDark"`
					} `json:"tag_content"`
					Type    string `json:"type"`
					TagType int    `json:"tag_type"`
				} `json:"behavior,omitempty"`
			} `json:"tag_strategy_map"`
			CardTitle          string `json:"card_title"`
			StockStatus        int    `json:"stock_status"`
			Height             int    `json:"height"`
			ItemId             string `json:"item_id"`
			Link               string `json:"link"`
			Width              int    `json:"width"`
			RecommendReasonTag struct {
				Icon            string `json:"icon"`
				IconHeight      int    `json:"icon_height"`
				IconWidth       int    `json:"icon_width"`
				RecommendReason string `json:"recommend_reason"`
			} `json:"recommend_reason_tag,omitempty"`
			Id          string `json:"id"`
			SellerId    string `json:"seller_id"`
			OnShelfTime int64  `json:"on_shelf_time"`
			PriceInfo   struct {
				ExpectedPrice struct {
					Price     float64 `json:"price"`
					PriceType int     `json:"price_type"`
				} `json:"expected_price"`
			} `json:"price_info"`
		} `json:"listing_items"`
	} `json:"data"`
}
