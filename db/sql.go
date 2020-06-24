package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

var db *gorm.DB

func Setup() {
	var (
		err error
	)
	db, err = gorm.Open("sqlite3", "./salary.db")
	if err != nil {
		panic(err.Error())
	}
	db.DB().SetConnMaxLifetime(80 * time.Second) // 设置链接重置时间
	db.LogMode(true)

	initTable()
}

func GetDB() *gorm.DB {
	return db
}

/**
 * 初始化表
 */
func initTable() {
	GetDB().AutoMigrate(Git{}, Salary{})
	//db.GetDB().Model(&entity.AuthLog{}).DropColumn("ip")
	//db.GetDB().Model(&entity.Rel_Server_IP{}).DropColumn("rel")

}
