package routes

import (
	"eventku/delivery/controllers/auth"

	"eventku/delivery/controllers/user"
	"eventku/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo,
	aa *auth.AuthController,
	uc *user.UserController,

) {

	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	//ROUTE index
	e.GET("/", aa.Index())
	e.GET("/login", aa.LoginPage())
	e.GET("/register", aa.RegisterPage())

	e.GET("/undangan/khitanan", aa.Post())
	e.GET("/dashboard", aa.Dashboard())

	//ROUTE REGISTER - LOGIN USERS
	e.POST("users/register", uc.Register())
	e.POST("users/login", aa.Login())
	// e.GET("google/login", aa.LoginGoogle())
	// e.GET("google/callback", aa.LoginGoogleCallback())
	e.POST("users/logout", aa.Logout(), middlewares.JwtMiddleware())

	//ROUTE USERS
	e.GET("/users/me", uc.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me", uc.Delete(), middlewares.JwtMiddleware())
	// e.GET("/users/me/search", uc.Search())
	// e.GET("/users/me/dummy", uc.Dummy())

}
