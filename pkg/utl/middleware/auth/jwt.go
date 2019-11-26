package jwt

import (
	"github.com/huynhdunguyen/boilerplate-golang-sample/pkg/utl/config"

	"github.com/gin-gonic/gin"
)

// New is func initialize
func New(cfg *config.JWT) {

}

// Authenticate token
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
