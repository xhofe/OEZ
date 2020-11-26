package oez

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Url   		string 	`json:"url" form:"url" gorm:"type:varchar(512)"`
	Times 		uint
	ClientIP	string	`gorm:"type:varchar(64)"`
}

func (url *Url)Create() (uint,error) {
	result:=DB.Create(url)
	if result.Error!=nil {
		return 0,result.Error
	}
	return url.ID,nil
}

func Get(str string) *Url {
	id:=Decode62To10(str)
	var url Url
	result:=DB.First(&url,id)
	if result.Error!=nil {
		return nil
	}
	url.Times=url.Times+1
	go func() {
		DB.Save(url)
	}()
	return &url
}