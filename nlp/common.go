package nlp

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	nlp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/nlp/v20190408"
)

type Param struct {
	SecretId  string
	SecretKey string
	Region    string
}

// initClient new Client
func initClient(p Param) (client *nlp.Client) {
	if p.Region == "" {
		p.Region = "ap-guangzhou"
	}

	credential := common.NewCredential(
		p.SecretId,
		p.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "nlp.tencentcloudapi.com"
	client, _ = nlp.NewClient(credential, p.Region, cpf)
	return
}

// NewChatBot 对话机器人
func NewChatBot(p Param) (chatBot *ChatBot) {
	return &ChatBot{
		client:  initClient(p),
		request: nlp.NewChatBotRequest(),
	}
}
