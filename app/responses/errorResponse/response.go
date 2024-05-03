package errorResponse

import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {

}

func ResponseWithObject(c *gin.Context, statusCode int) {

}
