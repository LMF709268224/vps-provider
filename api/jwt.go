package api

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"vps-provider/core/dao"
	"vps-provider/core/errors"
	"vps-provider/core/generated/model"
	"vps-provider/core/oplog"
	"vps-provider/utils"
)

const (
	loginStatusFailure = iota
	loginStatusSuccess
)

type login struct {
	Username   string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	VerifyCode string `form:"verify_code" json:"verify_code" binding:"required"`
}

type loginResponse struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}

var identityKey = "id"

func jwtGinMiddleware(secretKey string) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "User",
		Key:         []byte(secretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.Uuid,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				Uuid: claims[identityKey].(string),
			}
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"data": loginResponse{
					Token:  token,
					Expire: expire.Format(time.RFC3339),
				},
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
			})
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginParams login
			loginParams.Username = c.Query("username")
			loginParams.VerifyCode = c.Query("verify_code")
			loginParams.Password = c.Query("password")
			if loginParams.Username == "" {
				return "", jwt.ErrMissingLoginValues
			}
			if loginParams.VerifyCode == "" && loginParams.Password == "" {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginParams.Username
			password := loginParams.Password
			verifyCode := loginParams.VerifyCode
			userAgent := c.Request.Header.Get("User-Agent")
			ua := user_agent.New(userAgent)
			os := ua.OS()
			explorer, _ := ua.Browser()
			clientIP := utils.GetClientIP(c.Request)
			location := utils.GetLocationByIP(clientIP)
			var err error
			var user interface{}
			if verifyCode != "" {
				user, err = loginByVerifyCode(c.Request.Context(), userID, verifyCode)
			} else {
				user, err = loginByPassword(c.Request.Context(), userID, password)
			}

			if err != nil {
				oplog.AddLoginLog(&model.LoginLog{
					IpAddress:     clientIP,
					Browser:       explorer,
					Os:            os,
					Status:        loginStatusFailure,
					Msg:           err.Error(),
					LoginLocation: location,
				})
				return nil, err
			}

			oplog.AddLoginLog(&model.LoginLog{
				LoginUsername: userID,
				LoginLocation: location,
				IpAddress:     clientIP,
				Browser:       explorer,
				Os:            os,
				Status:        loginStatusSuccess,
				Msg:           "success",
			})
			return user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(model.User); ok && v.Username == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(200, gin.H{
				"code":    200,
				"msg":     message,
				"success": false,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}

func loginByPassword(ctx context.Context, username, password string) (interface{}, error) {
	user, err := dao.GetUserByUsername(ctx, username)
	if err != nil {
		log.Errorf("get user by username: %v", err)
		return nil, errors.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(password)); err != nil {
		log.Errorf("can't compare hash %s ans password %s: %v", user.PassHash, password, err)
		return nil, errors.ErrInvalidPassword
	}

	return &model.User{Uuid: user.Uuid, Username: user.Username, Role: user.Role}, nil
}

func loginByVerifyCode(ctx context.Context, username, userVerifyCode string) (interface{}, error) {
	verifyCode, err := GetVerifyCode(ctx, username+"2")
	if err != nil {
		log.Errorf("get user by verify code: %v", err)
		return nil, errors.ErrUnknown
	}
	if verifyCode == "" {
		return nil, errors.ErrVerifyCodeExpired
	}
	user, err := dao.GetUserByUsername(ctx, username)
	if err != nil {
		log.Errorf("get user by username: %v", err)
		return nil, errors.ErrUserNotFound
	}
	if verifyCode != userVerifyCode {
		return nil, errors.ErrVerifyCode
	}

	return &model.User{Uuid: user.Uuid, Username: user.Username, Role: user.Role}, nil
}