package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goCart/pkg/setting"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	t := time.Now()

	savePath := setting.AppSetting.ImageSavePath + fmt.Sprintf("%d%d%d/", t.Year(), t.Month(), t.Day())
	fullPath := setting.AppSetting.RuntimeRootPath + savePath

	if _, err := os.Stat(fullPath); err != nil {
		_ = os.MkdirAll(fullPath, os.ModePerm)
		_, _ = os.Create(fullPath + "index.html")
	}
	fileExt := path.Ext(file.Filename)
	rand.Seed(t.UnixNano())

	fileName := fmt.Sprintf("%v", rand.Uint64()) + fileExt

	if err := c.SaveUploadedFile(file, fullPath+fileName); err != nil {
		fmt.Println(err)

	} else {
		c.JSON(http.StatusOK, map[string]interface{}{"uploadStatus": "success", "uploadPercentage": 100, "code": 200, "data": savePath + fileName})
	}
}
func UploadMulti(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file[]"]
	var fileMap []string
	for _, file := range files {
		t := time.Now()
		savePath := setting.AppSetting.ImageSavePath + fmt.Sprintf("%d%d%d/", t.Year(), t.Month(), t.Day())
		fullPath := setting.AppSetting.RuntimeRootPath + savePath
		if _, err := os.Stat(fullPath); err != nil {
			_ = os.MkdirAll(fullPath, os.ModePerm)
			_, _ = os.Create(fullPath + "index.html")
		}
		fileExt := path.Ext(file.Filename)
		rand.Seed(t.UnixNano())

		fileName := fmt.Sprintf("%v", rand.Uint64()) + fileExt

		if err := c.SaveUploadedFile(file, fullPath+fileName); err != nil {
			fmt.Println(err)
		} else {
			fileMap = append(fileMap, savePath+fileName)
		}
	}
	c.JSON(http.StatusOK, map[string]interface{}{"uploadStatus": "success", "uploadPercentage": 100, "code": 200, "data": fileMap})
}
