package middleware

import (
	"context"
	"strings"

	"github.com/Tencent/WeKnora/internal/types"
	"github.com/gin-gonic/gin"
)

// Locale returns a middleware that parses the Accept-Language header
func Locale() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Accept-Language header
		acceptLang := c.GetHeader("Accept-Language")
		
		// Default to zh-CN if not provided
		locale := "zh-CN"
		if acceptLang != "" {
			// Take the first language tag
			parts := strings.Split(acceptLang, ",")
			if len(parts) > 0 {
				firstPart := strings.TrimSpace(parts[0])
				// Clean up quality values like en-US;q=0.9
				if idx := strings.Index(firstPart, ";"); idx != -1 {
					locale = strings.TrimSpace(firstPart[:idx])
				} else {
					locale = firstPart
				}
			}
		}

		// Store in gin context
		c.Set(string(types.LocaleContextKey), locale)

		// Store in request context
		ctx := context.WithValue(c.Request.Context(), types.LocaleContextKey, locale)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// GetLocale gets the locale from context
func GetLocale(ctx context.Context) string {
	if locale, ok := ctx.Value(types.LocaleContextKey).(string); ok {
		return locale
	}
	return "zh-CN"
}

