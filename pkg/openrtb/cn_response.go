package openrtb

type DspAccessCnResponse struct {
	Id      string                `json:"id,omitempty"`
	SeatBid []*DspAccessCnSeatBid `json:"seatbid,omitempty"`
	BidId   string                `json:"bidid"`
	Nbr     int                   `json:"nbr,omitempty"`
}

type DspAccessCnSeatBid struct {
	BidCn []*BidCn `json:"bidcn,omitempty"`
}

type BidCn struct {
	Id           string         `json:"id,omitempty"`
	ImpId        string         `json:"impid,omitempty"`
	DspName      string         `json:"dsp_name,omitempty"`
	ADid         string         `json:"adid"`
	CId          string         `json:"cid"`
	Price        float64        `json:"price"`
	PriceCny     float64        `json:"price_cny,omitempty"`
	InteractType int            `json:"interact_type"`
	DealId       string         `json:"dealid"`
	Ad           *AdCn          `json:"ad"`
	Tracking     *TrackingCn    `json:"tracking"`
	Action       *ActionCn      `json:"action"`
	DownloadApp  *DownloadAppCn `json:"downloadapp"`
}

type AdCn struct {
	Adtype  int        `json:"adtype"`
	Title   string     `json:"title,omitempty"`
	Desc    string     `json:"desc,omitempty"`
	Cta     string     `json:"cta,omitempty"`
	From    string     `json:"from,omitempty"`
	Icon    *ImageCn   `json:"icon,omitempty"`
	AddLogo int        `json:"add_logo"`
	Logo    *ImageCn   `json:"logo,omitempty"`
	Images  []*ImageCn `json:"images,omitempty"`
	Video   *VideoCn   `json:"video,omitempty"`
	Html    *HtmlCn    `json:"html,omitempty"`
}

type ImageCn struct {
	Url string `json:"url"`
	W   int    `json:"w,omitempty"`
	H   int    `json:"h,omitempty"`
}

type VideoCn struct {
	Url         string `json:"url"`
	W           int    `json:"w,omitempty"`
	H           int    `json:"h,omitempty"`
	Duration    int    `json:"duration,omitempty"`
	FileSize    int    `json:"file_size,omitempty"`
	Cover       string `json:"cover,omitempty"`
	CoverW      int    `json:"cover_w,omitempty"`
	CoverH      int    `json:"cover_h,omitempty"`
	EndpageHtml string `json:"endpage_html,omitempty"`
	EndCover    string `json:"endcover,omitempty"`
}

type HtmlCn struct {
	Url     string `json:"url"`
	Content string `json:"content,omitempty"`
}

type TrackingCn struct {
	Em     *EmCn    `json:"em"`
	Nurl   []string `json:"nurl"`
	Lurl   []string `json:"lurl"`
	Impurl []string `json:"impurl"`
	Clkurl []string `json:"clkurl"`
}

type EmCn struct {
	DeeplinkAttempt   []string `json:"deeplink_attempt,omitempty"`
	AppInvokeSuccess  []string `json:"app_invoke_success,omitempty"`
	AppInvokeLoss     []string `json:"app_invoke_loss,omitempty"`
	WxInvokeAttempt   []string `json:"wx_invoke_attempt,omitempty"`
	WxInvokeSuccess   []string `json:"wx_invoke_success,omitempty"`
	VideoPlay         []string `json:"video_play,omitempty"`
	VideoOneQuarter   []string `json:"video_one_quarter,omitempty"`
	VideoOneHalf      []string `json:"video_one_half,omitempty"`
	VideoThreeQuarter []string `json:"video_three_quarter,omitempty"`
	VideoOver         []string `json:"video_over,omitempty"`
	DownloadStart     []string `json:"download_start,omitempty"`
	DownloadFinish    []string `json:"download_finish,omitempty"`
	InstallStart      []string `json:"install_start,omitempty"`
	InstallFinish     []string `json:"install_finish,omitempty"`
}

type ActionCn struct {
	OpMode         int    `json:"op_mode"`
	OpSlogan       string `json:"op_slogan"`
	LandingpageUrl string `json:"landingpage_url,omitempty"`
	DeeplinkUrl    string `json:"deeplink_url,omitempty"`
	Wxoid          string `json:"wxoid,omitempty"`
	Wxp            string `json:"wxp,omitempty"`
	StoreId        string `json:"store_id,omitempty"`
	UniversalLink  string `json:"universal_link,omitempty"`
	MarketUrl      string `json:"market_url,omitempty"`
	AppDownloadUrl string `json:"app_download_url,omitempty"`
	QuickAppUrl    string `json:"quick_app_url,omitempty"`
}

type DownloadAppCn struct {
	AppName             string `json:"app_name"`          //安卓应用名称
	AppIcon             string `json:"app_icon"`          //应⽤图⽚地址
	AdvertiserName      string `json:"advertiser_name"`   //安卓应用开发者名称
	AppVersion          string `json:"app_version"`       //应用版本
	PrivacyAgreementUrl string `json:"privacy_agreement"` //隐私协议超链
	PermissionsUrl      string `json:"permissions_url"`   //权限列表超链
	AppPrivacy          string `json:"app_privacy"`       //隐私权限内容

	PackageName      string `json:"package_name,omitempty"` //包名
	PackageSizeBytes int64  `json:"package_size_bytes"`     //int	安卓应用包大小
}