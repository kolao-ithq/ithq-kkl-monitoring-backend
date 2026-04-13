package auth

import (
	"codelab/go-gin-template/internal/config"
	"codelab/go-gin-template/pkg/utils/encryption"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiKeyAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("x-api-key") == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}
		encrypt := encryption.EncodeB64(config.APIKeyGlobal.Value)
		if ctx.Request.Header.Get("x-api-key") != encrypt {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// ApiKeyAuthFromParam validates API Key from URL query or route params named "api_key".
func ApiKeyAuthFromParam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.Query("api_key")
		if apiKey == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}
		encrypt := encryption.EncodeB64(config.APIKeyGlobal.Value)
		if apiKey != encrypt {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// ApiKeyAuthAny validates API Key from header (x-api-key), query (?api_key=), or route param (/:api_key)
func ApiKeyAuthAny() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.Request.Header.Get("x-api-key")
		if apiKey == "" {
			apiKey = ctx.Query("api_key")
		}
		if apiKey == "" {
			apiKey = ctx.Param("api_key")
		}
		if apiKey == "" {
			apiKey, _ = ctx.Cookie("x-api-key")
		}
		if apiKey == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		encrypt := encryption.EncodeB64(config.APIKeyGlobal.Value)
		if apiKey != encrypt {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
			ctx.Abort()
			return
		}

		// Persist the validated key in a cookie so subsequent asset requests under /docs/* succeed in browser
		ctx.SetCookie("x-api-key", apiKey, 0, "/", "", false, true)

		ctx.Next()
	}
}
