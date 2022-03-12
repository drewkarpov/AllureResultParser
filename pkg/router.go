package pkg

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Handler struct {
}

func (h *Handler) GetResults(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "cannot find  file with suites.json")
		return
	}
	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, file); err != nil {
		c.String(http.StatusBadRequest, "cannot find  file with suites.json")
		return
	}

	result := Result{}

	err = json.Unmarshal(buffer.Bytes(), &result)

	if _, err := io.Copy(buffer, file); err != nil {
		c.String(http.StatusBadRequest, "cannot find  file with suites.json")
		return
	}

	var failedTests []Suite

	for _, suite := range result.Suites {
		suiteResult := SuiteResult{}
		failedTests = append(failedTests, suiteResult.findSuitesWithParentId(suite)...)
	}
	GetPreparedResults(failedTests)

	c.String(http.StatusOK, GetPreparedResults(failedTests))
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())
	router.POST("/test/results", h.GetResults)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
