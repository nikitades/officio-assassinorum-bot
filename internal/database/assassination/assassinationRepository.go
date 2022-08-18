package assassination

import (
	"OfficioAssassinorumBot/internal/conf"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	_db, err := gorm.Open(
		sqlite.Open(conf.CurrentConfig.DatabaseConnection),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err)
	}

	_db.AutoMigrate(&Assassination{})

	db = _db
}

func Add(chatId int, userId int, userName string, temple string) {
	a := &Assassination{
		ChatId:   chatId,
		UserId:   userId,
		UserName: userName,
		Temple:   temple,
	}

	db.Create(a)
}

func FindAllInChat(chatId int) []Assassination {
	var ass []Assassination

	db.Where("chat_id = ?", chatId).Find(&ass)

	return ass
}
