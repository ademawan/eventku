package auth

import (
	"eventku/delivery/controllers/common"
	"eventku/entities"
	"eventku/middlewares"
	"eventku/repository/auth"
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo auth.Auth
}

func New(repo auth.Auth) *AuthController {
	return &AuthController{
		repo: repo,
	}
}

func (ac *AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		Userlogin := LoginReqFormat{}

		c.Bind(&Userlogin)
		err_validate := c.Validate(&Userlogin)

		if err_validate != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		checkedUser, err_repo := ac.repo.Login(Userlogin.Email, Userlogin.Password)

		if err_repo != nil {
			var statusCode int = 500
			if err_repo.Error() == "email not found" {
				statusCode = http.StatusUnauthorized
			} else if err_repo.Error() == "incorrect password" {
				statusCode = http.StatusUnauthorized
			}
			return c.JSON(statusCode, common.InternalServerError(statusCode, err_repo.Error(), nil))
		}
		token, err := middlewares.GenerateToken(checkedUser)
		response := UserLoginResponse{
			User_uid: checkedUser.UserUid,
			Name:     checkedUser.Name,
			Email:    checkedUser.Email,
			Gender:   checkedUser.Gender,
			Token:    token,
		}

		if err != nil {
			return c.JSON(http.StatusNotAcceptable, common.BadRequest(http.StatusNotAcceptable, "error in process token", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Login successfully", response))

	}
}
func (ac *AuthController) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		userUid := middlewares.ExtractTokenUserUid(c)
		log.Info(userUid)
		token, _ := middlewares.GenerateToken(entities.User{UserUid: "xxx"})
		log.Info(token)

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Logout successfully", nil))

	}
}

// func (ac *AuthController) LoginGoogle() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		googleConfig := utils.SetUpConfig()
// 		url := googleConfig.AuthCodeURL("randomstate")
// 		fmt.Println(url, "F GoogleLogin")

// 		return c.JSON(http.StatusSeeOther, common.Success(http.StatusOK, "successfully", url))

// 	}
// }
// func (ac *AuthController) LoginGoogleCallback() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		state := c.QueryParam("state")

// 		if state != "randomstate" {
// 			fmt.Println("states dont match")
// 			return c.JSON(http.StatusBadRequest, common.InternalServerError(http.StatusBadRequest, state, nil))

// 		}
// 		code := c.QueryParam("code")
// 		googleConfig := utils.SetUpConfig()
// 		tokenGoogle, err := googleConfig.Exchange(context.Background(), code)

// 		if err != nil {
// 			fmt.Print("code -Token Exchange Failed")
// 		}
// 		resp, err2 := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tokenGoogle.AccessToken)
// 		userData, errresp := ioutil.ReadAll(resp.Body)

// 		if errresp != nil {
// 			fmt.Println("JSON parsing failed", err2)
// 		}

// 		var resGoogle UserLoginGoogle
// 		json.Unmarshal(userData, &resGoogle)

// 		checkedUser, err_repo := ac.repo.LoginGoogle(resGoogle.Email)

// 		if err_repo != nil {
// 			var statusCode int = 500
// 			if err_repo.Error() == "email not found" {
// 				statusCode = http.StatusUnauthorized
// 			} else if err_repo.Error() == "incorrect password" {
// 				statusCode = http.StatusUnauthorized
// 			}
// 			return c.JSON(statusCode, common.InternalServerError(statusCode, err_repo.Error(), nil))
// 		}

// 		token, _ := middlewares.GenerateToken(checkedUser)

// 		response := UserLoginResponse{
// 			User_uid: checkedUser.UserUid,
// 			Name:     checkedUser.Name,
// 			Email:    checkedUser.Email,
// 			Gender:   checkedUser.Gender,
// 			Token:    token,
// 		}

// 		return c.JSON(http.StatusSeeOther, common.Success(http.StatusOK, "successfully", response))

// 	}
// }
func (ac *AuthController) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		type dataMap map[string]interface{}
		data := dataMap{"message": "Hello World!"}
		// return c.Render(http.StatusOK, "index.html", data)
		return c.Render(http.StatusOK, "pengembangan.html", data)

	}
}
func (ac *AuthController) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		type dataMap map[string]interface{}
		data := dataMap{"message": "Hello World!"}
		return c.Render(http.StatusOK, "khitanan.html", data)

	}
}
func (ac *AuthController) Dashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		type dataMap map[string]interface{}
		data := dataMap{"message": "Hello World!"}
		return c.Render(http.StatusOK, "blank.html", data)

	}
}
