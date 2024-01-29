package db

import 
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(localhost:3306)/user?charset=utf8&parseTime=True&loc=Local", // data source name, 详情参考：https://github.com/go-sql-driver/mysql#dsn-data-source-name
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = db.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

}

// func GetDB() *gorm.DB {
// 	return db
// }
