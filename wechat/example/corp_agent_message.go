package main

import (
	"fmt"
	"github.com/go4it-x/component/wechat/corp"
)

func main() {
	am := corp.NewAgentMessage(corp.Param{
		AgentId:    "AgentId",
		CorpId:     "CorpId",
		CorpSecret: "CorpSecret",
	})

	fmt.Println(am.SendAgentTextMsg("1", "hello world"))
}
