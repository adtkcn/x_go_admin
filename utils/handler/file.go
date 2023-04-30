package handler

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 计算文件md5,还是放前端
func CalculateFileMD5() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "file is required",
			})
			return
		}

		src, err := file.Open()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "failed to open file",
			})
			return
		}
		defer src.Close()

		hash := md5.New()
		if _, err := io.Copy(hash, src); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "failed to calculate file md5",
			})
			return
		}

		md5Hash := fmt.Sprintf("%x", hash.Sum(nil))
		c.Set("file_md5", md5Hash)
		c.Next()
	}
}
