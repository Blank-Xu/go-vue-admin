package context

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Context struct {
	gin.Context

	UID    int
	Logger *zerolog.Logger
}
