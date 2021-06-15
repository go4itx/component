package main

import (
	"fmt"
	"github.com/go4it-x/util/wechat/corp"
)

func main() {
	app := corp.App{
		CorpId:     "xxxxxxx",
		AgentId:    "xxxxx",
		CorpSecret: "xxxxxxxxxxxxxxx",
	}

	fmt.Println(app.SendAgentTextMsg("1", "hello world"))
}
