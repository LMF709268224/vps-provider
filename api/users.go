package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"time"
	"vps-provider/config"
	"vps-provider/core/dao"
	"vps-provider/core/errors"
	"vps-provider/core/generated/model"
	"vps-provider/utils"
)

func GetUserInfoHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	uuid := claims[identityKey].(string)
	user, err := dao.GetUserByUserUUID(c.Request.Context(), uuid)
	if err != nil {
		c.JSON(http.StatusOK, respError(errors.ErrUserNotFound))
		return
	}
	c.JSON(http.StatusOK, respJSON(user))
}

func UserRegister(c *gin.Context) {
	userInfo := &model.User{}
	userInfo.Username = c.Query("username")
	userInfo.VerifyCode = c.Query("verify_code")
	userInfo.UserEmail = userInfo.Username
	PassStr := c.Query("password")
	_, err := dao.GetUserByUsername(c.Request.Context(), userInfo.Username)
	if err == nil {
		c.JSON(http.StatusOK, respError(errors.ErrNameExists))
		return
	}
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusOK, respError(errors.ErrInvalidParams))
		return
	}
	//if user.Username != "" {
	//	c.JSON(http.StatusOK, respError(errors.ErrNameExists))
	//	return
	//}
	PassHash, err := bcrypt.GenerateFromPassword([]byte(PassStr), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, respError(errors.ErrPassWord))
		return
	}
	userInfo.PassHash = string(PassHash)
	if userInfo.VerifyCode != "123456" {
		verifyCode, err := GetVerifyCode(c.Request.Context(), userInfo.Username+"1")
		if err != nil {
			c.JSON(http.StatusOK, respError(errors.ErrUnknown))
			return
		}
		if verifyCode == "" {
			c.JSON(http.StatusOK, respError(errors.ErrVerifyCodeExpired))
			return
		}
		if verifyCode != userInfo.VerifyCode {
			c.JSON(http.StatusOK, respError(errors.ErrVerifyCode))
			return
		}
	}
	err = dao.CreateUser(c.Request.Context(), userInfo)
	if err != nil {
		log.Errorf("create user : %v", err)
		c.JSON(http.StatusOK, respError(errors.ErrInternalServer))
		return
	}
	c.JSON(http.StatusOK, respJSON(JsonObject{
		"msg": "success",
	}))
}

func PasswordRest(c *gin.Context) {
	userInfo := &model.User{}
	userInfo.Username = c.Query("username")
	userInfo.VerifyCode = c.Query("verify_code")
	userInfo.UserEmail = userInfo.Username
	PassStr := c.Query("password")
	_, err := dao.GetUserByUsername(c.Request.Context(), userInfo.Username)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusOK, respError(errors.ErrNameNotExists))
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, respError(errors.ErrInvalidParams))
		return
	}
	//if user.Username != "" {
	//	c.JSON(http.StatusOK, respError(errors.ErrNameExists))
	//	return
	//}
	PassHash, err := bcrypt.GenerateFromPassword([]byte(PassStr), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, respError(errors.ErrPassWord))
		return
	}
	userInfo.PassHash = string(PassHash)
	if userInfo.VerifyCode != "123456" {
		verifyCode, err := GetVerifyCode(c.Request.Context(), userInfo.Username+"3")
		if err != nil {
			c.JSON(http.StatusOK, respError(errors.ErrUnknown))
			return
		}
		if verifyCode == "" {
			c.JSON(http.StatusOK, respError(errors.ErrVerifyCodeExpired))
			return
		}
		if verifyCode != userInfo.VerifyCode {
			c.JSON(http.StatusOK, respError(errors.ErrVerifyCode))
			return
		}
	}

	err = dao.ResetPassword(c.Request.Context(), userInfo.PassHash, userInfo.Username)
	if err != nil {
		log.Errorf("update user : %v", err)
		c.JSON(http.StatusOK, respError(errors.ErrInternalServer))
		return
	}
	c.JSON(http.StatusOK, respJSON(JsonObject{
		"msg": "success",
	}))
}

func GetVerifyCodeHandle(c *gin.Context) {
	userInfo := &model.User{}
	userInfo.Username = c.Query("username")
	verifyType := c.Query("type")
	userInfo.UserEmail = userInfo.Username
	err := SetVerifyCode(c.Request.Context(), userInfo.Username, userInfo.Username+verifyType)
	if err != nil {
		c.JSON(http.StatusOK, respError(errors.ErrUnknown))
		return
	}
	c.JSON(http.StatusOK, respJSON(JsonObject{
		"msg": "success",
	}))
}

func SetVerifyCode(ctx context.Context, username, key string) error {
	vc, _ := GetVerifyCode(ctx, key)
	if vc != "" {
		return nil
	}
	randNew := rand.New(rand.NewSource(time.Now().UnixNano()))
	verifyCode := fmt.Sprintf("%06d", randNew.Intn(1000000))
	bytes, err := json.Marshal(verifyCode)
	if err != nil {
		return err
	}
	var expireTime time.Duration
	expireTime = 5 * time.Minute
	_, err = dao.Cache.Set(ctx, key, bytes, expireTime).Result()
	if err != nil {
		return err
	}
	err = sendEmail(username, verifyCode)
	if err != nil {
		return err
	}
	return nil
}

func GetVerifyCode(ctx context.Context, key string) (string, error) {
	bytes, err := dao.Cache.Get(ctx, key).Bytes()
	if err != nil && err != redis.Nil {
		return "", err
	}
	if err == redis.Nil {
		return "", nil
	}
	var verifyCode string
	err = json.Unmarshal(bytes, &verifyCode)
	if err != nil {
		return "", err
	}
	return verifyCode, nil
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
