package models

import "time"

type Share struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Uuid      string    `json:"uuid"`
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	Data      string    `json:"data"`
	UpdatedAt time.Time `json:"-"`
}

func (m Share) TableName() string {
	return "my_share"
}

type ShareModel struct {

}

func NewShareModel() *ShareModel {
	return &ShareModel{}
}

func (m *ShareModel) Save(share *Share) error {
	return db.Create(share).Error
}

func (m *ShareModel) GetShareData(path, token string) (data string) {
	row := db.Table("my_share").Where("path = ? AND token = ?", path, token).Select("data").Row()
	row.Scan(&data)
	return
}
