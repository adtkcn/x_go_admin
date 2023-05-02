package system

import (
	"x-gin-admin/model"
	"x-gin-admin/utils/response"

	"github.com/adtkcn/dayjs_go/dayjs"
	"github.com/gin-gonic/gin"
)

type CommonController struct{}

var FilePrefix string = "."

func (common *CommonController) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	permission := c.PostForm("permission")
	var dstDir = "/uploads/" + dayjs.Dayjs().Format("YYYY-MM/DD/HH") + "/" + file.Filename

	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, FilePrefix+dstDir)

	var upload = model.UploadFile{
		Name:       file.Filename,
		Path:       dstDir,
		Md5:        c.GetString("file_md5"),
		Permission: permission,
	}
	result := model.DB.Create(&upload)
	if result.Error != nil {
		response.SendError(c, result.Error.Error(), nil)
		return
	}
	// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	response.Send(c, "ok", upload)
}

func (common *CommonController) FileByID(c *gin.Context) {
	id := c.Query("id")
	var upload = model.UploadFile{}

	model.DB.Where("id", id).First(&upload)
	if upload.Permission == "" {
		c.File(FilePrefix + upload.Path)
		return
	}
	// todo: 校验用户权限
	c.File(FilePrefix + upload.Path)
}

// md5复制文件记录
func (common *CommonController) CopyWithMd5(c *gin.Context) {
	md5 := c.Query("md5")
	var upload = model.UploadFile{}

	model.DB.Where("md5", md5).First(&upload)
	if upload.ID == 0 {
		response.SendError(c, "没有记录", nil)
		return
	}
	newUpload := upload // 复制记录
	newUpload.ID = 0    // 修改ID，使其成为新记录
	// newUpload.Name = "New Name" // 修改其他字段值

	model.DB.Model(&model.UploadFile{}).Create(&newUpload) // 创建新记录，只复制主记录

	// model.DB.Create(newUpload)

	response.Send(c, "ok", newUpload)
}
