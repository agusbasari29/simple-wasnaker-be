package helper

import "github.com/gin-gonic/gin"

type Route struct {
	Method      string
	Path        string
	HandlerFunc []gin.HandlerFunc
}

type Handler interface {
	Route() []Route
}
