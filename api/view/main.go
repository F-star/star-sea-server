package view

import (
	"fmt"
	"net/http"
	"os"
	"star-sea-server/api/controller"

	"github.com/gin-gonic/gin"
)

type File struct {
	Name string `uri:"name" binding:"required"`
}

func authCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != os.Getenv("TOKEN") {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}

func StartServer() {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.Use(authCheck())
	// upload
	v1.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": "恢复为文件失败" /* err.Error() */})
			return
		}

		n, err := controller.Upload(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": "Uploaded successfully",
			"name":    fmt.Sprintf("%s", n),
		})
	})

	// fileinfo
	v1.GET("/file/list", func(c *gin.Context) {
		fileInfoList, err := controller.GetFileList()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"list":       fileInfoList,
			"url_prefix": os.Getenv("STATIC_HOST"),
		})
	})

	router.Run(":8080")
}
