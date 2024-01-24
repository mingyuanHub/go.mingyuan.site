package models

type Dsp struct {
	Id        int    `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	UniqueKey string `json:"unique_key"`
	IsCn      int    `json:"is_cn"`
	Adm       string `json:"adm"`
}

func (m Dsp) TableName() string {
	return "my_dsp"
}

type DspModel struct {

}

func NewDspModel() *DspModel {
	return &DspModel{}
}


func (m *DspModel) GetDspList(isCn int) (dspList []*Dsp) {
	dspList = []*Dsp{}
	db.Where("is_cn = ?", isCn).Order("updatetime desc").Find(&dspList)
	return
}

func (m *DspModel) GetDspAdmById(id int) (adm string) {
	row := db.Table("my_dsp").Where("id = ?", id).Select("adm").Row()
	row.Scan(&adm)
	return
}

func (m *DspModel) GetDspByUniqueKey(uniqueKey string) *Dsp {
	var dsp = &Dsp{}
	db.Table("my_dsp").Where("unique_key = ?", uniqueKey).First(&dsp)
	return dsp
}

func (m *DspModel) Save(dsp *Dsp) error {
	return db.Save(dsp).Error
}
