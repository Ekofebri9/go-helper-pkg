package gin

import (
	"go/helperpkg/errors"
	"go/helperpkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	validator *validator.Validate
}

func (h *Handler) BindJSON(c *gin.Context, obj any) response.Response {
	if err := c.ShouldBindJSON(obj); err != nil {
		return h.ErrorResponse(errors.NewInvalidInputError("invalid paylaod", err))
	}
	return nil
}

func (h *Handler) ValidatePayload() *response.ErrorResponse {
	if h.validator == nil {
		h.validator = validator.New()
	}
	return nil
}

func (h *Handler) JSON(c *gin.Context, r response.Response) {
	c.JSON(r.GetStatusCode(), r)
}

func (h *Handler) AbortJSON(c *gin.Context, r response.Response) {
	c.AbortWithStatusJSON(r.GetStatusCode(), r)
}

func (h *Handler) ErrorResponse(err error) *response.ErrorResponse {
	if errr, ok := err.(*errors.Error); ok {
		return &response.ErrorResponse{
			Error: errr.Error(),
			Meta: response.Meta{
				ResponseCode:    errr.GetHTTPCode(),
				ResponseMessage: errr.GetMessage(),
			},
		}
	} else {
		return &response.ErrorResponse{
			Error: err.Error(),
			Meta: response.Meta{
				ResponseCode:    http.StatusInternalServerError,
				ResponseMessage: http.StatusText(http.StatusInternalServerError),
			},
		}
	}
}

func NewHandler(validator *validator.Validate) *Handler {
	return &Handler{
		validator: validator,
	}
}
