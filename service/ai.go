package service

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/util"
	"context"
	"errors"
	"io"
	"net/http"
	"sync"

	deepseek "github.com/cohesion-org/deepseek-go"
)

type aiChatService struct {
	client   *deepseek.Client
	onceInit sync.Once
}

func (s *aiChatService) StreamingChat(question string) (ch chan string) {
	// 懒加载初始化
	s.onceInit.Do(func() {
		if setting.Config.Ai.DeepSeek.ApiKey == nil {
			log.Default.Warn("未设置DeepSeek API Key，无法初始化AI Chat服务")
			return
		}
		s.client = deepseek.NewClient(*setting.Config.Ai.DeepSeek.ApiKey)
		log.Default.Info("AI Chat服务初始化成功")
	})

	if s.client == nil {
		panic(util.NewError(http.StatusInternalServerError, "AI Chat服务初始化失败，请联系管理员"))
	}

	ctx := context.Background()

	// 开启对话流
	stream, err := s.client.CreateChatCompletionStream(ctx, &deepseek.StreamChatCompletionRequest{
		Model: deepseek.DeepSeekChat,
		Messages: []deepseek.ChatCompletionMessage{
			{Role: deepseek.ChatMessageRoleUser, Content: question},
		},
		Stream: true,
	})
	if err != nil {
		log.Err.Errorf("创建AI Chat流失败: %v", err)
		panic(util.NewError(http.StatusInternalServerError, "创建AI Chat流失败，请联系管理员"))
	}

	ch = make(chan string)
	go func() {
		defer stream.Close()
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				close(ch)
				break
			}
			if err != nil {
				log.Err.Errorf("AI Chat流错误: %v", err)
				ch <- "发生了未知错误"
				close(ch)
				break
			}
			for _, choice := range response.Choices {
				ch <- choice.Delta.Content
			}
		}
	}()
	return
}
