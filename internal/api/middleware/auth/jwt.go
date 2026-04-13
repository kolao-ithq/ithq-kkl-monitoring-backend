package auth

import (
	"codelab/go-gin-template/internal/api/middleware/response"
	"codelab/go-gin-template/internal/config"
	"codelab/go-gin-template/pkg/utils/format"
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kataras/jwt"
)

type ClaimsToken struct {
	ID        int   `json:"id,omitempty"`
	IssuedAt  int64 `json:"iat,omitempty"`
	ExpiresAt int64 `json:"exp,omitempty"`
}

var (
	JWT_ACCESS_EXSPIRATION  = 7  // days
	JWT_REFRESH_EXSPIRATION = 30 // days
)

type TokenPair struct {
	AccessToken  json.RawMessage `json:"access_token,omitempty"`
	RefreshToken json.RawMessage `json:"refresh_token,omitempty"`
}

func GenerateJWTToken(id int) (*TokenPair, error) {
	standardClaims := ClaimsToken{
		ID:        id,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * time.Duration(JWT_ACCESS_EXSPIRATION)).Unix(),
	}
	encrypt, _, err := jwt.GCM([]byte(config.JWTGlobal.Secret), nil)
	if err != nil {

		return nil, err
	}
	token, err := jwt.SignEncrypted(jwt.HS256, []byte(config.JWTGlobal.Secret), encrypt, standardClaims, jwt.MaxAge(time.Duration(JWT_ACCESS_EXSPIRATION)*24*time.Hour)) //7Days
	if err != nil {

		return nil, err
	}
	reEncrypt, _, _ := jwt.GCM([]byte(config.JWTGlobal.RefreshSecret), nil)
	refreshToken, err := jwt.SignEncrypted(jwt.HS256, []byte(config.JWTGlobal.RefreshSecret), reEncrypt, standardClaims, jwt.MaxAge(time.Duration(JWT_REFRESH_EXSPIRATION)*24*time.Hour)) //30Days
	if err != nil {
		return nil, err
	}
	tokenPairData := jwt.NewTokenPair(token, refreshToken)
	return &TokenPair{
		AccessToken:  format.BytesQuote(tokenPairData.AccessToken),
		RefreshToken: format.BytesQuote(tokenPairData.RefreshToken),
	}, nil
}

// AccessToken is a Gin middleware that validates the access token and continues
// the handler chain if the token is valid. On failure it writes an error
// response and aborts the request.
func AccessToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, decrypt, _ := jwt.GCM([]byte(config.JWTGlobal.Secret), nil)
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			_ = response.NewErrorErrMsgUnauthorized(ctx)
			ctx.Abort()
			return
		}
		jwtFromHeader := strings.TrimSpace(auth[7:])
		verifiedToken, err := jwt.VerifyEncrypted(jwt.HS256, []byte(config.JWTGlobal.Secret), decrypt, []byte(jwtFromHeader))
		if err != nil {
			_ = response.NewErrorErrMsgUnauthorized(ctx)
			ctx.Abort()
			return
		}
		var claims ClaimsToken
		if err := verifiedToken.Claims(&claims); err != nil {
			_ = response.NewErrorErrMsgUnauthorized(ctx)
			ctx.Abort()
			return
		}

		// Optionally store claims in context for downstream handlers.
		ctx.Set("claims", claims)

		// Token is valid; continue to the next handler.
		ctx.Next()
	}
}

// AccessRefreshToken is a Gin middleware that validates the refresh token and
// continues the handler chain if the token is valid. On failure it writes an
// error response and aborts the request.
func AccessRefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, decrypt, _ := jwt.GCM([]byte(config.JWTGlobal.RefreshSecret), nil)
		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			_ = response.NewErrorErrMsgUnauthorized(ctx)
			ctx.Abort()
			return
		}
		jwtFromHeader := strings.TrimSpace(auth[7:])
		verifiedToken, err := jwt.VerifyEncrypted(jwt.HS256, []byte(config.JWTGlobal.RefreshSecret), decrypt, []byte(jwtFromHeader))
		if err != nil {
			_ = response.NewErrorErrMsgUnauthorized(ctx)
			ctx.Abort()
			return
		}
		var claims ClaimsToken
		if err := verifiedToken.Claims(&claims); err != nil {
			_ = response.NewErrorErrMsgUnauthorized(ctx)
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}

func GetOwnerAccessToken(ctx *gin.Context) (*ClaimsToken, error) {
	_, decrypt, _ := jwt.GCM([]byte(config.JWTGlobal.Secret), nil)
	auth := ctx.Request.Header.Get("Authorization")
	jwtFromHeader := strings.TrimSpace(auth[7:])
	verifiedToken, err := jwt.VerifyEncrypted(jwt.HS256, []byte(config.JWTGlobal.Secret), decrypt, []byte(jwtFromHeader))
	if err != nil {
		return nil, err
	}
	var claims ClaimsToken
	err = verifiedToken.Claims(&claims)
	if err != nil {
		return nil, err
	}
	return &claims, nil
}

func GetOwnerRefresh(ctx *gin.Context) (*ClaimsToken, error) {
	_, decrypt, _ := jwt.GCM([]byte(config.JWTGlobal.RefreshSecret), nil)
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		return nil, response.NewErrorErrMsgUnauthorized(ctx)
	}
	jwtFromHeader := strings.TrimSpace(auth[7:])

	verifiedToken, err := jwt.VerifyEncrypted(jwt.HS256, []byte(config.JWTGlobal.RefreshSecret), decrypt, []byte(jwtFromHeader))
	if err != nil {
		return nil, err
	}
	var claims ClaimsToken
	err = verifiedToken.Claims(&claims)
	if err != nil {
		return nil, err
	}
	resp := ClaimsToken{
		ID: claims.ID,
	}
	return &resp, nil
}

// TODO: Refresh token
func GenerateRefreshToken(ctx *gin.Context) (*TokenPair, error) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		return nil, response.NewErrorErrMsgUnauthorized(ctx)
	}
	claims, err := GetOwnerRefresh(ctx)
	if err != nil {
		return nil, err
	}
	standardClaims := ClaimsToken{
		ID:        claims.ID,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * time.Duration(JWT_ACCESS_EXSPIRATION)).Unix(),
	}
	encrypt, _, _ := jwt.GCM([]byte(config.JWTGlobal.Secret), nil)
	token, err := jwt.SignEncrypted(jwt.HS256, []byte(config.JWTGlobal.Secret), encrypt, standardClaims, jwt.MaxAge(time.Duration(JWT_ACCESS_EXSPIRATION)*24*time.Hour))
	if err != nil {
		return nil, err
	}
	reEncrypt, _, _ := jwt.GCM([]byte(config.JWTGlobal.RefreshSecret), nil)
	refreshToken, err := jwt.SignEncrypted(jwt.HS256, []byte(config.JWTGlobal.RefreshSecret), reEncrypt, standardClaims, jwt.MaxAge(time.Duration(JWT_REFRESH_EXSPIRATION)*24*time.Hour))
	if err != nil {

		return nil, err
	}
	tokenPairData := jwt.NewTokenPair(token, refreshToken)
	return &TokenPair{
		AccessToken:  format.BytesQuote(tokenPairData.AccessToken),
		RefreshToken: format.BytesQuote(tokenPairData.RefreshToken),
	}, nil
}
