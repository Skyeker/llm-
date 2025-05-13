# llm-traffic probe
基于Go 的高并发特性流量探针， 能够监听 MCP/HTTP2 等协议，分析并转存 LLM 调用请求与响应等

# demo
运行
```go run  main.go mcp_parser.go detector.go
```

再用curl发送一条调用 delete_user 的 MCP 请求：
```
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-4",
    "messages": [{"role": "user", "content": "请帮我删除一个用户"}],
    "tools": [{
      "function": {
        "name": "delete_user",
        "arguments": "{\"user\": \"admin\"}"
      }
    }]
  }'
```

#待完善
