package user

import (
	"database/sql"
	"errors"
	"eventku/entities"
	"eventku/middlewares"
	"fmt"
	"strings"
	"time"

	"github.com/lithammer/shortuuid"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Register(user entities.User) (entities.User, error) {

	user.Password, _ = middlewares.HashPassword(user.Password)
	uid := shortuuid.New()
	user.UserUid = uid
	_, err := ur.db.Exec("insert into user (uid,name,email,password,address,gender,created_at,updated_at,deleted_at) values (?, ?, ?, ?,?,?,?,?,?)", user.UserUid, user.Name, user.Email, user.Password, user.Address, user.Gender, user.CreatedAt.(int64), user.UpdatedAt.(int64), user.DeletedAt.(int64))
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) GetByUid(userUid string) (entities.User, error) {
	user := entities.User{}

	err := ur.db.QueryRow("SELECT * FROM user WHERE uid=?", userUid).Scan(&user.UserUid, &user.Name, &user.Email, &user.Password, &user.Address, &user.Gender, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return user, err
	}
	if user.UserUid == "" {
		return user, errors.New("record not found")
	}

	return user, nil
}

func (ur *UserRepository) Update(userUid string, newUser entities.User) (entities.User, error) {

	var user entities.User

	queryPrepareField := make([]string, 0)
	queryExec := make([]interface{}, 0)
	if newUser.Name != "" {
		queryPrepareField = append(queryPrepareField, "name")
		queryExec = append(queryExec, newUser.Name)
	}
	if newUser.Email != "" {
		queryPrepareField = append(queryPrepareField, "email")
		queryExec = append(queryExec, newUser.Email)
	}
	if newUser.Address != "" {
		queryPrepareField = append(queryPrepareField, "address")
		queryExec = append(queryExec, newUser.Address)
	}
	if newUser.Gender != "" {
		queryPrepareField = append(queryPrepareField, "gender")
		queryExec = append(queryExec, newUser.Gender)
	}
	timeLocation, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().In(timeLocation).Unix()
	queryPrepareField = append(queryPrepareField, "updated_at")
	queryExec = append(queryExec, timeNow)
	queryExec = append(queryExec, userUid)

	queryPrepareString := strings.Join(queryPrepareField, "= ? ,")
	queryPrepareString += "= ? "

	stmt, err := ur.db.Prepare("UPDATE user SET " + queryPrepareString + " where uid=?")

	if err != nil {
		return user, err
	}
	// execute
	if _, err := stmt.Exec(queryExec...); err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Delete(userUid string) error {

	_, err := ur.db.Exec("delete from user where uid = ?", userUid)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil

}

func (ur *UserRepository) Search(q string) ([]entities.User, error) {
	arrUser := []entities.User{}

	// if len(q) < 3 {
	// 	if len(q) == 1 {
	// 		ur.db.Debug().Where("a =?", q).Find(&arrUser)
	// 		return arrUser, nil
	// 	}
	// 	if len(q) == 2 {
	// 		ur.db.Debug().Where("b =?", q).Find(&arrUser)
	// 		return arrUser, nil

	// 	}
	// 	if len(q) == 3 {
	// 		ur.db.Debug().Where("c =?", q).Find(&arrUser)
	// 		return arrUser, nil

	// 	}
	// }
	// sql := "%" + q + "%"

	// result := ur.db.Debug().Where("name like ?", sql).Find(&arrUser)
	// if err := result.Error; err != nil {
	// 	return arrUser, err
	// }
	// if result.RowsAffected == 0 {
	// 	return arrUser, errors.New("record not found")
	// }

	return arrUser, nil
}

func (ur *UserRepository) Dummy(length int) bool {

	// names := []string{"roger", "joni", "mail", "bruto", "icon", "abeng", "jangkrik", "zeagger", "connie", "terlalu"}

	// for i := 0; i < length; i++ {
	// 	uid := shortuuid.New()

	// 	user := entities.User{
	// 		UserUid: uid,
	// 		Name:    names[rand.Intn(9)],
	// 		Email:   faker.Email(),
	// 		Address: "jl.dramaga no.22",
	// 		Gender:  "male",
	// 	}
	// 	user.Password = "xyz"
	// 	user.Password, _ = middlewares.HashPassword(user.Password)

	// 	if err := ur.db.Create(&user).Error; err != nil {
	// 		return false
	// 	}
	// }

	return true
}
