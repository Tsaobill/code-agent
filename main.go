package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Tsaobill/code-agent/internal/agent"
	"github.com/Tsaobill/code-agent/internal/api"
)

func main() {
	var prompt string
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) != os.ModeCharDevice {
		// 有管道输入
		scanner := bufio.NewScanner(os.Stdin)
		var input strings.Builder
		for scanner.Scan() {
			input.WriteString(scanner.Text())
		}
		prompt = input.String()
	} else {
		if len(os.Args) < 2 {
			fmt.Println("Usage: ag <prompt>")
			os.Exit(1)
		}
		prompt = strings.Join(os.Args[1:], " ")
	}

	// 初始化API客户端
	api_key := os.Getenv("ANTHROPIC_API_KEY")
	if api_key == "" {
		fmt.Println("ANTHROPIC_API_KEY environment variable is not set")
		os.Exit(1)
	}
	client := api.NewClaudeClient(
		api_key,
		"claude-3-5-sonnet-20240620", // 可根据需要修改模型版本
	)

	// 创建Agent
	codeAgent := agent.NewCodeAgent(client)

	// 处理请求
	response, err := codeAgent.Execute(prompt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(response)
}
