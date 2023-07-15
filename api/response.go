package api

import (
	err "vps-provider/errors"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type jsonObject map[string]interface{}

func respJSON(code int, v interface{}) gin.H {
	return gin.H{
		"code": code,
		"data": v,
	}
}

func respError(e error) gin.H {
	var genericError err.GenericError
	if !errors.As(e, &genericError) {
		genericError = err.ErrUnknown
	}

	return gin.H{
		"code": -1,
		"err":  genericError.Code,
		"msg":  genericError.Error(),
	}
}
