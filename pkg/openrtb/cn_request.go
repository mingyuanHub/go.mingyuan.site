package openrtb

type DspAccessCnRequest struct {
	Id           string             `json:"id"`
	ApiVersion   string             `json:"api_version"`
	SupportHttps int                `json:"support_https"`
	Imp          []*DspAccessCnImp  `json:"imp"`
	App          *DspAccessCnApp    `json:"app"`
	Device       *DspAccessCnDevice `json:"device"`
	Test         int                `json:"test"`
	TMax         int64              `json:"tmax"`
}

type DspAccessCnImp struct {
	Id          string          `json:"id"`
	AdType      int             `json:"ad_type"`
	W           int             `json:"w"`
	H           int             `json:"h"`
	MinDuration int             `json:"minduration,omitempty"`
	MaxDuration int             `json:"maxduration,omitempty"`
	Tagid       string          `json:"tagid"`
	BidFloor    float64         `json:"bidfloor"`
	Pmp         *DspAccessCnPmp `json:"pmp,omitempty"`
	Ext         *Ext            `json:"ext,omitempty"`
}

type DspAccessCnPmp struct {
	Deals []*DspAccessCnDeal `json:"deals"`
}

type DspAccessCnDeal struct {
	Id       string  `json:"id"`
	BidFloor float64 `json:"bidfloor"`
}

type DspAccessCnApp struct {
	AppId  string `json:"appid"`
	Name   string `json:"name,omitempty"`
	Bundle string `json:"bundle,omitempty"`
	Ver    string `json:"ver,omitempty"`
}

type DspAccessCnDevice struct {
	Ua                 string             `json:"ua,omitempty"`
	Geo                *DspAccessCnGeo    `json:"geo,omitempty"`
	Ip                 string             `json:"ip,omitempty"`
	Ipv6               string             `json:"ipv6,omitempty"`
	DeviceType         int                `json:"devicetype,omitempty"`
	Make               string             `json:"make,omitempty"`
	Model              string             `json:"model,omitempty"`
	Os                 string             `json:"os,omitempty"`
	Osv                string             `json:"osv,omitempty"`
	H                  int                `json:"h,omitempty"`
	W                  int                `json:"w,omitempty"`
	Ppi                int                `json:"ppi,omitempty"`
	Language           string             `json:"language,omitempty"`
	MccMnc             string             `json:"mccmnc,omitempty"`
	ConnectionType     int                `json:"connectiontype,omitempty"`
	Orientation        int                `json:"orientation,omitempty"`
	TimezoneNum        string             `json:"timezone_num,omitempty"`
	Dpi                int                `json:"dpi,omitempty"`
	Density            float64            `json:"density,omitempty"`
	DeviceStartSec     string             `json:"device_start_sec,omitempty"`
	DeviceNameMd5      string             `json:"device_name_md5,omitempty"`
	HardwareMachine    string             `json:"hardware_machine,omitempty"`
	PhysicalMemoryByte string             `json:"physical_memory_byte,omitempty"`
	HarddiskSizeByte   string             `json:"harddisk_size_byte,omitempty"`
	Mac                string             `json:"mac,omitempty"`
	MacMd5             string             `json:"mac_md5,omitempty"`
	Idfa               string             `json:"idfa,omitempty"`
	IdfaMd5            string             `json:"idfa_md5,omitempty"`
	Idfv               string             `json:"idfv,omitempty"`
	HardwareModel      string             `json:"hardware_model,omitempty"`
	Gaid               string             `json:"gaid,omitempty"`
	Oaid               string             `json:"oaid,omitempty"`
	OaidMd5            string             `json:"oaid_md5,omitempty"`
	AndroidId          string             `json:"android_id,omitempty"`
	AndroidIdMd5       string             `json:"android_id_md5,omitempty"`
	Imei               string             `json:"imei,omitempty"`
	ImeiMd5            string             `json:"imei_md5,omitempty"`
	ApiLevel           int                `json:"api_level,omitempty"`
	AppstoreVer        string             `json:"appstore_ver,omitempty"`
	Hms                *DspAccessCnHms    `json:"hms,omitempty"`
	System             *DspAccessCnSystem `json:"system,omitempty"`
	Wx                 *DspAccessCnWx     `json:"wx,omitempty"`
}

type DspAccessCnHms struct {
	HmsCore string `json:"hms_core,omitempty"`
}

type DspAccessCnSystem struct {
	SystemBootMark   string `json:"system_boot_mark,omitempty"`
	SystemUpdateMark string `json:"system_update_mark,omitempty"`
	SystemUpdateSec  string `json:"system_update_sec,omitempty"`
	SystemBirthTime  string `json:"system_birth_time,omitempty"`
	SystemMntId      string `json:"system_mnt_id,omitempty"`
}

type DspAccessCnWx struct {
	WxApiVer     int    `json:"wx_api_ver,omitempty"`
	WxInstalled  bool   `json:"wx_installed,omitempty"`
	WxOpensdkVer string `json:"wx_opensdk_ver,omitempty"`
}

type DspAccessCnGeo struct {
	Lat float64 `json:"lat,omitempty"`
	Lon float64 `json:"lon,omitempty"`
}