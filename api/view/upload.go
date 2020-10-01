package view

import (
	"fmt"
	"net/http"
	"star-sea-server/api/controller"

	"github.com/gin-gonic/gin"
)

type File struct {
	Name string `uri:"name" binding:"required"`
}

func StartServer() {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	files := v1.Group("/files")
	// upload
	files.POST("/", func(c *gin.Context) {
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

	// download
	files.GET("/:name/", func(c *gin.Context) {
		var file File
		if err := c.ShouldBindUri(&file); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}

		m, cn, err := controller.Download(file.Name)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "no such file"})
			return
		}
		c.Header("Content-Disposition", "attachment; filename="+file.Name)
		c.Data(http.StatusOK, m, cn)
	})

	router.Run(":8080")
}
