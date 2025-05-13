// main.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/mcp", handleMCPRequest)

	port := ":8080"
	fmt.Println("MCP Probe 正在监听端口", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("服务启动失败:", err)
	}
}

func handleMCPRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "仅支持 POST 请求", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "读取请求失败", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("收到 MCP 调用流量")
	parseMCPRequest(body)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
