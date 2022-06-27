package models

import "time"

type PixelWars struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Color     string    `json:"color"`
	UpdatedAt time.Time `json:"-"`
}

func (m PixelWars) TableName() string {
	return "my_pixelwars"
}

type PixelWarsModel struct {

}

func NewPixelWarsModel() *PixelWarsModel {
	return &PixelWarsModel{}
}

func (m *PixelWarsModel) GetColor() (pixel []*PixelWars) {
	pixel = []*PixelWars{}
	db.Find(&pixel)
	return
}

func (m *PixelWarsModel) Save(pixel *PixelWars) error {
	return db.Save(pixel).Error
}
