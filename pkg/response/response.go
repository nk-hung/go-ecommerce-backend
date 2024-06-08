package response

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

// success response
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// error response
func ErrorResponse(c *gin.Context, code int) {
	c.JSON(code, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
