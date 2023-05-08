package system

import (
	"mime/multipart"
	"x-gin-admin/db"
	"x-gin-admin/model"
	"x-gin-admin/utils/response"

	"github.com/adtkcn/dayjs_go/dayjs"
	"github.com/gin-gonic/gin"
)

type FileController struct{}

var FilePrefix string = "."

// 拼接文件保存路径
func JoinPath(filename string) string {
	return "/uploads/" + dayjs.Dayjs().Format("YYYY-MM/DD/HH") + "/" + filename
}

// 真实路径
func GetRealPath(savePath string) string {
	return FilePrefix + savePath
}

// 保存文件到真实路径
func SaveFile(c *gin.Context, file *multipart.FileHeader) (savePath string) {
	savePath = JoinPath(file.Filename)

	c.SaveUploadedFile(file, GetRealPath(savePath))
	return savePath
}

func (common *FileController) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	permission := c.PostForm("permission")
	// var dstDir = JoinPath(file.Filename)
	var savePath = SaveFile(c, file)

	// 上传文件至指定的完整文件路径
	// c.SaveUploadedFile(file, FilePrefix+dstDir)

	var upload = model.UploadFile{
		Name:       file.Filename,
		Path:       savePath,
		Md5:        c.GetString("file_md5"),
		Permission: permission,
	}
	result := db.Sql.Create(&upload)
	if result.Error != nil {
		response.SendError(c, result.Error.Error(), nil)
		return
	}
	// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	response.Send(c, "ok", upload)
}

func (common *FileController) FileByID(c *gin.Context) {
	id := c.Query("id")
	var upload = model.UploadFile{}

	db.Sql.Where("id", id).First(&upload)
	if upload.Permission == "" {
		c.File(GetRealPath(upload.Path))
		return
	}
	// todo: 校验用户权限
	c.File(FilePrefix + upload.Path)
}

// md5复制文件记录
func (common *FileController) UploadWithMd5(c *gin.Context) {
	md5 := c.Query("md5")
	var upload = model.UploadFile{}

	db.Sql.Where("md5", md5).First(&upload)
	if upload.ID == 0 {
		response.SendError(c, "没有记录", nil)
		return
	}
	newUpload := upload // 复制记录
	newUpload.ID = 0    // 修改ID，使其成为新记录
	// newUpload.Name = "New Name" // 修改其他字段值

	db.Sql.Model(&model.UploadFile{}).Create(&newUpload) // 创建新记录，只复制主记录

	// db.Sql.Create(newUpload)

	response.Send(c, "ok", newUpload)
}
