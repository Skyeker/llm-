// detector.go
package main

import (
	"encoding/json"
)

type DetectionResult struct {
	FunctionName string
	Reason       string
}

var blockedFunctions = []string{
	"get_sensitive_info",
	"delete_user",
	"shutdown_server",
}

func DetectAnomalies(mcpData map[string]interface{}) []DetectionResult {
	var results []DetectionResult

	tools, ok := mcpData["tools"].([]interface{})
	if !ok {
		return results
	}

	for _, tool := range tools {
		toolMap, ok := tool.(map[string]interface{})
		if !ok {
			continue
		}

		function, ok := toolMap["function"].(map[string]interface{})
		if !ok {
			continue
		}

		name, ok := function["name"].(string)
		if !ok {
			continue
		}

		for _, blocked := range blockedFunctions {
			if name == blocked {
				results = append(results, DetectionResult{
					FunctionName: name,
					Reason:       "调用了禁止函数",
				})
			}
		}

		// 进一步检测参数内容（如是否包含"root", "passwd"）
		if args, ok := function["arguments"].(string); ok {
			if suspiciousArguments(args) {
				results = append(results, DetectionResult{
					FunctionName: name,
					Reason:       "参数中包含敏感关键词",
				})
			}
		}
	}

	return results
}

func suspiciousArguments(args string) bool {
	keywords := []string{"root", "passwd", "token", "secret"}
	for _, kw := range keywords {
		if json.Valid([]byte(args)) && containsKeyword(args, kw) {
			return true
		}
	}
	return false
}

func containsKeyword(s, keyword string) bool {
	return len(s) > 0 && (stringContainsIgnoreCase(s, keyword))
}

func stringContainsIgnoreCase(s, substr string) bool {
	return len(s) >= len(substr) && (stringIndexIgnoreCase(s, substr) != -1)
}

func stringIndexIgnoreCase(s, substr string) int {
	return len([]rune(s)) - len([]rune(substr)) // 伪代码，可用 strings.Contains(s, substr)
}
