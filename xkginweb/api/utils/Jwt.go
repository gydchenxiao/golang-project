package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/scrypt"
	"io"
	"time"
)

// 完整的代码可见github: https://github.com/Messi-Q/Beego-Jwt
// go get github.com/dgrijalva/jwt-go
// 官网：https://jwt.io/introduction/

// JWT : HEADER PAYLOAD SIGNATURE
const (
	SecretKEY              string = "JWT-Secret-Key"
	DEFAULT_EXPIRE_SECONDS int    = 600 // default expired 10 minutes
	PasswordHashBytes             = 16
)

// This struct is the payload
type MyCustomClaims struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
}

// This struct is the parsing of token payload
type JwtPayload struct {
	Username  string `json:"username"`
	UserID    int    `json:"userID"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

// @Author feige
// @Description 创建token
// @Date  2022-12-12 20:23
// @Param
// @Return
func GenerateToken(username string, userID int, expiredSeconds int) (tokenString string, err error) {
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}

	// Create the Claims
	mySigningKey := []byte(SecretKEY)
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	logs.Info("Token will be expired at ", time.Unix(expireAt, 0))

	claims := MyCustomClaims{
		userID,
		jwt.StandardClaims{
			Issuer:    username,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireAt,
		},
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Signs the token with a secret
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", errors.New("error: failed to generate token")
	}

	return tokenStr, nil
}

// @Author feige
// @Description 验证token
// @Date  2022-12-12 20:23
// @Param
// @Return
func ValidateToken(tokenString string) (*JwtPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKEY), nil
		})

	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		logs.Info("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt)
		logs.Info("Token will be expired at ", time.Unix(claims.StandardClaims.ExpiresAt, 0))

		return &JwtPayload{
			Username:  claims.StandardClaims.Issuer,
			UserID:    claims.UserID,
			IssuedAt:  claims.StandardClaims.IssuedAt,
			ExpiresAt: claims.StandardClaims.ExpiresAt,
		}, nil
	} else {
		logs.Info(err.Error())
		return nil, errors.New("error: failed to validate token")
	}
}

// @Author feige
// @Description 刷新token
// @Date  2022-12-12 20:23
// @Param
// @Return
func RefreshToken(tokenString string) (newTokenString string, err error) {
	// get previous token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKEY), nil
		})

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return "", err
	}

	mySigningKey := []byte(SecretKEY)
	expireAt := time.Now().Add(time.Second * time.Duration(DEFAULT_EXPIRE_SECONDS)).Unix() //new expired
	newClaims := MyCustomClaims{
		claims.UserID,
		jwt.StandardClaims{
			Issuer:    claims.StandardClaims.Issuer, //name of token issue
			IssuedAt:  time.Now().Unix(),            //time of token issue
			ExpiresAt: expireAt,
		},
	}

	// generate new token with new claims
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	// sign the token with a secret
	tokenStr, err := newToken.SignedString(mySigningKey)
	if err != nil {
		return "", errors.New("error: failed to generate new fresh json api token")
	}

	return tokenStr, nil
}

// generate salt
func GenerateSalt() (salt string, err error) {
	buf := make([]byte, PasswordHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", errors.New("error: failed to generate user's salt")
	}

	return fmt.Sprintf("%x", buf), nil
}

// generate password hash
func GeneratePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, PasswordHashBytes)
	if err != nil {
		return "", errors.New("error: failed to generate password hash")
	}

	return fmt.Sprintf("%x", h), nil
}
