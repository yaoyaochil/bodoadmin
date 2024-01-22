package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/Wxglobal"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/passport"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/wx_request"
	"io"
)

type MediaApi struct {
}

// UploadNewsMedia 上传永久图文图片素材
func (m *MediaApi) UploadNewsMedia(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 将文件流转为字节切片
	fileBytes, err := io.ReadAll(file)

	// 上传临时素材
	data, err := passport.PassPortGroupApp.MediaPassPort.UploadNewsMedia(wx_request.UploadMediaRequest{
		Type:     "",
		Media:    fileBytes,
		FileName: header.Filename,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if data.Url == "" || data.Errcode != 0 {
		response.FailWithMessage("上传失败", c)
		return
	}
	response.OkWithMessage("上传成功", c)
}

// UploadOtherMedia 上传其它类型永久素材
func (m *MediaApi) UploadOtherMedia(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	type_of := c.Request.FormValue("type")
	// 将文件流转为字节切片
	fileBytes, err := io.ReadAll(file)
	// title和introduction只有视频素材才需要
	if type_of == "video" {
		title := c.Request.FormValue("title")
		introduction := c.Request.FormValue("introduction")
		fmt.Println(title, introduction)
		// 上传临时素材
		data, err := passport.PassPortGroupApp.MediaPassPort.UploadOtherMedia(wx_request.UploadMediaRequest{
			Type:         type_of,
			Media:        fileBytes,
			FileName:     header.Filename,
			Title:        title,
			Introduction: introduction,
		})
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		if data.Url == "" || data.Errcode != 0 {
			response.FailWithMessage("上传失败", c)
			return
		}
		response.OkWithMessage("上传成功", c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 上传其它素材
	data, err := passport.PassPortGroupApp.MediaPassPort.UploadOtherMedia(wx_request.UploadMediaRequest{
		Type:     type_of,
		Media:    fileBytes,
		FileName: header.Filename,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if data.Url == "" || data.Errcode != 0 {
		response.FailWithMessage("上传失败", c)
		return
	}
	response.OkWithMessage("上传成功", c)
}

// GetMediaList 获取素材列表
func (m *MediaApi) GetMediaList(c *gin.Context) {
	var request wx_request.GetMediaListRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := passport.PassPortGroupApp.MediaPassPort.GetMediaList(wx_request.GetMediaListRequest{
		Type:   request.Type,
		Offset: request.Offset,
		Count:  request.Count,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(data, c)
}

type GetMediaRequest struct {
	MediaId string `json:"media_id"`
	Type    string `json:"type"`
}

// GetMediaById 根据media_id获取素材
func (m *MediaApi) GetMediaById(c *gin.Context) {
	var info GetMediaRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if info.Type == "video" {
		data, err := Wxglobal.GlobalConfig.OfficialAccountApp.Material.GetVideo(Wxglobal.GlobalConfig.NullCtx, info.MediaId)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithData(data, c)
		return
	}
	if info.Type == "news" {
		data, err := Wxglobal.GlobalConfig.OfficialAccountApp.Material.GetNews(Wxglobal.GlobalConfig.NullCtx, info.MediaId)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithData(data, c)
		return
	}
	data, err := Wxglobal.GlobalConfig.OfficialAccountApp.Material.Get(Wxglobal.GlobalConfig.NullCtx, info.MediaId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	c.Writer.Header().Set("Content-Type", data.Header.Get("Content-Type"))
	c.Writer.Header().Set("Content-Disposition", data.Header.Get("Content-Disposition"))
	data.Write(c.Writer)
}

// DeleteMediaById 根据media_id删除素材
func (m *MediaApi) DeleteMediaById(c *gin.Context) {
	var info GetMediaRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := Wxglobal.GlobalConfig.OfficialAccountApp.Material.Delete(Wxglobal.GlobalConfig.NullCtx, info.MediaId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if data.ErrCode != 0 {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
