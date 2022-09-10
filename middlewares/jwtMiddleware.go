package middlewares

import (
	"errors"
	config "eventku/configs"
	"eventku/entities"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(u entities.User) (string, error) {
	if u.UserUid == "" {
		return "cannot Generate token", errors.New("user_uid == null")
	}

	codes := jwt.MapClaims{
		"user_uid": u.UserUid,
		// "email":    u.Email,
		// "password": u.Password,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"auth": true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, codes)
	// fmt.Println(token)
	return token.SignedString([]byte(config.JWT_SECRET))
}
func ExtractTokenUserUid(e echo.Context) string {
	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
	if user.Valid {
		codes := user.Claims.(jwt.MapClaims)
		id := codes["user_uid"].(string)
		return id
	}
	return ""
}

// func ExtractRoles(e echo.Context) bool {
// 	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
// 	if user.Valid {
// 		codes := user.Claims.(jwt.MapClaims)
// 		id := codes["roles"].(bool)
// 		return id
// 	}
// 	return false
// }
