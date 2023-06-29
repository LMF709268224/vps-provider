package api

import (
	"database/sql"
	"net/http"

	"vps-provider/config"
	"vps-provider/storage"
	"vps-provider/utils"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUserInfoHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uuid := claims[identityKey].(string)
	user, err := storage.GetUserByUserUUID(c.Request.Context(), uuid)
	if err != nil {
		c.JSON(http.StatusOK, respError(utils.ErrUserNotFound))
		return
	}
	c.JSON(http.StatusOK, respJSON(user))
}

func UserRegister(c *gin.Context) {
	userInfo := &utils.User{}
	userInfo.Username = c.Query("username")
	userInfo.VerifyCode = c.Query("verify_code")
	userInfo.UserEmail = userInfo.Username
	PassStr := c.Query("password")
	_, err := storage.GetUserByUsername(c.Request.Context(), userInfo.Username)
	if err == nil {
		c.JSON(http.StatusOK, respError(utils.ErrNameExists))
		return
	}
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusOK, respError(utils.ErrInvalidParams))
		return
	}
	//if user.Username != "" {
	//	c.JSON(http.StatusOK, respError(errors.ErrNameExists))
	//	return
	//}
	PassHash, err := bcrypt.GenerateFromPassword([]byte(PassStr), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, respError(utils.ErrPassWord))
		return
	}
	userInfo.PassHash = string(PassHash)
	err = storage.CreateUser(c.Request.Context(), userInfo)
	if err != nil {
		log.Errorf("create user : %v", err)
		c.JSON(http.StatusOK, respError(utils.ErrInternalServer))
		return
	}
	c.JSON(http.StatusOK, respJSON(JsonObject{
		"msg": "success",
	}))
}

func PasswordRest(c *gin.Context) {
	userInfo := &utils.User{}
	userInfo.Username = c.Query("username")
	userInfo.VerifyCode = c.Query("verify_code")
	userInfo.UserEmail = userInfo.Username
	PassStr := c.Query("password")
	_, err := storage.GetUserByUsername(c.Request.Context(), userInfo.Username)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusOK, respError(utils.ErrNameNotExists))
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, respError(utils.ErrInvalidParams))
		return
	}
	//if user.Username != "" {
	//	c.JSON(http.StatusOK, respError(errors.ErrNameExists))
	//	return
	//}
	PassHash, err := bcrypt.GenerateFromPassword([]byte(PassStr), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, respError(utils.ErrPassWord))
		return
	}
	userInfo.PassHash = string(PassHash)

	err = storage.ResetPassword(c.Request.Context(), userInfo.PassHash, userInfo.Username)
	if err != nil {
		log.Errorf("update user : %v", err)
		c.JSON(http.StatusOK, respError(utils.ErrInternalServer))
		return
	}
	c.JSON(http.StatusOK, respJSON(JsonObject{
		"msg": "success",
	}))
}

func sendEmail(sendTo string, vc string) error {
	var EData utils.EmailData
	EData.Subject = "[Application]: Your verify code Info"
	EData.Tittle = "please check your verify code "
	EData.SendTo = sendTo
	EData.Content = "<h1>Your verify code ：</h1>\n"

	EData.Content = vc + "<br>"

	err := utils.SendEmail(config.Cfg.Email, EData)
	if err != nil {
		return err
	}
	return nil
}
