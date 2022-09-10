package utils

import (
	"database/sql"
	"eventku/configs"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {
	// func InitDB(config *configs.AppConfig) (*sql.DB, error) {

	// connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
	// 	config.Database.Username,
	// 	config.Database.Password,
	// 	config.Database.Address,
	// 	config.Database.Port,
	// 	config.Database.Name,
	// )
	// db, err := sql.Open("mysql", connectionString)
	// if err != nil {
	// 	return nil, err
	// }

	DBURL := fmt.Sprintf("postgres://%s:%v@%s:%s/%s", config.Database.Username, config.Database.Password, config.Database.Address, config.Database.Port, config.Database.Name)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	// InitMigrate(db)
	return db
}

func InitMigrate(db *sql.DB) {
	// db.Migrator().DropTable(&entities.User{})
	// db.AutoMigrate(&entities.User{})
	rows, err := db.Query("CREATE TABLE IF NOT EXISTS `user` (`uid` varchar(32) NOT NULL,`name` varchar(50) NOT NULL,`email` varchar(50) NOT NULL,`password` varchar(100) NOT NULL,`address` varchar(50),`gender` varchar(10) NOT NULL,created_at INT,updated_at INT,deleted_at INT,UNIQUE KEY unique_uid (uid),UNIQUE KEY unique_email (email)) ENGINE=InnoDB DEFAULT CHARSET=latin1;")
	if err != nil {
		return
	}
	fmt.Println("sukses")
	rows.Close()
	// // var userUid []string

	// // layoutFormat := "2006-01-02T15:04"
	// // todoDateTime, _ := time.Parse(layoutFormat, "2022-03-31T12:26")

	// userUidtest := shortuuid.New()
	// password, _ := middlewares.HashPassword("xyz")

	// db.Create(&entities.User{
	// 	UserUid:  userUidtest,
	// 	Name:     "Ade Mawan",
	// 	Email:    "ademawan1210@gmail.com",
	// 	Password: password,
	// 	Address:  "jl.dramaga no.22",
	// 	Gender:   "male",
	// })
	// for i := 0; i < 1; i++ {
	// 	taskUid := shortuuid.New()

	// 	db.Create(&entities.Task{
	// 		TaskUid:        taskUid,
	// 		Title:          faker.TitleMale(),
	// 		Priority:       "hight",
	// 		UserUid:        userUidtest,
	// 		Status:         "waithing",
	// 		Note:           "catatan catatan catatan",
	// 		Todo_date_time: todoDateTime,
	// 	})

	// }

	// //testing insert many data to realation table

	// // layoutFormat3 := "2006-01-02T15:04"
	// // todoDateTime3, _ := time.Parse(layoutFormat3, "2022-03-31T12:26")

	// // userUidtest3 := shortuuid.New()
	// // password, _ := middlewares.HashPassword("xyz")

	// // tasks := []entities.Task{}
	// // for i := 0; i < 10; i++ {

	// // 	taskUidtest3 := shortuuid.New()
	// // 	tasktest := entities.Task{
	// // 		TaskUid:        taskUidtest3,
	// // 		Title:          faker.TitleMale(),
	// // 		Priority:       "hight",
	// // 		Status:         "waithing",
	// // 		Note:           "catatan catatan catatan",
	// // 		Todo_date_time: todoDateTime3,
	// // 	}

	// // 	tasks = append(tasks, tasktest)
	// // }

	// // db.Create(&entities.User{
	// // 	UserUid:  userUidtest3,
	// // 	Name:     "anonimustiga",
	// // 	Email:    "anonimus3@gmail.com",
	// // 	Password: password,
	// // 	Address:  "jl.dramaga no.22",
	// // 	Gender:   "male",
	// // 	Task:     tasks,
	// // })

	// for i := 0; i < 1; i++ {

	// 	userUid := shortuuid.New()
	// 	password, _ := middlewares.HashPassword("xyz")

	// 	db.Create(&entities.User{
	// 		UserUid:  userUid,
	// 		Name:     faker.Name(),
	// 		Email:    faker.Email(),
	// 		Password: password,
	// 		Address:  "jl.dramaga no.22",
	// 		Gender:   "female",
	// 	})
	// 	taskUid := shortuuid.New()

	// 	db.Create(&entities.Task{
	// 		TaskUid:        taskUid,
	// 		Title:          faker.TitleMale(),
	// 		Priority:       "hight",
	// 		UserUid:        userUid,
	// 		Status:         "waithing",
	// 		Note:           "catatan catatan catatan",
	// 		Todo_date_time: todoDateTime,
	// 	})

	// }

	// for i := 0; i < 1; i++ {

	// 	userUid := shortuuid.New()
	// 	password, _ := middlewares.HashPassword("xyz")

	// 	db.Create(&entities.User{
	// 		UserUid:  userUid,
	// 		Name:     faker.Name(),
	// 		Email:    faker.Email(),
	// 		Password: password,
	// 		Address:  "jl.dramaga no.22",
	// 		Gender:   "male",
	// 	})
	// 	taskUid := shortuuid.New()

	// 	db.Create(&entities.Task{
	// 		TaskUid:        taskUid,
	// 		Title:          faker.TitleMale(),
	// 		Priority:       "hight",
	// 		UserUid:        userUid,
	// 		Status:         "waithing",
	// 		Note:           "catatan catatan catatan",
	// 		Todo_date_time: todoDateTime,
	// 	})

	// }

}
