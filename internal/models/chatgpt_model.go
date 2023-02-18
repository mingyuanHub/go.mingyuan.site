package models

import "time"

type ChatGpt struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Ip        string    `json:"ip"`
	Type      int       `json:"type"`
	Text      string    `json:"text"`
	UpdatedAt time.Time `json:"-"`
}

func (m ChatGpt) TableName() string {
	return "my_chatgpt"
}

const (
	TypeText = 0
	TypeImg  = 1
)

type ChatGptModel struct {

}

func NewChatGptModel() *ChatGptModel {
	return &ChatGptModel{}
}

func (m *ChatGptModel) Save(chatGpt *ChatGpt) error {
	return db.Create(chatGpt).Error
}
