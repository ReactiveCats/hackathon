package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// not sure if it's good idea
//func GinResponse(ctx *gin.Context, data interface{}, payload ...interface{}) {
//	switch d := data.(type) {
//	case int:
//		GinOkResponse(ctx, d, payload...)
//	case error:
//		GinErrResponse(ctx, d)
//	default:
//		logrus.Panicf("unknown type to response: %T", d)
//	}
//}

func GinOkResponse(ctx *gin.Context, httpStatusCode int, payload ...interface{}) {
	switch len(payload) {
	case 0:
		ctx.AbortWithStatus(httpStatusCode)
	case 1:
		ctx.AbortWithStatusJSON(httpStatusCode, payload[0])
	default:
		logrus.Panicf("too many argument: %d", len(payload))
	}
}

func GinErrResponse(ctx *gin.Context, err error) {
	switch e := err.(type) {
	case ServerError:
		_ = ctx.Error(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	case ClientError:
		ctx.String(e.HttpStatusCode(), e.Message())
	default:
		logrus.Panicf("unknown error (%T) to response: %s", err, err.Error())
	}
	ctx.Abort()
}