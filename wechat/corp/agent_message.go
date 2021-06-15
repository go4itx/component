package corp

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

const agentMsgUrl = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%v"

type AgentMsgResult struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
}

// 发送应用文本消息
func (app *App) SendAgentTextMsg(party string, msg string) (res AgentMsgResult, err error) {
	accessToken, err := app.accessToken()
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
		SetBody([]byte(fmt.Sprintf(data, app.AgentId, party, msg))).
		Post(fmt.Sprintf(agentMsgUrl, accessToken))
	if err != nil {
		return
	}

	if res.ErrCode != 0 {
		log.Println(res)
		return res, errors.New(res.ErrMsg)
	}

	return
}
