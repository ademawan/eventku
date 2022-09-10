package auth

import (
	"database/sql"
	"eventku/entities"
	"eventku/middlewares"
	"fmt"

	"errors"
)

type AuthDb struct {
	db *sql.DB
}

func New(db *sql.DB) *AuthDb {
	return &AuthDb{
		db: db,
	}
}

func (ad *AuthDb) Login(email, password string) (entities.User, error) {
	user := entities.User{}

	err := ad.db.QueryRow("SELECT * FROM user WHERE email=?", email).Scan(&user.UserUid, &user.Name, &user.Email, &user.Password, &user.Address, &user.Gender, &user.CreatedAt, &user.DeletedAt, &user.DeletedAt)
	if err != nil {
		fmt.Println(err.Error())
		return user, errors.New("email not found")
	}

	// for rows.Next() {
	//     var err = rows.Scan(&user.UserUid, &user.Name, &user.Email,user.Password,user.Address,user.Gender,user.CreatedAt,user.DeletedAt)

	//     if err != nil {
	//         fmt.Println(err.Error())
	//         return user,err
	//     }

	//     result = append(result, each)
	// }

	match := middlewares.CheckPasswordHash(password, user.Password)

	if !match {
		return user, errors.New("incorrect password")
	}

	return user, nil
}

func (ad *AuthDb) LoginGoogle(email string) (entities.User, error) {
	user := entities.User{}

	err := ad.db.QueryRow("SELECT * FROM user WHERE email=?", email).Scan(&user.UserUid, &user.Name, &user.Email, user.Password, user.Address, user.Gender, user.CreatedAt, user.DeletedAt)
	if err != nil {
		return user, errors.New("email not found")
	}

	return user, nil
}
