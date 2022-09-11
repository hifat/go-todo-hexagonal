package configs

import (
	"github.com/hifat/go-todo-hexagonal/internal/repository"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func GormDB() *gorm.DB {
	// dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
	// 	os.Getenv("POSTGRES_HOST"),
	// 	os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PASSWORD"),
	// 	os.Getenv("POSTGRES_DB_NAME"),
	// 	os.Getenv("POSTGRES_PORT"),
	// )

	// dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&Loc=local",
	// 	os.Getenv("MYSQL_HOST"),
	// 	os.Getenv("MYSQL_USER"),
	// 	os.Getenv("MYSQL_PASSWORD"),
	// 	os.Getenv("MYSQL_DB_NAME"),
	// 	os.Getenv("MYSQL_PORT"),
	// )

	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/go_todo"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&repository.Task{})

	return db
}
