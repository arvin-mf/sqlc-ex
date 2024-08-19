package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, httpCode int, msg string, data interface{}) {
	c.JSON(httpCode, map[string]interface{}{
		"status":  "success",
		"message": msg,
		"data":    data,
	})
}

func FailOrError(c *gin.Context, httpCode int, msg string, err error) {
	c.JSON(httpCode, gin.H{
		"status":  "fail",
		"message": msg,
		"data": gin.H{
			"error": err.Error(),
		},
	})
}

func ErrorEmptyField(c *gin.Context) {
	c.JSON(400, gin.H{
		"status":  "fail",
		"message": "Please fill the empty field",
	})
}
