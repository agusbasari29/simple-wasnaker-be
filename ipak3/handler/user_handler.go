package handler

import (
	"net/http"
	"simpel-wasnaker/ipak3/helper"
	"simpel-wasnaker/ipak3/request"
	"simpel-wasnaker/ipak3/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Create(ctx *gin.Context) {
	var request request.UserRegisterRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	user, err := h.userService.Create(request)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", nil, user)
	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Update(ctx *gin.Context) {
	var request request.UserUpdateRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	user, err := h.userService.Update(request)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", nil, user)
	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Delete(ctx *gin.Context) {
	var request request.UserDeleteRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	err = h.userService.Delete(request)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", nil, nil)
	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) GetByID(ctx *gin.Context) {
	var request request.UserGetRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	user, err := h.userService.Find(request.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", nil, user)
	ctx.JSON(http.StatusOK, response)
}
