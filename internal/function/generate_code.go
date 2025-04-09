package function

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type CodeBlock struct {
	Language string
	Code     string
}

func extractCodeBlocks(answer string) []CodeBlock {
	var blocks []CodeBlock
	// 使用正则表达式匹配代码块
	re := regexp.MustCompile("```([a-zA-Z0-9]*)\n([\\s\\S]*?)```")
	matches := re.FindAllStringSubmatch(answer, -1)

	for _, match := range matches {
		language := match[1]
		code := match[2]
		blocks = append(blocks, CodeBlock{
			Language: language,
			Code:     code,
		})
	}
	return blocks
}

func generateFile(filename, content string) (string, error) {
	dir := filepath.Dir(filename)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", fmt.Errorf("创建目录失败: %v", err)
		}
	}

	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("写入文件失败: %v", err)
	}

	return fmt.Sprintf("成功生成文件: %s", filename), nil
}
