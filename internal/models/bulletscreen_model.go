package models

import "time"

type BulletScreen struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Ip        string    `json:"ip"`
	Message   string    `json:"message"`
	UpdatedAt time.Time `json:"-"`
}

func (m BulletScreen) TableName() string {
	return "my_bulletscreen"
}

type BulletScreenModel struct {

}

func NewBulletScreenModel() *BulletScreenModel {
	return &BulletScreenModel{}
}

func (m *BulletScreenModel) GetBulletScreen(offset, limit int) (message []*BulletScreen) {
	message = []*BulletScreen{}
	db.Offset(offset).Limit(limit).Find(&message)
	return
}

func (m *BulletScreenModel) Save(bulletScreen *BulletScreen) error {
	return db.Create(bulletScreen).Error
}
