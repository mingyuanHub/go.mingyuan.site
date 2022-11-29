package openrtb

type AdResponse struct {
	Id         string         `json:"id,omitempty"`
	SeatBid    []*SeatBid     `json:"seatbid,omitempty"`
	BidId      string         `json:"bidid,omitempty"`
	Cur        string         `json:"cur,omitempty"`
	CustomData string         `json:"customdata,omitempty"`
	Nbr        int            `json:"nbr,omitempty"`
	Ext        *AdResponseExt `json:"ext,omitempty"`
}

type SeatBid struct {
	Bid   []*Bid `json:"bid"`
	Seat  string `json:"seat,omitempty"`
	Group int    `json:"grup,omitempty"`
	Ext   *Ext   `json:"ext,omitempty"`
}

type AdResponseExt struct {
	TP *AdResponseExtTP `json:"tp,omitempty"`
}

type AdResponseExtTP struct {
	AppId        int `json:"app_id"`
	AdseatId     int `json:"adseat_id"`
	SegmentId    int `json:"segment_id"`
	BucketId     int `json:"bucket_id"`
	AspId        int `json:"asp_id"`
	DspAccountId int `json:"dsp_account_id"`
}

type Bid struct {
	Id             string         `json:"id"`
	ImpId          string         `json:"impid"`
	Price          float64        `json:"price"`
	NUrl           string         `json:"nurl,omitempty"`
	BUrl           string         `json:"burl,omitempty"`
	LUrl           string         `json:"lurl,omitempty"`
	Adm            string         `json:"adm"`
	ADid           string         `json:"adid,omitempty"`
	ADomain        []string       `json:"adomain,omitempty"`
	Bundle         string         `json:"bundle,omitempty"`
	IUrl           string         `json:"iurl,omitempty"`
	CId            string         `json:"cid,omitempty"`
	CrId           string         `json:"crid,omitempty"`
	Cat            []string       `json:"cat,omitempty"`
	Attr           []int             `json:"attr,omitempty"`
	Api            int               `json:"api,omitempty"`
	Protocol       int               `json:"protocol,omitempty"`
	QagMediaRating int               `json:"qagmediarating,omitempty"`
	DealId         string            `json:"dealid,omitempty"`
	W              int               `json:"w,omitempty"`
	H              int               `json:"h,omitempty"`
	WRatio         int               `json:"wratio,omitempty"`
	HRatio         int               `json:"hratio,omitempty"`
	Exp            int               `json:"exp,omitempty"`
	Ext            *AdResponseBidExt `json:"ext"`
}

type AdResponseBidExt struct {
	NUrl   []string `json:"nurl"`
	LUrl   []string `json:"lurl"`
	BUrl   []string `json:"burl"`
	ImpUrl []string `json:"impurl"`
	ClkUrl []string `json:"clkurl"`
}

//验证格式是否正确
func (resp *AdResponse) VerifyFormat() bool {
	if len(resp.SeatBid) == 0 ||
		len(resp.SeatBid[0].Bid) == 0 ||
		resp.SeatBid[0].Bid[0].Price == 0 ||
		len(resp.SeatBid[0].Bid[0].Adm) == 0 {
		return false
	}

	return true
}


func NewAdResponse() *AdResponse {
	adResponse := &AdResponse{
		Id: "",
		SeatBid: []*SeatBid{
			{
				Bid: []*Bid{
					{
						Id:             "",
						ImpId:          "",
						Price:          0,
						NUrl:           "",
						BUrl:           "",
						LUrl:           "",
						Adm:            "",
						ADid:           "",
						ADomain:        []string{},
						Bundle:         "",
						IUrl:           "",
						CId:            "",
						CrId:           "",
						Cat:            []string{},
						Attr:           []int{},
						Api:            0,
						Protocol:       0,
						QagMediaRating: 0,
						DealId:         "",
						W:              0,
						H:              0,
						WRatio:         0,
						HRatio:         0,
						Exp:            0,
						Ext:            &AdResponseBidExt{},
					},
				},
				Seat:  "",
				Group: 0,
				Ext:   &Ext{},
			},
		},
		BidId:      "",
		Cur:        "",
		CustomData: "",
		Nbr:        0,
		Ext:        &AdResponseExt{},
	}
	return adResponse
}