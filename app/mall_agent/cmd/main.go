package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"gomall/app/mall_agent/eino/mallagent"
)

func main() {
	ctx := context.Background()

	// 初始化ChatModel
	chatModel, err := mallagent.NewArkChatModel(ctx, &mallagent.ChatModelConfig{
		Timeout: 30 * time.Second,
	})
	if err != nil {
		fmt.Printf("初始化ChatModel失败: %v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("已成功连接ARK聊天模型，请输入您的问题（输入'exit'退出）:")

	for {
		fmt.Print("用户: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			break
		}

		// 创建聊天请求
		resp, err := chatModel.Chat(ctx, &ark.ChatRequest{
			Messages: []*ark.Message{
				{Role: "system", Content: mallagent.SystemPromptTemplate},
				{Role: "user", Content: input},
			},
		})

		if err != nil {
			fmt.Printf("请求失败: %v\n", err)
			continue
		}

		fmt.Printf("助手: %s\n\n", resp.Choices[0].Message.Content)
	}
}