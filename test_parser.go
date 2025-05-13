package main

func TestParser() {
	sampleJSON := `
	{
		"model": "gpt-4",
		"messages": [
			{"role": "user", "content": "请调用天气查询工具"}
		],
		"tools": [
			{
				"function": {
					"name": "get_weather",
					"arguments": "{\"location\": \"Shanghai\"}"
				}
			}
		]
	}`

	parseMCPRequest([]byte(sampleJSON))
}
