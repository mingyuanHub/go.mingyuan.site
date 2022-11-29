package openrtb

type AdRequest struct {
	Id     string   `json:"id"`
	Imp    []*Imp   `json:"imp"`
	App    *App     `json:"app,omitempty"`
	Device *Device  `json:"device,omitempty"`
	User   *User    `json:"user,omitempty"`
	Test   int      `json:"test,omitempty"`
	TMax   int64    `json:"tmax,omitempty"`
	Cur    []string `json:"cur,omitempty"`
	WLang  []string `json:"wlang,omitempty"`
	BCat   []string `json:"bcat,omitempty"`
	BAdv   []string `json:"badv,omitempty"`
	BApp   []string `json:"bapp,omitempty"`
	Source *Source  `json:"source,omitempty"`
	Regs   *Regs    `json:"regs,omitempty"`
	At     int      `json:"at,omitempty"`
	Ext    *Ext     `json:"ext,omitempty"`
}

type Imp struct {
	Id                string  `json:"id"`
	Banner            *Banner `json:"banner,omitempty"`
	Video             *Video  `json:"video,omitempty"`
	Native            *Native `json:"native,omitempty"`
	DisplayManager    string  `json:"displaymanager,omitempty"`
	DisplayManagerVer string  `json:"displaymanagerver,omitempty"`
	Instl             int     `json:"instl,omitempty"`
	TagId             string  `json:"tagid,omitempty"`
	BidFloor          float64 `json:"bidfloor"`
	BidFloorCur       string  `json:"bidfloorcur,omitempty"`
	ClickBrowser      int     `json:"clickbrowser,omitempty"`
	Secure            int     `json:"secure,omitempty"`
	Exp               int     `json:"exp,omitempty"`
	Ext               *ImpExt `json:"ext,omitempty"`
}

type Banner struct {
	Format   []*Format `json:"format,omitempty"`

	*AdSize

	BAttr    []int     `json:"battr,omitempty"`
	Pos      int       `json:"pos,omitempty"`
	Mimes    []string  `json:"mimes,omitempty"`
	TopFrame int       `json:"topframe,omitempty"`
	Api      []int     `json:"api,omitempty"`
	Id       string    `json:"id,omitempty"`
	Ext      *Ext      `json:"ext,omitempty"`
}

type Format struct {
	*AdSize

	WRatio int  `json:"wratio,omitempty"`
	HRatio int  `json:"hratio,omitempty"`
	WMin   int  `json:"wmin,omitempty"`
	Ext    *Ext `json:"ext,omitempty"`
}

type Video struct {
	Mimes          []string  `json:"mimes"`
	MinDuration    int       `json:"minduration"`
	MaxDuration    int       `json:"maxduration"`
	Protocols      []int     `json:"protocols"`

	*AdSize

	StartDelay     int       `json:"startdelay,omitempty"`
	Placement      int       `json:"placement,omitempty"`
	Linearity      int       `json:"linearity,omitempty"`
	Skip           int       `json:"skip,omitempty"`
	SkipMin        int       `json:"skipmin,omitempty"`
	SkipAfter      int       `json:"skipafter,omitempty"`
	Sequence       int       `json:"sequence,omitempty"`
	BAttr          []int     `json:"battr,omitempty"`
	MaxBitrate     int       `json:"maxbitrate,omitempty"`
	MinBitrate     int       `json:"mixbitrate,omitempty"`
	BoxingAllowed  int       `json:"boxingallowed,omitempty"`
	PlaybackMethod []int     `json:"playbackmethod,omitempty"`
	PlayBackend    int       `json:"playbackend,omitempty"`
	Delivery       []int     `json:"delivery,omitempty"`
	Pos            int       `json:"pos,omitempty"`
	CompanionAd    []*Banner `json:"companionad,omitempty"`
	Api            []int     `json:"api"`
	CompanionType  []int     `json:"companiontype,omitempty"`
	Ext            *VideoExt `json:"ext,omitempty"`
}

type VideoExt struct {
	Rewarded int `json:"rewarded,omitempty"`
}

type ImpExt struct {
	Skadn *ImpExtSkadn `json:"skadn,omitempty"`
}

type ImpExtSkadn struct {
	Version    string   `json:"version"`
	SourceApp  string   `json:"sourceApp"`
	SkadnetIds []string `json:"skadnetids"`
}

type Native struct {
	Request string `json:"request"`
	Ver     string `json:"ver,omitempty"`
	Api     []int  `json:"api,omitempty"`
	BAttr   []int  `json:"battr,omitempty"`
	Ext     *Ext   `json:"ext,omitempty"`
}

type App struct {
	Id            string   `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	Bundle        string   `json:"bundle,omitempty"`
	Domain        string   `json:"domain,omitempty"`
	StoreUrl      string   `json:"storeurl,omitempty"`
	Cat           []string `json:"cat,omitempty"`
	SectionCat    []string `json:"sectioncat,omitempty"`
	PageCat       []string `json:"pagecat,omitempty"`
	Ver           string   `json:"ver,omitempty"`
	Paid          string   `json:"paid,omitempty"`
	PrivacyPolicy int      `json:"privacypolicy,omitempty"`
	Keywords      string   `json:"keywords,omitempty"`
	Ext           *AppExt  `json:"ext,omitempty"`
}

type AppExt struct {
	Orientation int `json:"orientation,omitempty"`
}

type Device struct {
	Ua             string     `json:"ua,omitempty"`
	Geo            *Geo       `json:"geo,omitempty"`
	Lmt            int        `json:"lmt,omitempty"`
	Dnt            int        `json:"dnt,omitempty"`
	Ip             string     `json:"ip,omitempty"`
	Ipv6           string     `json:"ipv6,omitempty"`
	DeviceType     int        `json:"devicetype,omitempty"`
	Make           string     `json:"make,omitempty"`
	Model          string     `json:"model,omitempty"`
	Os             string     `json:"os,omitempty"`
	Osv            string     `json:"osv,omitempty"`
	Hwv            string     `json:"hwv,omitempty"`
	H              int        `json:"h,omitempty"`
	W              int        `json:"w,omitempty"`
	Ppi            int        `json:"ppi,omitempty"`
	PxRatio        float64    `json:"pxratio,omitempty"`
	Js             int        `json:"js,omitempty"`
	GeoFetch       int        `json:"geofetch,omitempty"`
	FlashVer       string     `json:"flashver,omitempty"`
	Language       string     `json:"language,omitempty"`
	Carrier        string     `json:"carrier,omitempty"`
	MccMnc         string     `json:"mccmnc,omitempty"`
	ConnectionType int        `json:"connectiontype,omitempty"`
	Ifa            string     `json:"ifa,omitempty"`
	Ext            *DeviceExt `json:"ext,omitempty"`
}

type Geo struct {
	Lat       float64 `json:"lat,omitempty"`
	Lon       float64 `json:"lon,omitempty"`
	Type      int     `json:"type,omitempty"`
	Country   string  `json:"country,omitempty"`
	Region    string  `json:"region,omitempty"`
	City      string  `json:"city,omitempty"`
	Zip       string  `json:"zip,omitempty"`
	Accuracy  int     `json:"accuracy,omitempty"`
	LastFix   int     `json:"lastfix,omitempty"`
	IpService int     `json:"ipservice,omitempty"`
	UtcOffset int     `json:"utcoffset,omitempty"`
}

type DeviceExt struct {
	Idfa string `json:"idfa,omitempty"`
	Idfv string `json:"idfv,omitempty"`
	Gaid string `json:"gaid,omitempty"`
	Atts int    `json:"atts,omitempty"`
}

type User struct {
	Id       string `json:"id,omitempty"`
	BuyerUid string `json:"buyerUid,omitempty"`
	Yob      int    `json:"yob,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Keywords string `json:"keywords,omitempty"`
	Ext      *Ext   `json:"ext,omitempty"`
}

type Source struct {
	Fd     int        `json:"fd,omitempty"`
	TId    string     `json:"tid,omitempty"`
	PChain string     `json:"pchain,omitempty"`
	Ext    *SourceExt `json:"ext,omitempty"`
}

type SourceExt struct {
	Schain *Schain `json:"schain"`
}

type Schain struct {
	Complete int           `json:"complete"`
	Ver      string        `json:"ver"`
	Nodes    []*SchainNode `json:"nodes"`
}

type SchainNode struct {
	Asi    string `json:"asi"`
	Sid    string `json:"sid"`
	Rid    string `json:"rid,omitempty"`
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain,omitempty"`
	Hp     int    `json:"hp"`
}

type Regs struct {
	Coppa int      `json:"coppa"`
	Ext   *RegsExt `json:"ext,omitempty"`
}

type RegsExt struct {
	Gdpr int `json:"gdpr"`
	Ccpa int `json:"ccpa"`
}

type Ext struct {
	TP *TP `json:"tp,omitempty"`
}

type TP struct {
	SegmentId      int                `json:"segmentid,omitempty"`
	BucketId       int                `json:"bucketid,omitempty"`
	Sdkv           string             `json:"sdkv,omitempty"`
	BidFloorConfig *AdxBidFloorConfig `json:"bidfloorconfig"`
}

type AdxBidFloorConfig struct {
	FloorPrice    float64            `json:"floor_price"`     //adx底价
	DspFloorPrice map[string]float64 `json:"dsp_floor_price"` //公司类型的dspAccountId和底价关系表
}

type AdSize struct {
	W int `json:"w"`
	H int `json:"h"`
}

