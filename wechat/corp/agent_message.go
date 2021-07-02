package corp

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

const agentMessageUrl = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%v"

type agentMessage struct {
	param Param
}

func NewAgentMessage(param Param) *agentMessage {
	return &agentMessage{
		param: param,
	}
}

type AgentMessageResult struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
}

// 发送应用文本消息
func (a *agentMessage) SendAgentTextMsg(party string, msg string) (res AgentMessageResult, err error) {
	accessToken, err := accessToken(a.param)
	if err != nil {
		return
	}

	data := `{
		"agentid" : %v,
		"toparty" : "%v",
		"msgtype" : "text",
		"text" : {
		   "content" : "%v"
		},
		"safe":0,
		"enable_id_trans": 0,
		"enable_duplicate_check": 0,
		"duplicate_check_interval": 1800
	}`

	_, err = resty.New().R().
		SetResult(&res).
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(fmt.Sprintf(data, a.param.AgentId, party, msg))).
		Post(fmt.Sprintf(agentMessageUrl, accessToken))
	if err != nil {
		return
	}

	if res.ErrCode != 0 {
		log.Println(res)
		return res, errors.New(res.ErrMsg)
	}

	return
}
