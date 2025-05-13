// mcp_parser.go
package main

import (
	"encoding/json"
	"fmt"
)

// MCPRequest 模拟的 LLM 工具调用结构
type MCPRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Tools    []Tool    `json:"tools"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Tool struct {
	Function Function `json:"function"`
}

type Function struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"` // 注意：也可能是结构体，可扩展
}

func parseMCPRequest(body []byte) {
	var raw map[string]interface{}
	err := json.Unmarshal(body, &raw)
	if err != nil {
		fmt.Println("JSON 解析失败:", err)
		return
	}

	fmt.Println("✅ MCP 请求解析成功")
	model, _ := raw["model"].(string)
	fmt.Println("模型:", model)

	tools, ok := raw["tools"].([]interface{})
	if ok {
		for _, tool := range tools {
			if toolMap, ok := tool.(map[string]interface{}); ok {
				if function, ok := toolMap["function"].(map[string]interface{}); ok {
					name, _ := function["name"].(string)
					args, _ := function["arguments"].(string)
					fmt.Printf("🛠️ 工具调用: %s 参数: %s\n", name, args)
				}
			}
		}
	}

	// 调用检测器并输出异常
	results := DetectAnomalies(raw)
	for _, res := range results {
		fmt.Printf("异常检测: 函数 [%s] 触发规则：%s\n", res.FunctionName, res.Reason)
	}
}
