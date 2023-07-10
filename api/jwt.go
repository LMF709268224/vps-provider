package api

import (
	"net/http"
	"time"

	"vps-provider/types"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	loginStatusFailure = iota
	loginStatusSuccess
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
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
			if v, ok := data.(*types.User); ok {
				return jwt.MapClaims{
					identityKey: v.UUID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &types.User{
				UUID: claims[identityKey].(string),
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
			loginParams.Password = c.Query("password")
			if loginParams.Username == "" || loginParams.Password == "" {
				return "", jwt.ErrMissingLoginValues
			}

			return "", jwt.ErrMissingLoginValues
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(types.User); ok && v.UserName == "admin" {
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
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})
}
