package jwtgo

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
	"xkginweb/commons/response"
	"xkginweb/model/entity/jwt"
	"xkginweb/utils"
)

var jwtService = JwtService{}

// 定义一个JWTAuth的中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过http header中的token解析来认证
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			response.Fail(701, "请求未携带token，无权限访问", c)
			c.Abort()
			return
		}

		if jwtService.IsBlacklist(authorization) {
			response.Fail(601, "您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}

		// 初始化一个JWT对象实例，并根据结构体方法来解析token
		j := NewJWT()
		// 解析token中包含的相关信息(有效载荷)
		claims, err := j.ParserToken(authorization)

		if err != nil {
			// 如果token过期
			if errors.Is(err, TokenExpired) {
				response.Fail(601, "token授权已过期，请重新申请授权", c)
				c.Abort()
				return
			}
			// 其他错误
			response.Fail(602, err.Error(), c)
			c.Abort()
			return
		}

		fmt.Println("claims.ExpiresAt-time.Now().Unix()", claims.ExpiresAt)
		fmt.Println("claims.ExpiresAt-time.Now().Unix()", time.Now().Unix())
		fmt.Println("claims.ExpiresAt-time.Now().Unix()", claims.ExpiresAt-time.Now().Unix())
		fmt.Println("claims.BufferTime", claims.BufferTime)
		// 判断过期时间和当前时间比较，是否已经小于了缓存时间，小于了就开始进行续期
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			// 获取过期时间
			dr, _ := utils.ParseDuration("2d")
			// 时间开始累加
			claims.ExpiresAt = time.Now().Add(dr).Unix()
			// 生成新的token
			newToken, _ := j.CreateTokenByOldToken(authorization, *claims)
			// 开始解析新的token
			newClaims, _ := j.ParserToken(newToken)
			// 返回给响应
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))

			// 判断当前用户是否在redis中
			RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
			if err != nil {
				fmt.Errorf("get redis jwt failed", zap.Error(err))
			} else {
				// 当之前的取成功时才进行拉黑操作
				_ = jwtService.JsonInBlacklist(jwt.JwtBlacklist{Jwt: RedisJwtToken})
			}
			// 无论如何都要记录当前的活跃状态
			_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
		}

		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("claims", claims)
		c.Next()
	}
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
//func (j *JWT) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
//	group := singleflight.Group{}
//	v, err, _ := group.Do("JWT:"+oldToken, func() (interface{}, error) {
//		return j.CreateToken(claims)
//	})
//	return v.(string), err
//}
