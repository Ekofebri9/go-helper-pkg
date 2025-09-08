package http

import (
	"go/helperpkg/response"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

type Context interface {
	gin.Context | fiber.Ctx
}

type Handler[T Context] interface {
	JSON(c *T, r response.Response)
	AbortJSON(c *T, r response.Response)
	ErrorResponse(err error) *response.ErrorResponse
	BindJSON(c *T, obj any) response.Response
}
