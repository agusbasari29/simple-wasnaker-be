package router

import (
	"simpel-wasnaker/ipak3/database"
	"simpel-wasnaker/ipak3/handler"
	"simpel-wasnaker/ipak3/helper"
	"simpel-wasnaker/ipak3/repository"
	"simpel-wasnaker/ipak3/service"

	"github.com/gin-gonic/gin"
)

type UserRoute struct{}

func (r UserRoute) Route() []helper.Route {
	db := database.SetupDBConnection()
	userRepo := repository.NewUserRepository(db)
	contactRepo := repository.NewContactRepository(db)
	userService := service.NewUserService(userRepo, contactRepo)
	userHandler := handler.NewUserHandler(userService)
	return []helper.Route{
		{
			Method:      "GET",
			Path:        "/user/profile/:id",
			HandlerFunc: []gin.HandlerFunc{userHandler.GetByID},
		},
		{
			Method:      "PUT",
			Path:        "/user/update",
			HandlerFunc: []gin.HandlerFunc{userHandler.Create},
		},
	}
}
