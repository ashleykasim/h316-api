package middleware

import (
  "github.com/gin-gonic/gin"

  "recipes-api/config"
)

func CORSMiddleware(c *gin.Context) {
  c.Writer.Header().Set("Access-Control-Allow-Origin", config.Get().CORSAllowOrigin)
  c.Writer.Header().Set("Access-Control-Allow-Credentials", config.Get().CORSAllowCredentials)
  c.Writer.Header().Set("Access-Control-Allow-Headers", config.Get().CORSAllowHeaders)
  c.Writer.Header().Set("Access-Control-Allow-Methods", config.Get().CORSAllowMethods)

  if c.Request.Method == "OPTIONS" {
      c.AbortWithStatus(204)
      return
  }

  c.Next()
}
