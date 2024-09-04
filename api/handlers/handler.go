package handlers

import (
	"encoding/json"
	"go_tg_api_gateway/api/http"
	"go_tg_api_gateway/api/models"
	"go_tg_api_gateway/config"
	"go_tg_api_gateway/pkg/logger"
	"go_tg_api_gateway/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/protobuf/runtime/protoiface"
)

type Handler struct {
	cfg      config.Config
	log      logger.LoggerI
	services services.ServiceManagerI
}

func NewHandler(cfg config.Config, log logger.LoggerI, svcs services.ServiceManagerI) Handler {
	return Handler{
		cfg:      cfg,
		log:      log,
		services: svcs,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	h.handleSuccessResponse(c, status.Code, status.Description, data)
}

func ProtoToStruct(s interface{}, p protoiface.MessageV1) error {
	var jm jsonpb.Marshaler

	jm.EmitDefaults = true
	jm.OrigName = true

	ms, err := jm.MarshalToString(p)

	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(ms), &s)

	return err
}

func (h *Handler) getOffsetParam(c *gin.Context) (offset int, err error) {
	if h.cfg.DefaultOffset != "" {
		offsetStr := c.DefaultQuery("offset", h.cfg.DefaultOffset)
		return strconv.Atoi(offsetStr)
	}

	offsetStr := c.DefaultQuery("offset", "0")
	return strconv.Atoi(offsetStr)
}

func (h *Handler) getLimitParam(c *gin.Context) (offset int, err error) {
	if h.cfg.DefaultLimit != "" {
		limitStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
		return strconv.Atoi(limitStr)
	}
	limitStr := c.DefaultQuery("limit", "10")
	return strconv.Atoi(limitStr)
}

// func (h *Handler) handleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
// 	h.log.Error(message, logger.Int("code", code), logger.Any("error", err))
// 	c.JSON(code, models.ResponseModel{
// 		Code:    code,
// 		Message: message,
// 		Error:   err,
// 	})
// }

func (h *Handler) handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, models.ResponseModel{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
