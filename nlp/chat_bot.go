package nlp

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	nlp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/nlp/v20190408"
	"log"
)

type ChatBot struct {
	client  *nlp.Client
	request *nlp.ChatBotRequest
}

// Request 发送请求
func (c *ChatBot) Request(query string, openId string, flag uint64) (response *nlp.ChatBotResponse, err error) {
	if flag != 0 {
		c.request.Flag = common.Uint64Ptr(flag)
	}

	if openId != "" {
		c.request.OpenId = common.StringPtr(openId)
	}

	c.request.Query = common.StringPtr(query)
	response, err = c.client.ChatBot(c.request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		//sdkErr.Message sdkErr.Code
		log.Printf("An API error has returned: %s", err)
		return
	}

	return
}
