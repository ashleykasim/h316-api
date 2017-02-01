package models

import (
  "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
  "testing"
)

func TestDownloadFile(t *testing.T) {

  gin.SetMode(gin.TestMode)

  handler := func(c *gin.Context) {
		c.String(http.StatusOK, "bar")
	}

	router := gin.New()
	router.GET("/files", handler)

	req, _ := http.NewRequest("GET", "/files", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "bar")
}
