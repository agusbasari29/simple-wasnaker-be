package router

import (
	"simpel-wasnaker/ipak3/database"
	"simpel-wasnaker/ipak3/handler"
	"simpel-wasnaker/ipak3/helper"
	"simpel-wasnaker/ipak3/repository"
	"simpel-wasnaker/ipak3/service"

	"github.com/gin-gonic/gin"
)

type AuthRoute struct{}

func (r AuthRoute) Route() []helper.Route {
	db := database.SetupDBConnection()
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	AuthHandler := handler.NewAuthHandler(authService)

	return []helper.Route{
		{
			Method:      "POST",
			Path:     "/login",
			HandlerFunc: []gin.HandlerFunc{AuthHandler.Login},
		},
	}
}
