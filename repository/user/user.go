package user

import (
	"errors"
	"eventku/entities"
	"eventku/middlewares"

	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) Register(user entities.User) (entities.User, error) {

	user.Password, _ = middlewares.HashPassword(user.Password)
	uid := shortuuid.New()
	user.UserUid = uid

	if err := ur.database.Create(&user).Error; err != nil {
		return user, errors.New("invalid input or this email was created (duplicated entry)")
	}

	return user, nil
}

func (ur *UserRepository) GetByUid(userUid string) (entities.User, error) {
	arrUser := entities.User{}

	result := ur.database.Preload("Task").Where("user_uid =?", userUid).First(&arrUser)
	if err := result.Error; err != nil {
		return arrUser, err
	}
	if result.RowsAffected == 0 {
		return arrUser, errors.New("record not found")
	}

	return arrUser, nil
}

func (ur *UserRepository) Update(userUid string, newUser entities.User) (entities.User, error) {

	var user entities.User
	result := ur.database.Where("user_uid =?", userUid).First(&user)

	if result.Error != nil {
		return entities.User{}, errors.New("failed to update user")
	}
	if result.RowsAffected == 0 {
		return entities.User{}, errors.New("user not found")
	}

	if err := ur.database.Model(&user).Where("user_uid =?", userUid).Updates(&newUser).Error; err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) Delete(userUid string) error {

	result := ur.database.Where("user_uid =?", userUid).Delete(&entities.User{})
	if result.Error != nil {
		return result.Error
	}

	return nil

}
