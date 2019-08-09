package config

import (
	"fmt"
	"ginApi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

/*
 数据库及连接配置
  return出db
*/

func Config() *gorm.DB {
	var err error
	db, err = gorm.Open("mysql", "root:709463253@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	} else {
		fmt.Println("数据库已连接！")
		// 关联数据表
		db.AutoMigrate(&models.Music{})

		// 检查模型`Movie`的表是否存在
		hasTable := db.HasTable(&models.Music{})
		fmt.Println(hasTable, "--")
		if !hasTable {
			// 为模型`Product`创建表,CHARSET=utf8设置数据库的字符集为utf8
			db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").CreateTable(&models.Music{})
		}
	}
	return db
}
