package passport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/Wxglobal"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/wx_request"
	"github.com/yaoyaochil/bodo-admin-server/server/utils"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type MediaPassPort struct{}

// UploadNewsMedia 上传永久图文素材
func (m *MediaPassPort) UploadNewsMedia(requests wx_request.UploadMediaRequest) (data wx_request.UploadMediaResponse, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=" + Wxglobal.GlobalConfig.AccessToken
	client := &http.Client{}
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := writer.CreateFormFile("media", requests.FileName)
	file.Write(requests.Media)
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = json.Unmarshal(body, &data)
	return
}

// UploadOtherMedia 上传其它类型永久素材
func (m *MediaPassPort) UploadOtherMedia(requests wx_request.UploadMediaRequest) (data wx_request.UploadMediaResponse, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s", Wxglobal.GlobalConfig.AccessToken, requests.Type)
	client := &http.Client{}
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := writer.CreateFormFile("media", requests.FileName)
	_, err = file.Write(requests.Media)
	if err != nil {
		return wx_request.UploadMediaResponse{}, err
	}

	if requests.Type == "video" {
		writer.WriteField("description", fmt.Sprintf("{\"title\":\"%s\",\"introduction\":\"%s\"}", requests.Title, requests.Introduction))
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	_ = json.Unmarshal(body, &data)
	return
}

// GetMediaList 获取素材列表
func (m *MediaPassPort) GetMediaList(requests wx_request.GetMediaListRequest) (data interface{}, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%s", Wxglobal.GlobalConfig.AccessToken)
	payload := fmt.Sprintf("{\"type\":\"%s\",\"offset\":%d,\"count\":%d}", requests.Type, requests.Offset-1, requests.Count)
	res, err := utils.SendRequest(url, "POST", payload, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	if requests.Type == "news" {
		var newsResponse wx_request.GetMediaNewsListResponse
		_ = json.Unmarshal(res, &newsResponse)
		return newsResponse, nil
	}
	var otherResponse wx_request.GetMediaOtherListResponse
	_ = json.Unmarshal(res, &otherResponse)
	return otherResponse, nil
}
