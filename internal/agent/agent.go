package agent

import (
	"fmt"
	"github.com/Tsaobill/code-agent/internal/api"
)

type CodeAgent struct {
	client *api.ClaudeClient
}

func NewCodeAgent(client *api.ClaudeClient) *CodeAgent {
	return &CodeAgent{client: client}
}

func (a *CodeAgent) Execute(prompt string) (string, error) {
	// 可以在这里添加预处理逻辑
	processedPrompt := fmt.Sprintf("你是一个专业的编程助手,被用在通过命令行与用户交互的agentic coding tool中，请尽可能的帮助用户, 如果用户没有特别要求，你需要尽可能保持回答的简洁。%s", prompt)

	// 调用API
	response, err := a.client.SendPrompt(processedPrompt)
	if err != nil {
		return "", err
	}

	// 可以在这里添加后处理逻辑
	return response, nil
}
