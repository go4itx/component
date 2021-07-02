package main

import (
	"fmt"
	"github.com/go4it-x/component/nlp"
)

func main() {
	chatBot := nlp.NewChatBot(nlp.Param{
		SecretId:  "SecretId",
		SecretKey: "SecretKey",
	})

	response, _ := chatBot.Request("哈喽啊", "", 0)
	fmt.Println(response.ToJsonString())
}
