package gptclient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/sashabaranov/go-openai"
)

func TestImage(t *testing.T) {
	config := openai.DefaultConfig("sk-Lskydw569JiYB3XW9bF9Fd413cB940D688C28c28Cc363e17")
	config.BaseURL = "https://kapkey.chatgptapi.org.cn/v1"
	client := openai.NewClientWithConfig(config)
	image, err := client.CreateImage(context.Background(), openai.ImageRequest{
		Prompt:         "设计一个带有字母etov的logo图像",
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	})
	client.CreateFile(context.Background(), openai.FileRequest{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(image.Data[0].URL)
}
func TestCompletions(t *testing.T) {
	config := openai.DefaultConfig("sk-Lskydw569JiYB3XW9bF9Fd413cB940D688C28c28Cc363e17")
	config.BaseURL = "https://kapkey.chatgptapi.org.cn/v1"
	client := openai.NewClientWithConfig(config)
	stream, err := client.CreateChatCompletionStream(context.Background(), openai.ChatCompletionRequest{
		Model: "gpt-4-turbo-128k",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "画一个Kubernetes集群的结构图",
			},
		},
		Stream: true,
	})
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
}
