package oez

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Url      string `json:"url" form:"url" gorm:"type:varchar(512)"`
	Times    uint
	ClientIP string `gorm:"type:varchar(64)"`
}

func (url *Url) Create() (uint, error) {
	result := DB.Create(url)
	if result.Error != nil {
		return 0, result.Error
	}
	return url.ID, nil
}

func Get(str string) (*Url, error) {
	id, err := Decode62To10(str)
	if err != nil {
		return nil, err
	}
	var url Url
	result := DB.First(&url, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &url, nil
}

func (url *Url) Look() {
	url.Times += 1
	DB.Save(url)
}
