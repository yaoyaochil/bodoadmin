package Api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wecom/ComGlobal"
	"io/ioutil"
)

type CallBackApi struct {
}

// 注册企业微信回调事件
// https://work.weixin.qq.com/api/doc/90000/90135/90930
func (call *CallBackApi) CallBack(c *gin.Context) {
	rs, err := ComGlobal.GlobalConfig.WeComApp.Server.Serve(c.Request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	text, _ := ioutil.ReadAll(rs.Body)
	fmt.Println(string(text))
	c.String(200, string(text))
}
