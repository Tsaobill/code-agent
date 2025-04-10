package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ClaudeClient struct {
	apiKey       string
	model        string
	systemPrompt string
	httpClient   *http.Client
}

func NewClaudeClient(apiKey, model string) *ClaudeClient {
	return &ClaudeClient{
		apiKey:       apiKey,
		model:        model,
		httpClient:   &http.Client{},
		systemPrompt: defaultPrompt,
	}
}

func (c *ClaudeClient) WithSysPrompt(prompt string) *ClaudeClient {
	c.systemPrompt = prompt
	return c
}

const defaultPrompt = "你是一个专业的编程助手,被用在通过命令行与用户交互的agentic coding tool中，请尽可能的帮助用户, 如果用户没有特别要求，你需要尽可能保持回答的简洁"
const URL = `https://caobiao.uk/proxy/anthropic/v1/messages`

func (c *ClaudeClient) SendPrompt(prompt string) (string, error) {
	requestBody := map[string]interface{}{
		"model": c.model,
		"messages": []map[string]string{
			{"role": "system", "content": c.systemPrompt},
			{"role": "user", "content": prompt},
		},
		"max_tokens": 4000,
		"functions":  nil,
	}

	jsonBody, _ := json.Marshal(requestBody)

	req, err := http.NewRequest(
		"POST",
		URL,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	// 假设respBody是你的map变量
	contentList, ok := result["content"].([]interface{})
	if !ok || len(contentList) == 0 {
		// 处理content不存在或为空的情况
		return "", fmt.Errorf("invalid response format: result")
	}

	firstContent, ok := contentList[0].(map[string]interface{})
	if !ok {
		// 处理content元素类型不正确的情况
		return "", fmt.Errorf("invalid response format: content")
	}

	text, ok := firstContent["text"].(string)
	if !ok {
		// 处理text不存在或类型不正确的情况
		return "", fmt.Errorf("invalid response format: text")
	}

	// 现在text变量中就包含了你要的内容

	return text, nil

}
