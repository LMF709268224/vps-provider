package api

import (
	"vps-provider/utils"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type JsonObject map[string]interface{}

func respJSON(v interface{}) gin.H {
	return gin.H{
		"code": 0,
		"data": v,
	}
}

func respError(e error) gin.H {
	var genericError utils.GenericError
	if !errors.As(e, &genericError) {
		genericError = utils.ErrUnknown
	}

	return gin.H{
		"code": -1,
		"err":  genericError.Code,
		"msg":  genericError.Error(),
	}
}
