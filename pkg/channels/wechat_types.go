package channels

type weChatBaseInfo struct {
	ChannelVersion string `json:"channel_version,omitempty"`
}

type weChatGetUpdatesRequest struct {
	GetUpdatesBuf string          `json:"get_updates_buf,omitempty"`
	BaseInfo      weChatBaseInfo  `json:"base_info"`
}

type weChatGetUpdatesResponse struct {
	Ret                  int              `json:"ret,omitempty"`
	ErrCode              int              `json:"errcode,omitempty"`
	ErrMsg               string           `json:"errmsg,omitempty"`
	Messages             []weChatMessage  `json:"msgs,omitempty"`
	GetUpdatesBuf        string           `json:"get_updates_buf,omitempty"`
	LongPollingTimeoutMS int              `json:"longpolling_timeout_ms,omitempty"`
}

type weChatCDNMedia struct {
	EncryptQueryParam string `json:"encrypt_query_param,omitempty"`
	AESKey            string `json:"aes_key,omitempty"`
	EncryptType       int    `json:"encrypt_type,omitempty"`
}

type weChatTextItem struct {
	Text string `json:"text,omitempty"`
}

type weChatImageItem struct {
	Media      *weChatCDNMedia `json:"media,omitempty"`
	ThumbMedia *weChatCDNMedia `json:"thumb_media,omitempty"`
	RawAESKey  string          `json:"aeskey,omitempty"`
	URL        string          `json:"url,omitempty"`
	MidSize    int             `json:"mid_size,omitempty"`
	ThumbSize  int             `json:"thumb_size,omitempty"`
	HDSize     int             `json:"hd_size,omitempty"`
}

type weChatVoiceItem struct {
	Media      *weChatCDNMedia `json:"media,omitempty"`
	EncodeType int             `json:"encode_type,omitempty"`
	SampleRate int             `json:"sample_rate,omitempty"`
	PlayTime   int             `json:"playtime,omitempty"`
	Text       string          `json:"text,omitempty"`
}

type weChatFileItem struct {
	Media    *weChatCDNMedia `json:"media,omitempty"`
	FileName string          `json:"file_name,omitempty"`
	MD5      string          `json:"md5,omitempty"`
	Len      string          `json:"len,omitempty"`
}

type weChatVideoItem struct {
	Media      *weChatCDNMedia `json:"media,omitempty"`
	ThumbMedia *weChatCDNMedia `json:"thumb_media,omitempty"`
	VideoSize  int             `json:"video_size,omitempty"`
	PlayLength int             `json:"play_length,omitempty"`
	VideoMD5   string          `json:"video_md5,omitempty"`
}

type weChatRefMessage struct {
	MessageItem *weChatMessageItem `json:"message_item,omitempty"`
	Title       string             `json:"title,omitempty"`
}

type weChatMessageItem struct {
	Type      int               `json:"type,omitempty"`
	TextItem  *weChatTextItem   `json:"text_item,omitempty"`
	ImageItem *weChatImageItem  `json:"image_item,omitempty"`
	VoiceItem *weChatVoiceItem  `json:"voice_item,omitempty"`
	FileItem  *weChatFileItem   `json:"file_item,omitempty"`
	VideoItem *weChatVideoItem  `json:"video_item,omitempty"`
	RefMsg    *weChatRefMessage `json:"ref_msg,omitempty"`
}

type weChatMessage struct {
	Sequence      int                 `json:"seq,omitempty"`
	MessageID     int64               `json:"message_id,omitempty"`
	FromUserID    string              `json:"from_user_id,omitempty"`
	ToUserID      string              `json:"to_user_id,omitempty"`
	ClientID      string              `json:"client_id,omitempty"`
	CreateTimeMS  int64               `json:"create_time_ms,omitempty"`
	SessionID     string              `json:"session_id,omitempty"`
	GroupID       string              `json:"group_id,omitempty"`
	MessageType   int                 `json:"message_type,omitempty"`
	MessageState  int                 `json:"message_state,omitempty"`
	ItemList      []weChatMessageItem `json:"item_list,omitempty"`
	ContextToken  string              `json:"context_token,omitempty"`
}

type weChatSendMessageRequest struct {
	Message weChatMessage `json:"msg"`
}

type weChatGetUploadURLRequest struct {
	FileKey         string `json:"filekey,omitempty"`
	MediaType       int    `json:"media_type,omitempty"`
	ToUserID        string `json:"to_user_id,omitempty"`
	RawSize         int    `json:"rawsize,omitempty"`
	RawFileMD5      string `json:"rawfilemd5,omitempty"`
	FileSize        int    `json:"filesize,omitempty"`
	ThumbRawSize    int    `json:"thumb_rawsize,omitempty"`
	ThumbRawFileMD5 string `json:"thumb_rawfilemd5,omitempty"`
	ThumbFileSize   int    `json:"thumb_filesize,omitempty"`
	NoNeedThumb     bool   `json:"no_need_thumb,omitempty"`
	AESKey          string `json:"aeskey,omitempty"`
	BaseInfo        weChatBaseInfo `json:"base_info"`
}

type weChatGetUploadURLResponse struct {
	UploadParam      string `json:"upload_param,omitempty"`
	ThumbUploadParam string `json:"thumb_upload_param,omitempty"`
}

type weChatGetConfigRequest struct {
	ILinkUserID string         `json:"ilink_user_id,omitempty"`
	ContextToken string        `json:"context_token,omitempty"`
	BaseInfo    weChatBaseInfo `json:"base_info"`
}

type weChatGetConfigResponse struct {
	Ret          int    `json:"ret,omitempty"`
	ErrMsg       string `json:"errmsg,omitempty"`
	TypingTicket string `json:"typing_ticket,omitempty"`
}

type weChatSendTypingRequest struct {
	ILinkUserID  string         `json:"ilink_user_id,omitempty"`
	TypingTicket string         `json:"typing_ticket,omitempty"`
	Status       int            `json:"status,omitempty"`
	BaseInfo     weChatBaseInfo `json:"base_info"`
}

type weChatQRCodeResponse struct {
	QRCode           string `json:"qrcode,omitempty"`
	QRCodeImageURL   string `json:"qrcode_img_content,omitempty"`
}

type weChatQRStatusResponse struct {
	Status      string `json:"status,omitempty"`
	BotToken    string `json:"bot_token,omitempty"`
	ILinkBotID  string `json:"ilink_bot_id,omitempty"`
	BaseURL     string `json:"baseurl,omitempty"`
	ILinkUserID string `json:"ilink_user_id,omitempty"`
}

const (
	weChatMessageTypeText  = 1
	weChatMessageTypeImage = 2
	weChatMessageTypeVoice = 3
	weChatMessageTypeFile  = 4
	weChatMessageTypeVideo = 5

	weChatUploadMediaImage = 1
	weChatUploadMediaVideo = 2
	weChatUploadMediaFile  = 3

	weChatMessageTypeBot   = 2
	weChatMessageStateDone = 2
)
