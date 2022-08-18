package assassination

import "gorm.io/gorm"

type Assassination struct {
	gorm.Model
	Id       uint `gorm:"primaryKey;autoIncrement"`
	ChatId   int
	UserId   int
	UserName string
	Temple   string
}
