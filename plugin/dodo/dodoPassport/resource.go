package dodoPassport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	BODO_Global "github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/dodo_global"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/model"
	"github.com/yaoyaochil/bodo-admin-server/server/utils"
	"io"
	"mime/multipart"
	"regexp"
)

type DoDoResourcePassport struct{}
type ResourceReponse struct {
	Data struct {
		Height int    `json:"height"`
		Url    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// UploadResourceToDodo 上传图片资源到DODO并保存相关信息到数据库
func (d *DoDoResourcePassport) UploadResourceToDodo(c *gin.Context) (image ResourceReponse, err error) {
	// 从请求中读取文件
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return image, err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return image, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		return image, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return image, err
	}
	err = writer.Close()
	if err != nil {
		return image, err
	}
	url := "https://botopen.imdodo.com/api/v2/resource/picture/upload"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = dodo_global.GlobalConfig.Authorization
	headers["Content-Type"] = writer.FormDataContentType()
	requestFile, err := utils.SendRequestFile(url, "POST", body, headers)
	if err != nil {
		return ResourceReponse{}, err
	}
	err = json.Unmarshal(requestFile, &image)
	if err != nil {
		return ResourceReponse{}, err
	}
	// 保存图片资源到数据库
	name, err := d.GetImageSourceName(image.Data.Url)
	if err != nil {
		return image, err
	}
	err = BODO_Global.BODO_DB.Create(&model.DoDoImageSource{
		ImageName: name,
		Url:       image.Data.Url,
		Height:    image.Data.Height,
		Width:     image.Data.Width,
	}).Error
	return
}

// GetImageSourceName 获取图片资源名称
func (d *DoDoResourcePassport) GetImageSourceName(url string) (imageName string, err error) {
	re := regexp.MustCompile(`\/cdn\/(.*?)\.(jpg|png|webp)`)
	match := re.FindStringSubmatch(url)
	if len(match) < 1 {
		return "", fmt.Errorf("图片资源名称获取失败")
	}
	imageName = match[1]
	return imageName, nil
}
