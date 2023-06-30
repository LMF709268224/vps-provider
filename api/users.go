package api

import (
	"database/sql"
	"net/http"

	"vps-provider/storage/mysql"
	"vps-provider/types"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"vps-provider/errors"
)

func getUserInfoHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uuid := claims[identityKey].(string)
	user, err := mysql.GetUserByUserUUID(c.Request.Context(), uuid)
	if err != nil {
		c.JSON(http.StatusOK, respError(errors.ErrUserNotFound))
		return
	}
	c.JSON(http.StatusOK, respJSON(user))
}

func userRegisterHandler(c *gin.Context) {
	userInfo := &types.User{}
	userInfo.UserName = c.Query("username")
	passStr := c.Query("password")
	_, err := mysql.GetUserByUserName(c.Request.Context(), userInfo.UserName)
	if err == nil {
		c.JSON(http.StatusOK, respError(errors.ErrNameExists))
		return
	}
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusOK, respError(errors.ErrInvalidParams))
		return
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(passStr), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, respError(errors.ErrPassWord))
		return
	}
	userInfo.UUID = uuid.NewString()
	userInfo.PassHash = string(passHash)
	err = mysql.SaveUserInfo(c.Request.Context(), userInfo)
	if err != nil {
		log.Errorf("create user : %v", err)
		c.JSON(http.StatusOK, respError(errors.ErrInternalServer))
		return
	}
	c.JSON(http.StatusOK, respJSON(jsonObject{
		"msg": "success",
	}))
}

func resetPasswordHandler(c *gin.Context) {
	userInfo := &types.User{}
	userInfo.UserName = c.Query("username")
	passStr := c.Query("password")
	_, err := mysql.GetUserByUserName(c.Request.Context(), userInfo.UserName)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusOK, respError(errors.ErrNameNotExists))
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, respError(errors.ErrInvalidParams))
		return
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(passStr), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, respError(errors.ErrPassWord))
		return
	}
	userInfo.PassHash = string(passHash)

	err = mysql.ResetPassword(c.Request.Context(), userInfo.PassHash, userInfo.UserName)
	if err != nil {
		log.Errorf("update user : %v", err)
		c.JSON(http.StatusOK, respError(errors.ErrInternalServer))
		return
	}
	c.JSON(http.StatusOK, respJSON(jsonObject{
		"msg": "success",
	}))
}
