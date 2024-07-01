package middleware

import (
	"encoding/json"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Response setting gin.JSON
func Response(c *gin.Context, req interface{}, res models.Response) {
	// LOGGER
	reqByte, _ := json.Marshal(req)
	resByte, _ := json.Marshal(res)
	logger.Infof("[mvp-shop-backend:log] [RequestURL] : %s, [RequestMethod] : %s, [RequestBody] : %s, [ResponseData] : %s", c.Request.RequestURI, c.Request.Method, string(reqByte), string(resByte))

	c.JSON(res.Code, res)
}
