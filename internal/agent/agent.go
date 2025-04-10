package agent

import (
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

	// 调用API
	response, err := a.client.SendPrompt(prompt)
	if err != nil {
		return "", err
	}

	// 可以在这里添加后处理逻辑
	return response, nil
}
